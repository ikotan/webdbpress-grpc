#!/usr/bin/env bash

PROTO_DIR="./uploader/proto/"
PLUGINS="grpc"
DEST_PATH="./uploader/proto/"
GO_OUT="plugins=$PLUGINS:$DEST_PATH"
PROTO_FILE_PATH="./uploader/proto/file.proto"

protoc --proto_path $PROTO_DIR --go_out=$GO_OUT $PROTO_FILE_PATH