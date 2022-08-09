package main

import (
	networkPkg "github.com/ahmadateya/flotta-edge-example/cmd/network"
	sensorsPkg "github.com/ahmadateya/flotta-edge-example/cmd/sensors"
	"os"
	"sync"
	"time"
)

var Destination = os.Getenv("CLUSTER_ADDRESS")
var LogsDir = os.Getenv("LOGS_DIR")

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		networkPkg.Start(Destination, time.Minute, LogsDir)
	}()

	go func() {
		defer wg.Done()
		sensorsPkg.Start(1)
	}()
	wg.Wait()
}
