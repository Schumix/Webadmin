#!/bin/bash

sourcedir=`ls -d $PWD`
gopath="$sourcedir/.gopath"
bin="bin"
filename="Webadmin"
golangdir="src"
configdir="configs"
sqldir="sql"

if [ ! -e $gopath ]; then
	mkdir $gopath
fi

export GOPATH=$gopath

go get github.com/mattn/go-sqlite3
go get github.com/hoisie/web
go get github.com/mattn/go-session-manager
go get github.com/Jackneill/gosemver

cd $golangdir
go build -o $filename
cd ..

if [ ! -e $bin ]; then
	mkdir $bin
fi

if [ ! -e $golangdir/$filename ]; then
	echo "Build error!"
	exit
fi

mv $golangdir/$filename $bin/$filename
cp -rf www $bin/www
cp $configdir/config.json $bin/config.json
cp $sqldir/Schumix.db3 $bin/Schumix.db3
