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
