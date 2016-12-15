package main

import (
	"flag"

	"github.com/gvsro/plaincast/server"
)

func main() {
	flag.Parse()

	server.Serve()
}
