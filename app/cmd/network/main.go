package network

import (
	"github.com/ahmadateya/flotta-edge-example/pkg/traceroute"
	"time"
)

func Start(server string, intervalInMin int) {
	options := traceroute.Options{}
	for {
		traceroute.GoTraceroute(server, options, "./tmp")
		time.Sleep(time.Duration(intervalInMin) * time.Minute)
	}
}
