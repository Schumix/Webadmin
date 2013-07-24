#!/bin/bash

sourcedir=`ls -d $PWD`
gopath="$sourcedir/.gopath"
bin="bin"

if [ ! -e $gopath ]; then
	echo "Not build!"
	exit
fi

export GOPATH=$gopath

if [ ! -e $bin ]; then
	echo "Not build!"
	exit
fi

cd $bin
./Webadmin
