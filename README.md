# go-gin-test

USAGE:
--------------
main [global options] command [command options] [arguments...]


Quick Start:
--------------
 1. git clone
 2. go mod tidy
 3. go run cmd/main.go start
--------------

How to watch go file changes?

1.https://gitee.com/liudng/dogo
2.mkdir bin
3.add dogo.json file
dogo.json example:
{
  "WorkingDir": "./",
  "SourceDir": [ "./" ],
  "SourceExt": ".go|.c|.cpp|.h",
   "BuildCmd": "go build -o ./bin/go-gin-test",
   "RunCmd": "./bin/go-gin-test --debug start"
}