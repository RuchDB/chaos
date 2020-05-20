package main

import (
	"github.com/RuchDB/chaos/cache/log"

	"github.com/RuchDB/chaos/lib/g/net"
)

func main() {
	log.Info("Welcome to Chaos World!")

	net.NewTcp4Server(net.IPV4_ANY, 2020).Start()
}
