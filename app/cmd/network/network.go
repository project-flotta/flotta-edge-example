package network

import (
	"github.com/ahmadateya/flotta-edge-example/pkg/traceroute"
	"os"
	"time"
)

var LogsDir = os.Getenv("LOGS_DIR")

func Start(server string, d time.Duration) {
	options := traceroute.Options{}
	for {
		traceroute.GoTraceroute(server, options, LogsDir)
		time.Sleep(d)
	}
}
