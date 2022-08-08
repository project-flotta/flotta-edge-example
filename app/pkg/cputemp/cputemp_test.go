package cputemp

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSaveCPUTempFunc(t *testing.T) {
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

func TestGetCPUTemp(t *testing.T) {
	raw, _ := os.ReadFile(ThermalZoneFile)
	temp := getCPUTemp()
	cpuTempInt, _ := strconv.Atoi(strings.TrimSpace(string(raw))) // e.g. 55306

	if temp.Temp != float64(cpuTempInt)/1000.0 {
		t.Errorf("%v does not equal to %v testing data", temp.Temp, cpuTempInt)
	}
}
