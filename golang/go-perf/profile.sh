#!/bin/sh
go build -o main main.go
rm mem.prof
./main
go tool pprof -http=":8082" ./main ./mem.prof
