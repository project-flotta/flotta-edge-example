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
	cpuTemp := getCPUTemp()
	saveCPUTemp(pathToSave, cpuTemp)
}

func getCPUTemp() CPUTemp {
	raw, err := os.ReadFile(ThermalZoneFile)
	if err != nil {
		fmt.Printf("Failed to read temperature from %q: %v", ThermalZoneFile, err)
	}
	cpuTempInt, err := strconv.Atoi(strings.TrimSpace(string(raw))) // e.g. 55306
	if err != nil {
		fmt.Printf("%q does not contain an integer: %v", ThermalZoneFile, err)
	}

	return CPUTemp{
		Temp: float64(cpuTempInt) / 1000.0,
	}
}

func saveCPUTemp(pathToSave string, cpuTemp CPUTemp) {
	f, err := os.OpenFile(fmt.Sprintf("%s/cputemp.log", pathToSave), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Could not open file %q: %v", fmt.Sprintf("%s/cputemp.log", pathToSave), err)
	}
	defer f.Close()

	log.SetOutput(f)
	b, err := json.Marshal(cpuTemp)
	if err != nil {
		fmt.Printf("Error marshaling cpuTemp: %v", err)
	}
	log.Println(string(b))
}
