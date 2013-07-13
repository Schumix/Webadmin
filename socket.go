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
	"fmt"
	"net"
	"strconv"
)

const MAX_BUFFER_SIZE = 262144 // 8^6
const PACKET_SEPARATOR = "|;|"

const (
	SCMSG_PACKET_NULL = iota
	CMSG_REQUEST_AUTH
	SMSG_AUTH_APPROVED
	SMSG_AUTH_DENIED
	CMSG_CLOSE_CONNECTION
	SMSG_CLOSE_CONNECTION
)

var conn net.Conn

func connectToSocket(host string) {
	fmt.Print("[SOCKET] Connecting to ", host, "...\n")
	var err error
	conn, err = net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("[SOCKET] Done. ")

	go regConnection()
	listenToSocket()
}

func listenToSocket() {
	fmt.Printf("Listening...\n")
	buffer := make([]byte, MAX_BUFFER_SIZE)
	for {
		n, err := conn.Read(buffer[:])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("-- START PACKET --", n, "bytes", "--")
		fmt.Println(string(buffer[:n]))
		fmt.Println("-- END PACKET --")
	}
}

func regConnection() {
	msg := strconv.Itoa(CMSG_REQUEST_AUTH) + PACKET_SEPARATOR + "schumix webadmin (reg GUID)" + PACKET_SEPARATOR + md5_gen("schumix") + PACKET_SEPARATOR
	fmt.Fprint(conn, msg)
}
