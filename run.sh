#!/bin/bash

sourcedir=`ls -d $PWD`
gopath="$sourcedir/.gopath"
bin="bin"

if [ ! -e $gopath ]; then
	echo "$gopath not found! Run build.sh."
	exit
fi

export GOPATH=$gopath

if [ ! -e $bin ]; then
	echo "$bin not found! Run build.sh."
	exit
fi

cd $bin
./Webadmin
