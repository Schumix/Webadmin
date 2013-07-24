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
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func sha1_gen(data string) string {
	chiperer := sha1.New()
	chiperer.Write([]byte(data))
	bs := chiperer.Sum(nil)
	return hex.EncodeToString(bs)
}

func md5_gen(data string) string {
	chiperer := md5.New()
	chiperer.Write([]byte(data))
	bs := chiperer.Sum(nil)
	return hex.EncodeToString(bs)
}
