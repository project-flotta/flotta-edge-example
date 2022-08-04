package main

import (
	networkPkg "github.com/ahmadateya/flotta-edge-example/cmd/network"
	sensorsPkg "github.com/ahmadateya/flotta-edge-example/cmd/sensors"
	"sync"
)

const Destination = "google.com"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		networkPkg.Start(Destination, 1)
	}()

	go func() {
		defer wg.Done()
		sensorsPkg.Start(1)
	}()
	wg.Wait()
}
