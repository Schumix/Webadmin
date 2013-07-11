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
	"github.com/hoisie/web"
	"github.com/mattn/go-session-manager"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func loadServer(port string) {
	fmt.Print("Starting web server on localhost", port, "...\n")

	manager.OnStart(func(session *session.Session) {
		logger.Printf("Start session(\"%s\")", session.Id)
	})
	manager.OnEnd(func(session *session.Session) {
		logger.Printf("End session(\"%s\")", session.Id)
	})
	manager.SetTimeout(300)

	web.Config.CookieSecret = "7C19QRmwf3mHZ9CPAaPQ0hsWeufKd"
	web.Config.StaticDir = config["WebDir"].(string)

	web.Get("/", func(ctx *web.Context) {
		session := getSession(ctx, manager)
		if session.Value != nil && session.Value.(*User).SuccesLogin {
			session.Value.(*User).SuccesLogin = false
			HomeSuccess(ctx, "Successful login!")
			return
		}

		HandleDefaultFunc(ctx, "index.tpl", "index.tpl", "Home", "home")
	})
	web.Get("/login", func(ctx *web.Context) {
		session := getSession(ctx, manager)
		if session.Value != nil {
			ctx.Redirect(302, "/")
		}

		HandleDefaultFunc(ctx, "login.tpl", "login.tpl", "Login", "login")
	})
	web.Post("/login", func(ctx *web.Context) {
		session := getSession(ctx, manager)
		if session.Value != nil {
			ctx.Redirect(302, "/")
		}

		userid := getParam(ctx, "userid")
		password := getParam(ctx, "password")
		if userid != "" && password != "" {
			// find user
			password = sha1_gen(password)
			st, _ := db.Prepare("select Name, Password from admins where Name = ? and Password = ? and ServerName = 'rizon'")
			r, e := st.Query(userid, password)
			if e != nil {
				logger.Print(e)
				return
			}
			if !r.Next() {
				// not found
				LoginError(ctx, "User not found!")
				return
			}
			var userid, password string
			e = r.Scan(&userid, &password)
			if e != nil {
				logger.Print(e)
				return
			}

			// store User object to sessino
			session.Value = &User{userid, password, true}
			logger.Printf("User \"%s\" login", session.Value.(*User).UserId)
			ctx.Redirect(302, "/")
		}

		if userid == "" && password != "" {
			LoginError(ctx, "Username is missing!")
			return
		}
		if userid != "" && password == "" {
			LoginError(ctx, "Password is missing!")
			return
		}
		if userid == "" && password == "" {
			LoginError(ctx, "Username and password are missing!")
			return
		}
	})
	web.Get("/logout", func(ctx *web.Context) {
		session := getSession(ctx, manager)
		if session.Value != nil {
			// abandon
			logger.Printf("User \"%s\" logout", session.Value.(*User).UserId)
			session.Abandon()
			HomeSuccess(ctx, "Successful logout!")
			return
		}
		ctx.Redirect(302, "/")
	})
	web.Get("/about", func(ctx *web.Context) {
		HandleDefaultFunc(ctx, "about.tpl", "about.tpl", "About", "about")
	})
	web.Get("/stats", func(ctx *web.Context) {
		HandleDefaultFunc(ctx, "stats.tpl", "stats.tpl", "Public Stats", "stats")
	})

	fmt.Print("Done. Serving...\n")
	web.Run(port)
}
