#!/usr/bin/env bash
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
## env GOOS=linux GOARCH=amd64 go build