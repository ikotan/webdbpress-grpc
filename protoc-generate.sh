#!/usr/bin/env bash

PROTO_DIR="./downloader/proto/"
PLUGINS="grpc"
DEST_PATH="./downloader/proto/"
GO_OUT="plugins=$PLUGINS:$DEST_PATH"
PROTO_FILE_PATH="./downloader/proto/file.proto"

protoc --proto_path $PROTO_DIR --go_out=$GO_OUT $PROTO_FILE_PATH