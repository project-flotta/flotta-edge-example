package traceroute

import (
	"encoding/json"
	"fmt"
	"net"
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
	ipAddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return
	}

	fmt.Printf("traceroute to %v (%v), %v hops max, %v byte packets\n", host, ipAddr, options.MaxHops(), options.PacketSize())

	c := make(chan Hop, 0)
	go func() {
		for {
			hop, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
			printHop(hop)
		}
	}()

	result, err := traceroute(host, &options, c)
	if err != nil {
		fmt.Printf("Error: ", err)
	}

	if pathToSave != "" {
		f, _ := os.OpenFile(fmt.Sprintf("%s/%s.json", pathToSave, host), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()

		b, _ := json.Marshal(result)

		if _, err = f.WriteString("\n" + string(b)); err != nil {
			panic(fmt.Sprintf("cant write to log file %v ", err))
		}
	}
}