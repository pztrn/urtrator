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
    "errors"
    "fmt"
    "net"
    "runtime"
    "strconv"
    "strings"
    "sync"
    "time"

    // local
    "github.com/pztrn/urtrator/datamodels"
)

type Pooler struct {
    // Maximum number of simultaneous requests running.
    maxrequests int
    // Packet prefix.
    pp string
    // Current requests counter mutex.
    cur_requests_mutex sync.Mutex
}

func (p *Pooler) Initialize() {
    fmt.Println("Initializing requester goroutine pooler...")
    // ToDo: figure out how to make this work nice.
    p.maxrequests = 200
    _ = runtime.GOMAXPROCS(runtime.NumCPU() * 4)
    p.pp = "\377\377\377\377"
    fmt.Println("Pooler initialized")
}

func (p *Pooler) PingOneServer(server_address string) {
    var wait sync.WaitGroup

    server := Cache.Servers[server_address].Server

    wait.Add(1)
    go func(srv *datamodels.Server) {
        defer wait.Done()
        p.pingServersExecutor(srv)
    }(server)
    wait.Wait()
}

// Servers pinging pooler. Should be started as goroutine to prevent
// UI blocking.
func (p *Pooler) PingServers(servers_type string) {
    fmt.Println("About to ping " + servers_type + " servers...")

    cur_requests := 0
    var wait sync.WaitGroup

    for _, server_to_ping := range Cache.Servers {
        if servers_type == "favorites" && server_to_ping.Server.Favorite != "1" {
            continue
        }
        for {
            if cur_requests == p.maxrequests {
                time.Sleep(time.Second * 1)
            } else {
                break
            }
        }
        wait.Add(1)
        p.cur_requests_mutex.Lock()
        cur_requests += 1
        p.cur_requests_mutex.Unlock()
        go func(srv *datamodels.Server) {
            defer wait.Done()
            p.pingServersExecutor(srv)
            p.cur_requests_mutex.Lock()
            cur_requests -= 1
            p.cur_requests_mutex.Unlock()
        }(server_to_ping.Server)
    }
    wait.Wait()
}

func (p *Pooler) pingServersExecutor(server *datamodels.Server) error {
    srv := server.Ip + ":" + server.Port
    fmt.Println("Pinging " + srv)
    // Dial to server.
    start_p := time.Now()
    conn_ping, err2 := net.Dial("udp", srv)
    if err2 != nil {
        fmt.Println("Error dialing to server " + srv + "!")
        return errors.New("Error dialing to server " + srv + "!")
    }
    // Set deadline, so we won't wait forever.
    ddl_ping := time.Now()
    // This should be enough. Maybe, you should'n run URTrator on modem
    // connections? :)
    ddl_ping = ddl_ping.Add(time.Second * 10)
    conn_ping.SetDeadline(ddl_ping)

    msg_ping := []byte(p.pp + "getinfo")
    conn_ping.Write(msg_ping)

    // UDP Buffer.
    var received_buf_ping []byte = make([]byte, 128)
    // Received buffer.
    var raw_received_ping []byte
    _, err := conn_ping.Read(received_buf_ping)
    if err != nil {
        fmt.Println("PING ERROR")
    }
    raw_received_ping = append(raw_received_ping, received_buf_ping...)
    conn_ping.Close()

    delta := strconv.Itoa(int(time.Since(start_p).Nanoseconds()) / 1000000)
    server.Ping = delta

    return nil
}

func (p *Pooler) UpdateOneServer(server_address string) {
    var wait sync.WaitGroup

    server := Cache.Servers[server_address].Server

    wait.Add(1)
    go func(server *datamodels.Server) {
        defer wait.Done()
        p.UpdateSpecificServer(server)
    }(server)
    wait.Wait()
    p.PingOneServer(server_address)
    Eventer.LaunchEvent("flushServers", map[string]string{})

    Eventer.LaunchEvent("loadAllServers", map[string]string{})
    Eventer.LaunchEvent("loadFavoriteServers", map[string]string{})
    Eventer.LaunchEvent("serversUpdateCompleted", map[string]string{})
}

func (p *Pooler) UpdateServers(servers_type string) {
    var wait sync.WaitGroup

    for _, server := range Cache.Servers {
        if servers_type == "favorites" && server.Server.Favorite != "1" {
            continue
        }
        wait.Add(1)
        go func(server *datamodels.Server) {
            defer wait.Done()
            p.UpdateSpecificServer(server)
        }(server.Server)
    }
    wait.Wait()
    p.PingServers(servers_type)
    Eventer.LaunchEvent("flushServers", map[string]string{})

    Eventer.LaunchEvent("loadAllServers", map[string]string{})
    Eventer.LaunchEvent("loadFavoriteServers", map[string]string{})
    Eventer.LaunchEvent("serversUpdateCompleted", map[string]string{})
}

// Updates information about specific server.
func (p *Pooler) UpdateSpecificServer(server *datamodels.Server) error {
    server_addr := server.Ip + ":" + server.Port
    fmt.Println("Updating server: " + server_addr)

    // Dial to server.
    conn, err1 := net.Dial("udp", server_addr)
    if err1 != nil {
        fmt.Println("Error dialing to server " + server_addr + "!")
        return errors.New("Error dialing to server " + server_addr + "!")
    }

    // Set deadline, so we won't wait forever.
    ddl := time.Now()
    // This should be enough. Maybe, you should'n run URTrator on modem
    // connections? :)
    ddl = ddl.Add(time.Second * 2)
    conn.SetDeadline(ddl)

    msg := []byte(p.pp + "getstatus")
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
    conn.Close()

    // First line is "infoResponse" string, which we should skip by
    // splitting response by "\n".
    received_lines := strings.Split(string(raw_received), "\n")
    // We have server's data!
    if len(received_lines) > 1 {
        srv_config := strings.Split(received_lines[1], "\\")
        // Parse server configuration into passed server's datamodel.
        for i := 0; i < len(srv_config); i = i + 1 {
            if srv_config[i] == "g_modversion" {
                server.Version = srv_config[i + 1]
            }
            if srv_config[i] == "g_gametype" {
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
            if srv_config[i] == "sv_hostname" {
                server.Name = srv_config[i + 1]
            }
            if srv_config[i] == "g_needpass" {
                if srv_config[i + 1] == "0" {
                    server.IsPrivate = "0"
                } else {
                    server.IsPrivate = "1"
                }
            }
            server.ExtendedConfig = received_lines[1]
        }
        if len(received_lines) >= 2 {
            // Here we go, players information.
            players := received_lines[2:]
            var real_players int = 0
            var bots int = 0
            // Calculate players!
            if len(players) == 1 && len(players[0]) > 255 {
                server.Players = "0"
                server.Bots = "0"
            } else {
                // Looks like we have last element to be empty, due to
                // strings.Split() call before.
                for i := range players {
                    // Get slice with data for bots-humans parsing.
                    player_data := strings.Split(string(players[i]), " ")
                    // If slice length isn't equal 3 - this is not what
                    // we want.
                    if len(player_data) != 3 {
                        continue
                    }

                    if player_data[1] == "0" {
                        bots++
                    } else {
                        real_players++
                    }
                }
                //server.Players = strconv.Itoa(len(players) - 1)
                server.Players = strconv.Itoa(real_players)
                server.Bots = strconv.Itoa(bots)
                fmt.Println(server.Players, server.Bots)
            }
            server.PlayersInfo = strings.Join(received_lines[2:], "\\")
        }
    }

    // ToDo: Calculate ping. 0 for now.
    server.Ping = "0"
    return nil
}
