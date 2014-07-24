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
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
)

const VERSION = "0.2.2"
const MIN_SCHUMIX_VERSION = "4.1.6"

var shutdown bool

func main() {
	shutdown = false
	go beforeShutdown()

	r := gin.Default()
	r.Static("/static", "www/static")
	r.LoadHTMLFiles("www/template/footer.tpl", "www/template/header.tpl",
		"www/template/menu.tpl", "www/index.tpl", "www/about.tpl", "www/login.tpl",
		"www/stats.tpl", "www/status.tpl", "www/status_build.tpl")

	r.GET("/", GET_index)
	r.GET("/about", GET_about)
	r.GET("/stats", GET_stats)
	r.GET("/status", GET_status)
	r.GET("/status-build", GET_status_build)
	r.GET("/login", GET_login)

	r.Run(":8080")
}

func GET_index(ctx *gin.Context) {
	ctx.HTML(200, "index.tpl", nil)
}

func GET_about(ctx *gin.Context) {
	ctx.HTML(200, "about.tpl", nil)
}

func GET_stats(ctx *gin.Context) {
	ctx.HTML(200, "stats.tpl", nil)
}

func GET_status(ctx *gin.Context) {
	ctx.HTML(200, "status.tpl", nil)
}

func GET_status_build(ctx *gin.Context) {
	ctx.HTML(200, "status_build.tpl", nil)
}

func GET_login(ctx *gin.Context) {
	ctx.HTML(200, "login.tpl", nil)
}

func beforeShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Got signal:", s)
	shutdown = true
	os.Exit(1)
}
