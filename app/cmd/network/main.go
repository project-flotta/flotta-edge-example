package network

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//generateDummyLogs()
}

func TracerouteCmd(server string) {

	cmd := exec.Command("traceroute", server)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	//fmt.Println(string(stdout))
	f, err := os.Create(fmt.Sprintf("./%s.txt", server))
	fmt.Fprintf(f, "======= server: %s \n %v\n", server, string(stdout))
}
