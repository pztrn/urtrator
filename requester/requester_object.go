// URTrator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016-2018, Stanslav N. a.k.a pztrn (or p0z1tr0n) and
// URTrator contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package requester

import (
	// stdlib
	"bytes"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Requester struct {
	// Pooler.
	Pooler *Pooler
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
	r.pp = "\377\377\377\377"
	r.ip_delimiter = 92
	r.Pooler = &Pooler{}
	r.Pooler.Initialize()
}

// Gets all available servers from master server.
// This isn't in pooler, because it have no need to be pooled.
func (r *Requester) getServers() error {
	// Get master server address and port from configuration.
	var master_server string = ""
	var master_server_port string = ""
	master_server_raw, ok := Cfg.Cfg["/servers_updating/master_server"]
	if ok {
		master_server = strings.Split(master_server_raw, ":")[0]
		master_server_port = strings.Split(master_server_raw, ":")[1]
	} else {
		master_server = "master.urbanterror.info"
		master_server_port = "27900"
	}
	fmt.Println("Using master server address: " + master_server + ":" + master_server_port)
	// IP addresses we will compose to return.
	conn, err1 := net.Dial("udp", master_server+":"+master_server_port)
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
		raw_received = append(raw_received, received_buf...)
		fmt.Println("Received " + strconv.Itoa(len(raw_received)) + " bytes")

		if err != nil {
			// A bit hacky - if we have received data length lower or
			// equal to 4k - we have an error. Looks like conn.Read()
			// reads by 4k.
			if len(raw_received) < 4097 {
				fmt.Println("Error dialing to master server!")
				Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "Failed to connect to master server!"})
				return errors.New("Failed to connect to master server!")
			}
			break
		}
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
			Cache.ServersMutex.Lock()
			Cache.Servers[addr].Server.Ip = ip
			Cache.Servers[addr].Server.Port = port
			Cache.ServersMutex.Unlock()
		}
	}

	return nil
}

// Updates information about all available servers from master server and
// parses it to usable format.
func (r *Requester) UpdateAllServers(task bool) {
	fmt.Println("Starting all servers updating procedure...")
	err := r.getServers()
	if err != nil {
		return
	}
	r.Pooler.UpdateServers("all")

	if task {
		Eventer.LaunchEvent("taskDone", map[string]string{"task_name": "Server's autoupdating"})
	}
}

func (r *Requester) UpdateFavoriteServers() {
	fmt.Println("Updating favorites servers...")
	r.Pooler.UpdateServers("favorites")
}

func (r *Requester) UpdateOneServer(server_address string) {
	fmt.Println("Updating server " + server_address)
	r.Pooler.UpdateOneServer(server_address)
}
