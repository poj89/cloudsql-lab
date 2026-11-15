package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Raspberry Pi Health Monitor ===")

	printHostname()
	printUptime()
	printLoadAverage()
	printTemperature()
	printMemory()
	printDisk()
}

func printHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Hostname: unavailable")
		return
	}

	fmt.Printf("Hostname: %s\n", hostname)
}

func printUptime() {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		fmt.Println("Uptime: unavailable")
		return
	}

	fields := strings.Fields(string(data))
	if len(fields) == 0 {
		return
	}

	seconds, _ := strconv.ParseFloat(fields[0], 64)

	hours := int(seconds) / 3600
	minutes := (int(seconds) % 3600) / 60

	fmt.Printf("Uptime: %dh %dm\n", hours, minutes)
}

func printLoadAverage() {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		fmt.Println("Load Average: unavailable")
		return
	}

	fields := strings.Fields(string(data))
	if len(fields) < 3 {
		return
	}

	fmt.Printf(
		"Load Average: %s %s %s\n",
		fields[0],
		fields[1],
		fields[2],
	)
}

func printTemperature() {
	data, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		fmt.Println("CPU Temp: unavailable")
		return
	}

	tempMilliC, _ := strconv.Atoi(strings.TrimSpace(string(data)))
	tempC := float64(tempMilliC) / 1000.0

	fmt.Printf("CPU Temp: %.1f°C\n", tempC)
}

func printMemory() {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println("Memory: unavailable")
		return
	}

	var totalKB int
	var availableKB int

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "MemTotal:") {
			fmt.Sscanf(line, "MemTotal: %d kB", &totalKB)
		}

		if strings.HasPrefix(line, "MemAvailable:") {
			fmt.Sscanf(line, "MemAvailable: %d kB", &availableKB)
		}
	}

	usedKB := totalKB - availableKB
	usedPercent := float64(usedKB) / float64(totalKB) * 100

	fmt.Printf(
		"Memory: %.1f%% (%d MB / %d MB)\n",
		usedPercent,
		usedKB/1024,
		totalKB/1024,
	)
}

func printDisk() {
	fmt.Println("Disk: use 'df -h' (coming in next version)")
}
