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
	//ipAddr, err := net.ResolveIPAddr("ip", host)
	//if err != nil {
	//	return
	//}

	//fmt.Printf("traceroute to %v (%v), %v hops max, %v byte packets\n", host, ipAddr, options.MaxHops(), options.PacketSize())

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
		fmt.Printf("Error: %v", err)
	}

	if pathToSave != "" {
		f, _ := os.OpenFile(fmt.Sprintf("%s/traceroute-to-%s.log", pathToSave, host), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()

		b, _ := json.Marshal(result)

		log.SetOutput(f)
		log.Println(string(b))
	}
}
