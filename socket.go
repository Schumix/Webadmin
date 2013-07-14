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
	"strings"
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
	CMSG_PING
	SMSG_PING
	CMSG_PONG
	SMSG_PONG
	CMSG_SCHUMIX_VERSION
	SMSG_SCHUMIX_VERSION
)

var conn net.Conn
var isconnected bool

func connectToSocket(host string) {
	fmt.Print("[SOCKET] Connecting to ", host, "...\n")
	var err error
	conn, err = net.Dial("tcp", host)
	if err != nil {
		isconnected = false
		fmt.Println(err)
		fmt.Println("[SOCKET] Fail.")
	} else {
		isconnected = true
		fmt.Print("[SOCKET] Done. ")
		go regConnection()
		listenToSocket()
	}
}

func listenToSocket() {
	fmt.Printf("Listening...\n")
	buffer := make([]byte, MAX_BUFFER_SIZE)
	for {
		if !isconnected {
			break
		}

		n, err := conn.Read(buffer[:])
		if err != nil {
			fmt.Println(err)
		}
		handlePacket(string(buffer[:n]), n)
	}
}

func handlePacket(data string, size int) {
	// separate packet to its elements
	packet := strings.Split(data, PACKET_SEPARATOR)

	fmt.Println("-- START PACKET --", size, "bytes", "--")
	fmt.Print("-- Opcode: ", packet[0], " -- ")
	opcode, _ := strconv.Atoi(packet[0])
	switch opcode {
	case SMSG_AUTH_APPROVED:
		fmt.Print("Auth request approved.")
	case SMSG_AUTH_DENIED:
		isconnected = false
		fmt.Print("Auth request denied.")
	case SMSG_CLOSE_CONNECTION:
		isconnected = false
		fmt.Print("Server sent closing signal. Connection closed.")
		conn.Close()
	case SMSG_PING:
		sendPong()
	case SMSG_PONG:
		//sendPong()
		break
	case SMSG_SCHUMIX_VERSION:
		break
	default:
		fmt.Print("Unknown opcode.")
	}
	fmt.Println(" --")
	fmt.Println(packet[1:])
	fmt.Println("-- END PACKET --")
}

func shutdownSocket() {
	if isconnected {
		fmt.Println("Shutting down socket connection...")
		sendCloseSignal()
		conn.Close()
	}
}

func sendPing() {
	msg := strconv.Itoa(CMSG_PING) + PACKET_SEPARATOR
	fmt.Fprint(conn, msg)
}

func sendPong() {
	msg := strconv.Itoa(CMSG_PONG) + PACKET_SEPARATOR
	fmt.Fprint(conn, msg)
}

func sendCloseSignal() {
	msg := strconv.Itoa(CMSG_CLOSE_CONNECTION) + PACKET_SEPARATOR + "uh. stomachache. shutting down for now." + PACKET_SEPARATOR
	fmt.Fprint(conn, msg)
}

func regConnection() {
	msg := strconv.Itoa(CMSG_REQUEST_AUTH) + PACKET_SEPARATOR + "schumix webadmin (reg GUID)" + PACKET_SEPARATOR + md5_gen("schumix") + PACKET_SEPARATOR
	fmt.Fprint(conn, msg)
}

func requestVersion() {
	msg := strconv.Itoa(CMSG_SCHUMIX_VERSION) + PACKET_SEPARATOR
	fmt.Fprint(conn, msg)
}
