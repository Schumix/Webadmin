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
	"github.com/Jackneill/gosemver"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

const MAX_BUFFER_SIZE = 262144 // 8^6 byte = 256 MB
const PACKET_SEPARATOR = "|;|"

const (
	SCMSG_PACKET_NULL     = 0x0
	CMSG_REQUEST_AUTH     = 0x01
	SMSG_AUTH_APPROVED    = 0x02
	SMSG_AUTH_DENIED      = 0x03
	CMSG_CLOSE_CONNECTION = 0x04
	SMSG_CLOSE_CONNECTION = 0x05
	CMSG_PING             = 0x06
	SMSG_PING             = 0x07
	CMSG_PONG             = 0x08
	SMSG_PONG             = 0x09
	CMSG_SCHUMIX_VERSION  = 0x10
	SMSG_SCHUMIX_VERSION  = 0x11
)

var conn net.Conn
var connectionState = make(chan bool)
var isConnected bool
var mHost string

func connectToSocket(host string) {
	mHost = host
	fmt.Print("[SOCKET] Connecting to ", host, "...\n")
	var err error
	conn, err = net.Dial("tcp", host)
	go reConnect()
	if err != nil {
		connectionState <- false
		fmt.Println(err)
		fmt.Println("[SOCKET] Fail.")
	} else {
		connectionState <- true
		fmt.Print("[SOCKET] Done. ")
		go regConnection()
		listenToSocket()
		defer conn.Close()
	}
}

func listenToSocket() {
	fmt.Printf("Listening...\n")
	buffer := make([]byte, MAX_BUFFER_SIZE)
	for {
		if !isConnected || shutdown {
			break
		}
		n, err := conn.Read(buffer[:])
		if err != nil {
			fmt.Println(err)
		}
		if err == io.EOF {
			fmt.Println("[SOCKET] Remote server closed connection.")
			connectionState <- false
			break
		}
		handlePacket(string(buffer[:n]), n)
	}
}

func handlePacket(data string, size int) {
	// separate packet to its elements
	packet := strings.Split(data, PACKET_SEPARATOR)
	if packet[0] == "" {
		fmt.Print("Empty packet.")
		return
	}
	fmt.Print("-- START PACKET -- ", size, " bytes")
	fmt.Print(" -- Opcode: ", packet[0], " -- ")
	opcode, _ := strconv.Atoi(packet[0])
	switch opcode {
	case SMSG_AUTH_APPROVED:
		fmt.Println("Auth request approved.")
		requestVersion()
	case SMSG_AUTH_DENIED:
		connectionState <- false
		fmt.Println("Auth request denied.")
	case SMSG_CLOSE_CONNECTION:
		connectionState <- false
		fmt.Println("Server sent closing signal. Connection closed.")
		conn.Close()
	case SMSG_PING:
		fmt.Println("SMSG_PING")
		sendPong()
	case SMSG_PONG:
		fmt.Println("SMSG_PONG")
	case SMSG_SCHUMIX_VERSION:
		checkVersion(packet[1])
	default:
		fmt.Println("Unknown opcode.")
	}
	fmt.Println(packet[1:])
	fmt.Println("-- END PACKET --")
}

func reConnect() {
	for which := range connectionState {
		if !which {
			isConnected = false
			dur, err := time.ParseDuration(config["Timeout"])
			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(dur)
			fmt.Println("[SOCET] Reconnecting...")
			go connectToSocket(mHost)
		} else {
			isConnected = true
		}
	}
}

func shutdownSocket() {
	if isConnected {
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
	msg := strconv.Itoa(CMSG_CLOSE_CONNECTION) + PACKET_SEPARATOR +
		"uh. stomachache. shutting down for now." + PACKET_SEPARATOR
	fmt.Fprint(conn, msg)
}

func regConnection() {
	msg := strconv.Itoa(CMSG_REQUEST_AUTH) + PACKET_SEPARATOR +
		"schumix webadmin (reg GUID)" + PACKET_SEPARATOR + md5_gen("schumix") + PACKET_SEPARATOR
	fmt.Fprint(conn, msg)
}

func requestVersion() {
	msg := strconv.Itoa(CMSG_SCHUMIX_VERSION) + PACKET_SEPARATOR
	fmt.Fprint(conn, msg)
}

func checkVersion(ver string) {
	res := gosemver.Compare(MIN_SCHUMIX_VERSION, ver)
	if res == 0 || res == 2 {
		fmt.Println("Version check OK.")
		fmt.Println("[VERSION] Webadmin:", VERSION, "Min Schumix:",
			MIN_SCHUMIX_VERSION, "Schumix connected:", ver)
	} else {
		fmt.Println("Schumix version is too low...")
		shutdownSocket()
	}
}
