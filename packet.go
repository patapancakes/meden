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

import "net"

type Opcode uint8

const (
	Ping Opcode = iota
	Login
	Connect
	Disconnect
	Scan
	ScanComplete
	ConnectBSSID
	Chat
)

const (
	MACAddressLength = 6

	ProductCodeLength = 9

	GroupNameLength = 8
	NicknameLength  = 128

	ChatMessageLength = 64
)

type MACAddress [MACAddressLength]uint8

type ProductCode [ProductCodeLength]byte

type GroupName [GroupNameLength]byte
type Nickname [NicknameLength]byte

type ChatMessage [ChatMessageLength]byte

// C2S Packets

// (01) C2S Login Packet
type LoginPacketC2S struct {
	mac  MACAddress
	name Nickname
	game ProductCode
}

// (02) C2S Connect Packet
type ConnectPacketC2S struct {
	group GroupName
}

// (07) C2S Chat Packet
type ChatPacketC2S struct {
	message ChatMessage
}

// S2C Packets

// (02) S2C Connect Packet
type ConnectPacketS2C struct {
	name Nickname
	mac  MACAddress
	ip   net.IP
}

// (03) S2C Disconnect Packet
type DisconnectPacketS2C struct {
	ip net.IP
}

// (04) S2C Scan Packet
type ScanPacketS2C struct {
	group GroupName
	mac   MACAddress
}

// (06) S2C Connect BSSID Packet
type ConnectBSSIDPacketS2C struct {
	mac MACAddress
}

// (07) S2C Chat Packet
type ChatPacketS2C struct {
	message ChatMessage
	name    Nickname
}
