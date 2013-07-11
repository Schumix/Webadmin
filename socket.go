/*
 * This file is part of Schumix Webadmin.
 *
 * Copyright (C) 2013 Schumix Team <http://schumix.eu/>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
)

const MAX_BUFFER_SIZE = 262144

func connectToSocket() {
	origin := "http://localhost/"
	url := "ws://localhost:36200/ws"

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		fmt.Println(err)
	}

	var msg = make([]byte, MAX_BUFFER_SIZE)
	if n, err := ws.Read(msg); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received: %s.\n", msg[:n])
}
