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

import "sync"

type Group struct {
	clients []*Client
}

type GroupList struct {
	data map[GroupName]*Group
	mtx  sync.RWMutex
}

func NewGroupList() *GroupList {
	return &GroupList{data: make(map[GroupName]*Group)}
}

func (gl *GroupList) get(key GroupName) *Group {
	gl.mtx.RLock()

	defer gl.mtx.RUnlock()

	return gl.data[key]
}

func (gl *GroupList) getAll() map[GroupName]*Group {
	gl.mtx.RLock()

	defer gl.mtx.RUnlock()

	return gl.data
}

func (gl *GroupList) set(key GroupName, value *Group) {
	gl.mtx.Lock()

	gl.data[key] = value

	gl.mtx.Unlock()
}

/*func (gl *GroupList) del(key GroupName) {
	gl.mtx.Lock()

	delete(gl.data, key)

	gl.mtx.Unlock()
}

func (gl *GroupList) len() int {
	gl.mtx.RLock()

	defer gl.mtx.RUnlock()

	return len(gl.data)
}*/
