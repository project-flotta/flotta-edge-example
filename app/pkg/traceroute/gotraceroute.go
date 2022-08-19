package traceroute

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func printHop(hop Hop) {
	addr := fmt.Sprintf("%v.%v.%v.%v", hop.Address[0], hop.Address[1], hop.Address[2], hop.Address[3])
	hostOrAddr := addr
	if hop.Host != "" {
		hostOrAddr = hop.Host
	}
	if hop.Success {
		fmt.Printf("%-3d %v (%v)  %v\n", hop.TTL, hostOrAddr, addr, hop.ElapsedTime)
	} else {
		fmt.Printf("%-3d *\n", hop.TTL)
	}
}

func GoTraceroute(host string, options Options, pathToSave string) {
	c := make(chan Hop, 0)
	go func() {
		for {
			_, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
			//printHop(hop)
		}
	}()

	result, err := traceroute(host, &options, c)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	if pathToSave != "" {
		f, err := os.OpenFile(fmt.Sprintf("%s/traceroute-to-%s.log", pathToSave, host), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Printf("could not open file %q: %v\n", fmt.Sprintf("%s/traceroute-to-%s.log", pathToSave, host), err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Printf("could not close file %q: %v\n", f.Name(), err)
			}
		}(f)

		b, err := json.Marshal(result)
		if err != nil {
			log.Printf("Error marshaling result: %v", err)
		}

		log.SetOutput(f)
		log.Println(string(b))
	}
}
