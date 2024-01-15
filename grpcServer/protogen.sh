#!/bin/bash


protoc -I=. --go-grpc_out=../../.. --go_out=../../.. pb/**/**/*.proto

