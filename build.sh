#!/bin/bash
if [ "$GOPATH" == "" ]; then
	export GOPATH="$PWD"
fi

last_path="$(pwd)"

cd $GOPATH/src/github.com/Virepri/Shoraldele
go install
cd $GOPATH/bin
./Shoraldele
cd "$last_path"
