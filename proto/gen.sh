#!/usr/bin/env bash
protoc -I./ rpc.proto --go_out=plugins=grpc:../rpc
