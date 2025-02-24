/*****************************
Author: Ashwin Nair
Date: 2025-02-24
Project name: main4.go
Package: main4
Summary: Tests localhost IP pinging dashboard (very rudimentary)
*****************************/

package main4

import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

// PingIP executes the ping command and extracts latency
func PingIP(ip string) (string, error) {
	cmd := exec.Command("ping", "-c", "1", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	outStr := string(output)
	re := regexp.MustCompile(`time=([\d.]+)\s*ms`)
	matches := re.FindStringSubmatch(outStr)

	if len(matches) < 2 {
		return "Latency not found", nil
	}
	return matches[1] + " ms", nil
}

// HTTP handler to get latency
func latencyHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "Missing 'ip' query parameter", http.StatusBadRequest)
		return
	}

	latency, err := PingIP(ip)
	if err != nil {
		http.Error(w, "Error pinging IP", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Latency to %s: %s\n", ip, latency)
}

func main() {
	http.HandleFunc("/latency", latencyHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

