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
	"database/sql"
	"fmt"
	_ "github.com/Schumix/go-sqlite3"
)

var db *sql.DB

func connectToSql() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", config["SQLiteFile"])
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
