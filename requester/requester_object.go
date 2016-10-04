// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package requester

import (
    // stdlib
    "bytes"
    "errors"
    "fmt"
    "net"
    "strconv"
    "strings"
    "sync"
    "time"

    // local
    "github.com/pztrn/urtrator/datamodels"
)

type Requester struct {
    // Master server address
    master_server string
    // Master server port
    master_server_port string
    // Packet prefix.
    pp string
    // Showstopper for delimiting IP addresses.
    // As we are receiving bytes, we will put bytes representation of "\"
    // character.
    ip_delimiter int
}

// Requester's initialization.
func (r *Requester) Initialize() {
    fmt.Println("Initializing Requester...")
    r.master_server = "master.urbanterror.info"
    r.master_server_port = "27900"
    r.pp = "\377\377\377\377"
    r.ip_delimiter = 92
}

// Gets all available servers from master server.
func (r *Requester) getServers(callback chan [][]string) {
    // IP addresses we will compose to return.
    var received_ips [][]string
    conn, err1 := net.Dial("udp", r.master_server + ":" + r.master_server_port)
    if err1 != nil {
        fmt.Println("Error dialing to master server!")
    }
    defer conn.Close()

    // Set deadline, so we won't wait forever.
    ddl := time.Now()
    // This should be enough. Maybe, you should'n run URTrator on modem
    // connections? :)
    ddl = ddl.Add(time.Second * 2)
    conn.SetDeadline(ddl)
    fmt.Println("Master server connection deadline set: " + ddl.String())

    // Request.
    // 68 - protocol version for 4.3
    // 70 - protocol version for 4.2.x
    msg := []byte(r.pp + "getservers 68 full empty")
    conn.Write(msg)

    // UDP Buffer.
    var received_buf []byte = make([]byte, 4096)
    // Received buffer.
    var raw_received []byte
    // Received IP addresses.
    //var received []string
    fmt.Println("Receiving servers list...")
    for {
        _, err := conn.Read(received_buf)
        if err != nil {
            break
        }
        raw_received = append(raw_received, received_buf...)
        fmt.Println("Received " + strconv.Itoa(len(raw_received)) + " bytes")
    }

    // Obtaining list of slices.
    var raw_received_slices [][]byte = bytes.Split(raw_received, []byte("\\"))

    // Every ip-port pair contains:
    // 1. IP as first 4 bytes
    // 2. Port as last 2 bytes.
    // Every package is a 7-bytes sequence, which starts with "\"
    // (code 92), which we used before to obtain list of slices.
    for _, slice := range raw_received_slices {
        // We need only 6-bytes slices. All other aren't represent
        // server's address.
        if len(slice) != 6 {
            continue
        }
        // Generate IP.
        ip := strconv.Itoa(int(slice[0])) + "." + strconv.Itoa(int(slice[1])) + "." + strconv.Itoa(int(slice[2])) + "." + strconv.Itoa(int(slice[3]))
        // Generate port from last two bytes.
        // This is a very shitty thing. Don't do this in real world.
        // Maybe bitshifting will help here, but I can't get it to work :(
        // Get first byte as integer and multiply it on 256 and summing with
        // second byte.
        p1 := int(slice[4]) * 256
        port := strconv.Itoa(p1 + int(slice[5]))

        // Create a slice with IP and port.
        ip_and_port := []string{ip, port}
        // Add it to received_ips.
        received_ips = append(received_ips, ip_and_port)
    }

    fmt.Println("Parsed " + strconv.Itoa(len(received_ips)) + " addresses")
    callback <- received_ips
}

// Updates information about all available servers from master server and
// parses it to usable format.
func (r *Requester) UpdateAllServers(done_chan chan map[string]*datamodels.Server, error_chan chan bool) {
    fmt.Println("Starting all servers updating procedure...")

    callback := make(chan [][]string)
    go r.getServers(callback)

    servers := make(map[string]*datamodels.Server)

    select {
    case data := <- callback:
        // Yay, we got data! :)
        fmt.Println("Received " + strconv.Itoa(len(data)) + " servers")
        servers = r.updateServerGoroutineDispatcher(data)
        break
    case <- time.After(time.Second * 10):
        // Timeouted? Okay, push error back.
        error_chan <- true
    }

    done_chan <- servers
}

func (r *Requester) UpdateFavoriteServers(servers [][]string, done_chan chan map[string]*datamodels.Server, error_chan chan bool) {
    fmt.Println("Updating favorites servers...")
    updated_servers := r.updateServerGoroutineDispatcher(servers)
    done_chan <- updated_servers
}

func (r *Requester) updateServerGoroutineDispatcher(data [][]string) map[string]*datamodels.Server {
    var wait sync.WaitGroup
    var lock = sync.RWMutex{}
    done_updating := 0
    servers := make(map[string]*datamodels.Server)

    for _, s := range data {
        s := datamodels.Server{
            Ip: s[0],
            Port: s[1],
        }
        go func(s *datamodels.Server, servers map[string]*datamodels.Server) {
            wait.Add(1)
            defer wait.Done()
            r.UpdateServer(s)
            done_updating = done_updating + 1
            lock.Lock()
            servers[s.Ip + ":" + s.Port] = s
            lock.Unlock()
        }(&s, servers)
    }
    wait.Wait()
    return servers
}

// Updates information about specific server.
func (r *Requester) UpdateServer(server *datamodels.Server) error {
    srv := server.Ip + ":" + server.Port
    fmt.Println("Updating server: " + srv)

    // Dial to server.
    conn, err1 := net.Dial("udp", srv)
    if err1 != nil {
        fmt.Println("Error dialing to server " + srv + "!")
        return errors.New("Error dialing to server " + srv + "!")
    }
    defer conn.Close()

    // Set deadline, so we won't wait forever.
    ddl := time.Now()
    // This should be enough. Maybe, you should'n run URTrator on modem
    // connections? :)
    ddl = ddl.Add(time.Second * 2)
    conn.SetDeadline(ddl)

    msg := []byte(r.pp + "getinfo")
    conn.Write(msg)

    // UDP Buffer.
    var received_buf []byte = make([]byte, 4096)
    // Received buffer.
    var raw_received []byte
    for {
        _, err := conn.Read(received_buf)
        if err != nil {
            break
        }
        raw_received = append(raw_received, received_buf...)
    }

    // First line is "infoResponse" string, which we should skip by
    // splitting response by "\n".
    received_lines := strings.Split(string(raw_received), "\n")
    // We have server's data!
    if len(received_lines) > 1 {
        srv_config := strings.Split(received_lines[1], "\\")
        // Parse server configuration into passed server's datamodel.
        for i := 0; i < len(srv_config); i = i + 1 {
            if srv_config[i] == "modversion" {
                server.Version = srv_config[i + 1]
            }
            if srv_config[i] == "gametype" {
                server.Gamemode = srv_config[i + 1]
            }
            if srv_config[i] == "sv_maxclients" {
                server.Maxplayers = srv_config[i + 1]
            }
            if srv_config[i] == "clients" {
                server.Players = srv_config[i + 1]
            }
            if srv_config[i] == "mapname" {
                server.Map = srv_config[i + 1]
            }
            if srv_config[i] == "hostname" {
                server.Name = srv_config[i + 1]
            }
        }
    }

    // ToDo: Calculate ping. 0 for now.
    server.Ping = "0"

    return nil
}
