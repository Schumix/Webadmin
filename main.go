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
	"os"
	"os/signal"
)

var shutdown bool

func main() {
	shutdown = false
	go beforeShutdown()
	loadConfig()
	db = connectToSql()
	defer db.Close()
	go connectToSocket("localhost:36200")
	loadServer(":" + config["Port"].(string))
}

func beforeShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Got signal:", s)
	shutdown = true
	shutdownSocket()
	os.Exit(1)
}
