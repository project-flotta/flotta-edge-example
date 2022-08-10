package sensors

import (
	"github.com/ahmadateya/flotta-edge-example/pkg/cputemp"
	"time"
)

func Start(sleep time.Duration, pathToSave string) {
	for {
		cputemp.ReadCPUTemp(pathToSave)
		time.Sleep(sleep)
	}
}
