#!/usr/bin/env bash
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
##env env GOOS=linux GOARCH=amd64 go build