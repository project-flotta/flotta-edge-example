package main

import (
	networkPkg "github.com/ahmadateya/flotta-edge-example/cmd/network"
	sensorsPkg "github.com/ahmadateya/flotta-edge-example/cmd/sensors"
	"os"
	"sync"
	"time"
)

var Destination = os.Getenv("CLUSTER_ADDRESS")
var LogsDir = os.Getenv("DEVICE_NAME")

func main() {
	// TODO make sure that /export/deviceName is created with the proper permissions
	if LogsDir == "" {
		LogsDir = "./export/device"
	}
	if Destination == "" {
		Destination = "google.com"
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		networkPkg.Start(Destination, time.Minute, LogsDir)
	}()

	go func() {
		defer wg.Done()
		sensorsPkg.Start(time.Minute, LogsDir)
	}()
	wg.Wait()
}
