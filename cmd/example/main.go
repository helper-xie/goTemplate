package main

import (
	"flag"
)

var configFile = flag.String("c", "/etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	ser := initServ(*configFile)
	_ = ser.Start()
}
