#!/bin/bash
if ["$GOPATH" == ""]; then
	export GOPATH="$PWD"
fi

cd $GOPATH/src/github.com/Virepri/Shoraldele
go install
cd $GOPATH/bin
./Shoraldele
