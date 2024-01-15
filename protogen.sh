#!/bin/bash

protoc  -I=. --java_out=./src/main/java pb/**/**/*.proto

