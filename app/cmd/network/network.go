package network

import (
	"github.com/ahmadateya/flotta-edge-example/pkg/traceroute"
	"time"
)

func Start(server string, sleep time.Duration, pathToSave string) {
	if pathToSave == "" {
		pathToSave = "./tmp"
	}
	if server == "" {
		server = "google.com"
	}
	options := traceroute.Options{}
	for {
		traceroute.GoTraceroute(server, options, pathToSave)
		time.Sleep(sleep)
	}
}
