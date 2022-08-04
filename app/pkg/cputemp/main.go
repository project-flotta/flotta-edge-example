package cputemp

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const ThermalZoneFile = "/sys/class/thermal/thermal_zone0/temp"

type CPUTemp struct {
	Temp float64
}

func ReadCPUTemp(pathToSave string) {
	raw, err := os.ReadFile(ThermalZoneFile)
	if err != nil {
		fmt.Printf("Failed to read temperature from %q: %v", ThermalZoneFile, err)
	}

	cpuTempStr := strings.TrimSpace(string(raw))
	cpuTempInt, err := strconv.Atoi(cpuTempStr) // e.g. 55306
	if err != nil {
		fmt.Printf("%q does not contain an integer: %v", ThermalZoneFile, err)
	}

	cpuTemp := CPUTemp{
		Temp: float64(cpuTempInt) / 1000.0,
	}

	f, _ := os.OpenFile(fmt.Sprintf("%s/cputemp.log", pathToSave), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	log.SetOutput(f)
	b, _ := json.Marshal(cpuTemp)
	log.Println(string(b))
}
