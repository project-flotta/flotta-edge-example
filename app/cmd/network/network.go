package network

import (
	"github.com/ahmadateya/flotta-edge-example/pkg/traceroute"
	"time"
)

func Start(server string, d time.Duration) {
	options := traceroute.Options{}
	for {
		traceroute.GoTraceroute(server, options, "./tmp")
		time.Sleep(d)
	}
}
