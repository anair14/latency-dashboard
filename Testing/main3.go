/*
Summary: Pings google as a test and returns latency.

*/

package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

// PingIP pings an IP address and extracts latency
func PingIP(ip string) (string, error) {
	// Execute the ping command (works on Linux/macOS)
	cmd := exec.Command("ping", "-c", "1", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	// Convert output to string
	outStr := string(output)

	// Regex to extract time=x.x ms
	re := regexp.MustCompile(`time=([\d.]+)\s*ms`)
	matches := re.FindStringSubmatch(outStr)

	if len(matches) < 2 {
		return "Latency not found", nil
	}
	
	return matches[1] + " ms", nil
}

func main() {
	ip := "8.8.8.8" // Example: Google's DNS
	latency, err := PingIP(ip)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Latency to %s: %s\n", ip, latency)
}

