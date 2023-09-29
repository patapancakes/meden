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
	"net"
)

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

// Decoders

func decodeLogin(data []byte) (LoginPacketC2S, error) {
	var login LoginPacketC2S

	if len(data) < MACAddressLength+NicknameLength+ProductCodeLength {
		return LoginPacketC2S{}, fmt.Errorf("invalid packet size: expected at least %d, got %d", MACAddressLength+NicknameLength+ProductCodeLength, len(data))
	}

	for i := range login.mac {
		login.mac[i] = data[i]
	}

	for i := range login.name {
		if data[MACAddressLength+i] == 0 {
			break
		}

		login.name[i] = data[MACAddressLength+i]
	}

	for i := range login.game {
		if data[MACAddressLength+NicknameLength+i] == 0 {
			break
		}

		login.game[i] = data[MACAddressLength+NicknameLength+i]
	}

	return login, nil
}

func decodeConnect(data []byte) (ConnectPacketC2S, error) {
	var connect ConnectPacketC2S

	if len(data) < GroupNameLength {
		return ConnectPacketC2S{}, fmt.Errorf("invalid packet size: expected at least %d, got %d", GroupNameLength, len(data))
	}

	for i := range connect.group {
		if data[i] == 0 {
			break
		}

		connect.group[i] = data[i]
	}

	return connect, nil
}

func decodeChat(data []byte) (ChatPacketC2S, error) {
	var chat ChatPacketC2S

	if len(data) < ChatMessageLength {
		return ChatPacketC2S{}, fmt.Errorf("invalid packet size: expected at least %d, got %d", ChatMessageLength, len(data))
	}

	for i := range chat.message {
		if data[i] == 0 {
			break
		}

		chat.message[i] = data[i]
	}

	return chat, nil
}

// Encoders

func encodeConnect(data ConnectPacketS2C) []byte {
	nameTruncated := cstrTruncate(data.name[:])

	connect := make([]byte, 0, 1+len(nameTruncated)+MACAddressLength+4)

	connect = append(connect, uint8(Connect))
	connect = append(connect, nameTruncated...)
	connect = append(connect, data.mac[:]...)
	connect = append(connect, data.ip.To4()...)

	return connect
}

func encodeDisconnect(data DisconnectPacketS2C) []byte {
	disconnect := make([]byte, 0, 1+4)

	disconnect = append(disconnect, uint8(Disconnect))
	disconnect = append(disconnect, data.ip.To4()...)

	return disconnect
}

func encodeScan(data ScanPacketS2C) []byte {
	scan := make([]byte, 0, 1+GroupNameLength+MACAddressLength)

	scan = append(scan, uint8(Scan))
	scan = append(scan, data.group[:]...)
	scan = append(scan, data.mac[:]...)

	return scan
}

func encodeScanComplete() []byte {
	scanComplete := make([]byte, 0, 1)

	scanComplete = append(scanComplete, uint8(ScanComplete))

	return scanComplete
}

func encodeConnectBSSID(data ConnectBSSIDPacketS2C) []byte {
	connectBSSID := make([]byte, 0, 1+MACAddressLength)

	connectBSSID = append(connectBSSID, uint8(ConnectBSSID))
	connectBSSID = append(connectBSSID, data.mac[:]...)

	return connectBSSID
}

func encodeChat(data ChatPacketS2C) []byte {
	messageTruncated := cstrTruncate(data.message[:])
	nameTruncated := cstrTruncate(data.name[:])

	chat := make([]byte, 0, 1+len(messageTruncated)+len(nameTruncated))

	chat = append(chat, uint8(Chat))
	chat = append(chat, messageTruncated...)
	chat = append(chat, nameTruncated...)

	return chat
}
