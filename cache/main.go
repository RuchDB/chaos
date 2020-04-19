package main

import (
	"github.com/RuchDB/chaos/lib/log"
	"github.com/RuchDB/chaos/lib/net"
)

func main() {
	log.Info("Welcome to Chaos World!")

	net.NewTcp4Server(net.IPV4_ANY, 2020).Start()
}
