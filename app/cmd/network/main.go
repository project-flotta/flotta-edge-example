package network

import (
	tracerouteCmd "github.com/ahmadateya/traceroute/cmd"
	"github.com/ahmadateya/traceroute/traceroute"
)

func main() {
	//generateDummyLogs()
}

func TracerouteCmd(server string) {
	options := traceroute.TracerouteOptions{}
	tracerouteCmd.GoTraceroute(server, options, "./tmp")
}
