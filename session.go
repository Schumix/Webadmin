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
	"github.com/hoisie/web"
	"github.com/mattn/go-session-manager"
	"html/template"
	"strings"
)

type Page struct {
	Title        string
	Body         string
	ProjectName  string
	PageName     string
	SessionValue interface{}
	IsLoggedIn   bool
	Error        bool
	Success      bool
}

type User struct {
	UserId      string
	Password    string
	SuccesLogin bool
}

var manager = session.NewSessionManager(logger)

func getSession(ctx *web.Context, manager *session.SessionManager) *session.Session {
	id, _ := ctx.GetSecureCookie("SessionId")
	session := manager.GetSessionById(id)
	ctx.SetSecureCookie("SessionId", session.Id, int64(manager.GetTimeout()))
	ctx.SetHeader("Pragma", "no-cache", true)
	return session
}

func getParam(ctx *web.Context, name string) string {
	value, found := ctx.Params[name]
	if found {
		return strings.Trim(value, " ")
	}
	return ""
}

func HandleDefaultFunc(ctx *web.Context, filename string, filelocation string, title string, pagename string) {
	t, _ := template.New(filename).Funcs(
		template.FuncMap{
			"eq": func(a, b string) bool {
				return a == b
			},
		}).ParseFiles(config["WebDir"].(string)+"/template/header.tpl", config["WebDir"].(string)+"/template/menu.tpl", config["WebDir"].(string)+"/"+filelocation, config["WebDir"].(string)+"/template/footer.tpl")
	p := PageSettings(ctx, title, pagename)
	t.Execute(ctx, p)
}

func HandleFunc(ctx *web.Context, page Page, filename string, filelocation string) {
	t, _ := template.New(filename).Funcs(
		template.FuncMap{
			"eq": func(a, b string) bool {
				return a == b
			},
		}).ParseFiles(config["WebDir"].(string)+"/template/header.tpl", config["WebDir"].(string)+"/template/menu.tpl", config["WebDir"].(string)+"/"+filelocation, config["WebDir"].(string)+"/template/footer.tpl")
	p := page
	t.Execute(ctx, p)
}

func PageSettings(ctx *web.Context, title string, pagename string) Page {
	return Page{Title: title + " - " + config["Title"].(string), ProjectName: config["ProjectName"].(string), PageName: pagename, SessionValue: getSession(ctx, manager).Value, IsLoggedIn: IsLoggedIn(ctx)}
}

func LoginError(ctx *web.Context, message string) {
	p := Page{Title: "Login" + " - " + config["Title"].(string), Body: message, ProjectName: config["ProjectName"].(string), PageName: "login", SessionValue: nil, Error: true}
	HandleFunc(ctx, p, "login.tpl", "login.tpl")
}

func HomeSuccess(ctx *web.Context, message string) {
	session := getSession(ctx, manager)
	p := Page{Title: "Home" + " - " + config["Title"].(string), Body: message, ProjectName: config["ProjectName"].(string), PageName: "home", SessionValue: session.Value, Success: true, IsLoggedIn: IsLoggedIn(ctx)}
	HandleFunc(ctx, p, "index.tpl", "index.tpl")
}

func IsLoggedIn(ctx *web.Context) bool {
	session := getSession(ctx, manager)

	if session.Value != nil {
		return true
	}
	return false
}
