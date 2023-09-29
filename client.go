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
	"time"
)

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc

	conn   *net.TCPConn
	outbox chan []byte

	mac  MACAddress
	name Nickname

	game  *Game
	group *Group
}

/*type ClientList struct {
	data map[uint32]*Client
	mtx  sync.RWMutex
}

func NewClientList() *ClientList {
	return &ClientList{
		data: make(map[uint32]*Client),
	}
}

func (cl *ClientList) get(key uint32) *Client {
	cl.mtx.RLock()

	defer cl.mtx.RUnlock()

	return cl.data[key]
}

func (cl *ClientList) getAll() map[uint32]*Client {
	cl.mtx.RLock()

	defer cl.mtx.RUnlock()

	return cl.data
}

func (cl *ClientList) set(key uint32, value *Client) {
	cl.mtx.Lock()

	cl.data[key] = value

	cl.mtx.Unlock()
}

func (cl *ClientList) del(key uint32) {
	cl.mtx.Lock()

	delete(cl.data, key)

	cl.mtx.Unlock()
}

/*func (cl *GameList) len() int {
	cl.mtx.RLock()

	defer cl.mtx.RUnlock()

	return len(cl.data)
}*/

func (c *Client) listen() {
	buf := make([]byte, 1024)
	for {
		select {
		case <-c.ctx.Done():
			c.disconnect()

			return
		default:
			c.conn.SetReadDeadline(time.Now().Add(time.Second * 5))
			n, err := c.conn.Read(buf)
			if err != nil {
				log.Printf("[ERR] failed to read from %s: %s.", c.conn.RemoteAddr(), err)

				c.cancel()

				continue
			}

			switch Opcode(buf[0]) {
			case Ping:
				continue
			case Login:
				login, err := decodeLogin(buf[1:n])
				if err != nil {
					log.Printf("[WRN] failed to decode login packet from %s: %s.", c.conn.RemoteAddr(), err)
				}

				err = c.handleLogin(login)
				if err != nil {
					log.Printf("[WRN] failed to handle login packet from %s: %s.", c.conn.RemoteAddr(), err)
				}
			case Connect:
				connect, err := decodeConnect(buf[1:n])
				if err != nil {
					log.Printf("[WRN] failed to decode connect packet from %s: %s.", c.conn.RemoteAddr(), err)
				}

				err = c.handleConnect(connect)
				if err != nil {
					log.Printf("[WRN] failed to handle connect packet from %s: %s.", c.conn.RemoteAddr(), err)
				}
			case Disconnect:
				err = c.handleDisconnect()
				if err != nil {
					log.Printf("[WRN] failed to handle disconnect packet from %s: %s.", c.conn.RemoteAddr(), err)
				}
			case Scan:
				err = c.handleScan()
				if err != nil {
					log.Printf("[WRN] failed to handle scan packet from %s: %s.", c.conn.RemoteAddr(), err)
				}
			case Chat:
				chat, err := decodeChat(buf[1:n])
				if err != nil {
					log.Printf("[WRN] failed to decode chat packet from %s: %s.", c.conn.RemoteAddr(), err)
				}

				err = c.handleChat(chat)
				if err != nil {
					log.Printf("[WRN] failed to handle chat packet from %s: %s.", c.conn.RemoteAddr(), err)
				}
			default:
				log.Printf("[WRN] received packet with invalid opcode from %s: %d.", c.conn.RemoteAddr(), buf[0])
			}
		}
	}
}

func (c *Client) messageWriter() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case data := <-c.outbox:
			c.conn.SetWriteDeadline(time.Now().Add(time.Second * 3))
			_, err := c.conn.Write(data)
			if err != nil {
				log.Printf("[ERR] failed to write to %s: %s.", c.conn.RemoteAddr(), err)

				return
			}
		}
	}
}

func (c *Client) disconnect() {
	err := c.conn.Close()
	if err != nil {
		log.Printf("[ERR] failed to close connection from %s: %s.", c.conn.RemoteAddr(), err)
	} else {
		log.Printf("[INF] closed connection from %s.", c.conn.RemoteAddr())
	}

	c.handleDisconnect()
}
