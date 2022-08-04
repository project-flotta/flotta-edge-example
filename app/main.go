package main

import (
	networkPkg "github.com/ahmadateya/flotta-edge-example/cmd/network"
	"sync"
)

const Destination = "google.com"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		networkPkg.Start(Destination, 1)
	}()

	wg.Wait()
}
