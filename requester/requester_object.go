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
    "time"
)

type Requester struct {
    // Pooler.
    Pooler *Pooler
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
    r.Pooler = &Pooler{}
    r.Pooler.Initialize()
}

// Gets all available servers from master server.
// This isn't in pooler, because it have no need to be pooled.
func (r *Requester) getServers() error {
    // IP addresses we will compose to return.
    conn, err1 := net.Dial("udp", r.master_server + ":" + r.master_server_port)
    if err1 != nil {
        fmt.Println("Error dialing to master server!")
        Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "Failed to connect to master server!"})
        return errors.New("Failed to connect to master server!")
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
        addr := ip + ":" + port

        // Check if we already have this server added previously. If so - do nothing.
        _, ok := Cache.Servers[addr]
        if !ok {
            // Create cached server.
            Cache.CreateServer(addr)
            Cache.Servers[addr].Server.Ip = ip
            Cache.Servers[addr].Server.Port = port
        }
    }

    return nil
}

// Updates information about all available servers from master server and
// parses it to usable format.
func (r *Requester) UpdateAllServers() {
    fmt.Println("Starting all servers updating procedure...")
    r.getServers()
    r.Pooler.UpdateServers("all")
}

func (r *Requester) UpdateFavoriteServers() {
    fmt.Println("Updating favorites servers...")
    r.Pooler.UpdateServers("favorites")
}

func (r *Requester) UpdateOneServer(server_address string) {
    fmt.Println("Updating server " + server_address)
    r.Pooler.UpdateOneServer(server_address)
}
