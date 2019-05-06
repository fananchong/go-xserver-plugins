#!/bin/bash

set -e

CUR_DIR=$PWD
SRC_DIR=$PWD
BIN_DIR=$PWD/bin
CONF_DIR=$PWD/config
FLAGS=-race

export GOPROXY=https://goproxy.io
cd $SRC_DIR
go generate ./...
cd $SRC_DIR
go vet ./...
echo "start build ..."
cd $SRC_DIR
plugins=`ls -l | grep '^d' | awk '{print $9}' | grep -v 'internal\|bin'` 
for plugin_name in $plugins; do
    go build $FLAGS -buildmode=plugin -o $BIN_DIR/$plugin_name.so ./$plugin_name;
done
echo "done"

cd $CUR_DIR

exit 0

