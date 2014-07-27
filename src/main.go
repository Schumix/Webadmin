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
	"github.com/Schumix/Protocol"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// connection states
const (
	STATE_OPENING = iota
	STATE_OPEN
	STATE_CLOSING
	STATE_CLOSED
)

var Timeout = "30m"
var connectionState = make(chan int)
var reConnectState = make(chan bool)
var state = STATE_CLOSED
var host string
var conn net.Conn

var admin_acc = gin.Accounts{"admin": "admin_pass"}

func setupProtocol() {
	protocol.UseConn(conn)
}

func connectToClient(addr string) {
	connectionState <- STATE_OPENING

	host = addr
	fmt.Print("[SOCKET] Connecting to ", addr, "...\n")

	var err error
	conn, err = net.Dial("tcp", addr)

	setupProtocol()

	if err != nil {
		connectionState <- STATE_CLOSED

		fmt.Println(err)
		fmt.Println("[SOCKET] Fail.")
	} else {

		fmt.Println("[SOCKET] Done connecting, registring...")

		protocol.RegConnection()

		connectionState <- STATE_OPEN

		fmt.Println("[SOCKET] Registered. Listening...")

		listenToSocket()

		defer conn.Close()
	}
}

func listenToSocket() {
	buffer := make([]byte, protocol.MAX_BUFFER_SIZE)
	for {
		if state == STATE_CLOSED {
			break
		}
		n, err := conn.Read(buffer[:])
		if err != nil {
			fmt.Println(err)
		}
		if err == io.EOF {
			fmt.Println("[SOCKET] Remote server closed connection.")
			connectionState <- STATE_CLOSED
			break
		}
		go handlePacket(string(buffer[:n]), n)
	}
}

func handlePacket(data string, size int) {
	// separate packet to its elements
	packet := strings.Split(data, protocol.PACKET_SEPARATOR)
	if packet[0] == "" {
		fmt.Print("Empty packet.")
		return
	}
	fmt.Print("-- START PACKET -- ", size, " bytes")
	fmt.Print(" -- Opcode: ", packet[0], " -- ")
	opcode, _ := strconv.Atoi(packet[0])
	switch opcode {
	case protocol.SMSG_AUTH_APPROVED:
		fmt.Println("Auth request approved.")
		protocol.RequestVersion()
	case protocol.SMSG_AUTH_DENIED:
		fmt.Println("Auth request denied.")
	case protocol.SMSG_CLOSE_CONNECTION:
		connectionState <- STATE_CLOSING
		fmt.Println("Server sent closing signal. Connection closed.")
		conn.Close()
	case protocol.SMSG_PING:
		fmt.Println("SMSG_PING")
		protocol.SendPong()
	case protocol.SMSG_PONG:
		fmt.Println("SMSG_PONG")
	case protocol.SMSG_SCHUMIX_VERSION:
		protocol.CheckVersion(packet[1])
	default:
		fmt.Println("Unknown opcode.")
	}
	fmt.Println(packet[1:])
	fmt.Println("-- END PACKET --")
}

func stateWatcher() {
	for which := range connectionState {
		switch which {
		case STATE_OPEN:
			state = STATE_OPEN
		case STATE_OPENING:
			state = STATE_OPENING
		case STATE_CLOSING:
			state = STATE_CLOSING
		case STATE_CLOSED:
			state = STATE_CLOSED
			reConnectState <- true
		}
	}
}

func reConnectWatcher() {
	for which := range reConnectState {
		switch which {
		case true:
			fmt.Println("[SOCET] Reconnecting in ", Timeout)
			dur, err := time.ParseDuration(Timeout)
			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(dur)
			fmt.Println("[SOCKET] Reconnecting...")
			connectToClient(host)
		}
	}
}

func shutdownSocket() {
	if conn != nil && state == STATE_OPEN || state == STATE_OPENING {
		state = STATE_CLOSING
		fmt.Println("[SOCKET] Shutting down socket connection...")
		protocol.SendCloseSignal()
		conn.Close()
	}
}

func main() {
	// get most out of multiple CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())

	go beforeShutdown(shutdownSocket)
	go stateWatcher()
	go connectToClient(":36200")
	go reConnectWatcher()

	r := gin.Default()
	r.Static("/static", "www/static")
	r.LoadHTMLFiles("www/template/footer.tpl", "www/template/header.tpl",
		"www/template/menu.tpl", "www/index.tpl", "www/about.tpl", "www/login.tpl",
		"www/stats.tpl", "www/status.tpl", "www/status_build.tpl")

	r.GET("/", GET_index)
	r.GET("/about", GET_about)
	r.POST("/logout", POST_logout)

	authrequired_stats := r.Group("/stats", gin.BasicAuth(admin_acc))
	authrequired_stats.GET("/", GET_stats)

	authrequired_status := r.Group("/status", gin.BasicAuth(admin_acc))
	authrequired_status.GET("/", GET_status)

	authrequired_status_build := r.Group("/status-build", gin.BasicAuth(admin_acc))
	authrequired_status_build.GET("/", GET_status_build)

	r.Run(":8080")
}

func GET_index(ctx *gin.Context) {
	ctx.HTML(200, "index.tpl", nil)
}

func GET_about(ctx *gin.Context) {
	ctx.HTML(200, "about.tpl", nil)
}

func GET_stats(ctx *gin.Context) {
	user, err := ctx.Get(gin.AuthUserKey)
	if err != nil {
		fmt.Println(err)
	} else {
		if _, ok := admin_acc[user.(string)]; ok {
			ctx.HTML(200, "stats.tpl", gin.H{"User": user.(string)})
		} else {
			ctx.HTML(200, "stats.tpl", gin.H{"User": "Unauthorized"})
		}
	}
}

func GET_status(ctx *gin.Context) {
	user, err := ctx.Get(gin.AuthUserKey)
	if err != nil {
		fmt.Println(err)
	} else {
		if _, ok := admin_acc[user.(string)]; ok {
			ctx.HTML(200, "status.tpl", gin.H{"User": user.(string)})
		} else {
			ctx.HTML(200, "status.tpl", gin.H{"User": "Unauthorized"})
		}
	}
}

func GET_status_build(ctx *gin.Context) {
	user, err := ctx.Get(gin.AuthUserKey)
	if err != nil {
		fmt.Println(err)
	} else {
		if _, ok := admin_acc[user.(string)]; ok {
			ctx.HTML(200, "status_build.tpl", gin.H{"User": user.(string)})
		} else {
			ctx.HTML(200, "status_build.tpl", gin.H{"User": "Unauthorized"})
		}
	}
}

func POST_logout(ctx *gin.Context) {
	ctx.Set(gin.AuthUserKey, nil)
	ctx.Writer.Header().Set("Location", "/stats")
	ctx.Abort(302)
}

func beforeShutdown(callback func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Got signal:", s)

	// for protocol cleanup/stuff
	callback()

	os.Exit(1)
}
