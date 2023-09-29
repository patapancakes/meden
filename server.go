/*
	Copyright (C) 2023  patapancakes <maru@myyahoo.com>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"context"
	"log"
	"net"
)

var (
	games   = NewGameList()
	//clients = NewClientList()
)

func main() {
	log.Printf("[INF] Meden started.")

	listener, err := net.ListenTCP("tcp4", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 27312})
	if err != nil {
		log.Fatalf("[FTL] failed to create matchmaking listener: %s.", err)
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("[ERR] failed to accept connection from %s: %s.", conn.RemoteAddr(), err)

			continue
		}

		client := &Client{
			conn:   conn,
			outbox: make(chan []byte, 16),
		}

		client.ctx, client.cancel = context.WithCancel(context.Background())

		//clients.set(addrToUint32(conn.RemoteAddr()), client)

		go client.messageWriter()
		go client.listen()

		log.Printf("[INF] accepted connection from %s.", conn.RemoteAddr())
	}
}

/*func addrToUint32(addr net.Addr) uint32 {
	return binary.LittleEndian.Uint32(addrToIP(addr))
}*/

func addrToIP(addr net.Addr) net.IP {
	host, _, _ := net.SplitHostPort(addr.String())

	return net.ParseIP(host)
}

func cstrToString(data []byte) string {
	var output string

	for _, char := range data {
		if char == 0 {
			break
		}

		output += string(char)
	}

	return output
}

func cstrTruncate(data []byte) []byte {
	var output []byte

	for _, char := range data {
		if char == 0 {
			output = append(output, char)

			break
		}

		output = append(output, char)
	}

	return output
}
