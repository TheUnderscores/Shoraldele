#!/bin/bash
if [ "$GOPATH" != "$PWD" ]; then
	export GOPATH="$PWD"
fi

cd $GOPATH/src/github.com/Virepri/Shoraldele
go install
cd $GOPATH/bin
