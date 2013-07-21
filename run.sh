#!/bin/bash

sourcedir=`ls -d $PWD`
gopath="$sourcedir/.gopath"

if [ ! -e $gopath ]; then
	mkdir $gopath
fi

export GOPATH=$gopath

go get github.com/mattn/go-sqlite3
go get github.com/hoisie/web
go get github.com/mattn/go-session-manager
go get github.com/Jackneill/gosemver

go run *.go
