#!/bin/bash

sourcedir=`ls -d $PWD`
gopath="$sourcedir/.gopath"
bin="$sourcedir/bin"

if [ ! -e $gopath ]; then
	mkdir $gopath
fi

export GOPATH=$gopath

go get github.com/mattn/go-sqlite3
go get github.com/hoisie/web
go get github.com/mattn/go-session-manager
go get github.com/Jackneill/gosemver

go build

if [ ! -e $bin ]; then
	mkdir $bin
fi

mv Webadmin $bin/Webadmin
cp -rf www $bin/www
cp config.json $bin/config.json
cp Schumix.db3 $bin/Schumix.db3
