package network

import (
	"github.com/ahmadateya/flotta-edge-example/pkg/traceroute"
	"time"
)

func Start(server string, sleep time.Duration, pathToSave string) {
	options := traceroute.Options{}
	for {
		traceroute.GoTraceroute(server, options, pathToSave)
		time.Sleep(sleep)
	}
}
