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
)

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

		login.name[i] = data[MACAddressLength+1]
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
