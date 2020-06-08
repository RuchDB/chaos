package main

import (
	"github.com/RuchDB/chaos/cache/log"

	"github.com/RuchDB/chaos/lib/g/net"
	"github.com/RuchDB/chaos/lib/g/net/codec"
	"github.com/RuchDB/chaos/lib/g/util"
)

func main() {
	log.Logger.Info("Welcome to Chaos World!")

	util.SetOpenFileLimit(2000)

	creator := net.NewServerCreator()
	creator.ConfigListenAddrByIpPort(net.IPV4_ANY, 2020)
	creator.ConfigNetCodecById(codec.CODEC_ID_REDIS)
	creator.ConfigMaxConns(1000)

	server, err := creator.CreateServer()
	if err != nil {
		log.Logger.Errorf("Create TCP server error: [%v]", err)
		return
	}

	server.Run()
}
