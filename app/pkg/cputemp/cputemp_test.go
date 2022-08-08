package cputemp

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSaveCPUTempFunc(t *testing.T) {
	fmt.Println("Testing saveCPUTemp Func...")
	tmp := CPUTemp{Temp: 55.5}
	saveCPUTemp("/tmp", tmp)
	//check if the file exists
	f, _ := os.OpenFile(fmt.Sprintf("%s/cputemp.log", "/tmp"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	//check if the file contains the correct data
	raw, err := os.ReadFile(fmt.Sprintf("%s/cputemp.log", "/tmp"))
	if err != nil {
		t.Errorf("Failed to read temperature from %q: %v", "/tmp", err)
	}

	existed := strings.Contains(string(raw), "55.5")
	if !existed {
		t.Errorf("%q does not contain the correct data", "/tmp")
	}
}

func TestReadCPUTemp(t *testing.T) {
	fmt.Println("Testing Read CPU Temp")
	//out, err := traceroute("google.com", new(Options))
	//if err == nil {
	//	if len(out.Hops) == 0 {
	//		t.Errorf("TestTraceroute failed. Expected at least one hop")
	//	}
	//} else {
	//	t.Errorf("TestTraceroute failed due to an error: %v", err)
	//}
	//
	//for _, hop := range out.Hops {
	//	printHop(hop)
	//}
	//fmt.Println()
}
