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
	"bufio"
	"fmt"
	"net"
)

//const MAX_BUFFER_SIZE = 262144

func connectToSocket(host string) {
	fmt.Print("[SOCKET] Connecting to", host, "...\n")
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("[SOCKET] Done. Listening...\n")
	//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		fmt.Println("----------")
		fmt.Println(status)
	}
	fmt.Println(status)
}
