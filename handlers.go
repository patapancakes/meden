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
	"fmt"
	"log"
	"regexp"
)

var validGroupName = regexp.MustCompile("^[a-zA-Z0-9]+\x00$").Match

func (c *Client) handleLogin(login LoginPacketC2S) error {
	if c.game != nil {
		return fmt.Errorf("client already logged in")
	}

	if login.mac == [6]uint8{0x00, 0x00, 0x00, 0x00, 0x00, 0x00} || login.mac == [6]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF} {
		return fmt.Errorf("invalid mac address: %X", login.mac)
	}

	var baseProductCode ProductCode
	if link, ok := productLinks[cstrToString(login.game[:])]; ok {
		for i := range baseProductCode {
			baseProductCode[i] = link[i]
		}
	} else {
		baseProductCode = login.game
	}

	gameName, ok := productNames[cstrToString(baseProductCode[:])]
	if !ok {
		return fmt.Errorf("invalid product code: %s", baseProductCode)
	}

	game := games.get(baseProductCode)
	if game == nil {
		game = &Game{groups: NewGroupList()}

		games.set(baseProductCode, game)
	}

	c.game = game

	c.mac = login.mac
	c.name = login.name

	log.Printf("[INF] %s (MAC: %X) started game \"%s\".", cstrToString(login.name[:]), login.mac, gameName)

	return nil
}

func (c *Client) handleConnect(connect ConnectPacketC2S) error {
	if c.game == nil {
		return fmt.Errorf("client not in game")
	}

	if c.group != nil {
		return fmt.Errorf("client already in group")
	}

	if !validGroupName(connect.group[:]) {
		return fmt.Errorf("invalid group name")
	}

	group := c.game.groups.get(connect.group)
	if group == nil {
		group = &Group{}

		c.game.groups.set(connect.group, group)
	}

	bssid := c.mac

	payload := encodeConnect(ConnectPacketS2C{
		name: c.name,
		mac:  c.mac,
		ip:   addrToIP(c.conn.RemoteAddr()),
	})

	for i, client := range group.clients {
		if client == nil {
			continue
		}

		if i == 0 {
			bssid = client.mac
		}

		client.outbox <- payload

		c.outbox <- encodeConnect(ConnectPacketS2C{
			name: client.name,
			mac:  client.mac,
			ip:   addrToIP(client.conn.RemoteAddr()),
		})
	}

	c.group.clients = append(c.group.clients, c)

	c.group = group

	c.outbox <- encodeConnectBSSID(ConnectBSSIDPacketS2C{
		mac: bssid,
	})

	log.Printf("[INF] %s (MAC: %X) joined group \"%s\".", cstrToString(c.name[:]), c.mac, cstrToString(connect.group[:]))

	return nil
}

func (c *Client) handleDisconnect() error {
	if c.game == nil {
		return fmt.Errorf("client not in game")
	}

	if c.group == nil {
		return fmt.Errorf("client not in group")
	}

	c.cancel()

	payload := encodeDisconnect(DisconnectPacketS2C{
		ip: addrToIP(c.conn.RemoteAddr()),
	})

	for _, client := range c.group.clients {
		if client == c {
			continue
		}

		client.outbox <- payload
	}

	clients.del(addrToUint32(c.conn.RemoteAddr()))

	for i, client := range c.group.clients {
		if client == c {
			c.group.clients = append(c.group.clients[:i], c.group.clients[i+1:]...)

			break
		}
	}

	c.game.groups.mtx.Lock()

	if len(c.group.clients) == 0 {
		for groupName, group := range c.game.groups.data {
			if group == c.group {
				delete(c.game.groups.data, groupName)

				break
			}
		}
	}

	c.game.groups.mtx.Unlock()

	games.mtx.Lock()

	if len(c.game.groups.data) == 0 {
		for productCode, game := range games.data {
			if game == c.game {
				delete(games.data, productCode)

				break
			}
		}
	}

	games.mtx.Unlock()

	log.Printf("[INF] %s (MAC: %X) disconnected.", cstrToString(c.name[:]), c.mac)

	return nil
}

func (c *Client) handleScan() error {
	if c.game == nil {
		return fmt.Errorf("client not in game")
	}

	if c.group != nil {
		return fmt.Errorf("client already in group")
	}

	for name, group := range c.game.groups.getAll() {
		c.outbox <- encodeScan(ScanPacketS2C{
			group: name,
			mac:   group.clients[0].mac,
		})
	}

	c.outbox <- encodeScanComplete()

	log.Printf("[INF] %s (MAC: %X) scanned groups.", cstrToString(c.name[:]), c.mac)

	return nil
}

func (c *Client) handleChat(chat ChatPacketC2S) error {
	if c.game == nil {
		return fmt.Errorf("client not in game")
	}

	if c.group == nil {
		return fmt.Errorf("client not in group")
	}

	for _, client := range c.group.clients {
		if client == c {
			continue
		}

		client.outbox <- encodeChat(ChatPacketS2C{
			message: chat.message,
			name:    c.name,
		})
	}

	log.Printf("[INF] %s (MAC: %X) sent message \"%s\".", cstrToString(c.name[:]), c.mac, cstrToString(chat.message[:]))

	return nil
}
