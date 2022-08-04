package sensors

import (
	"github.com/ahmadateya/flotta-edge-example/pkg/cputemp"
	"time"
)

func Start(intervalInMin int) {
	for {
		cputemp.ReadCPUTemp("./tmp")
		time.Sleep(time.Duration(intervalInMin) * time.Minute)
	}
}
