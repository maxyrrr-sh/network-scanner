package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const (
	timeout  = 2 * time.Second
	maxHosts = 254
)

func pingHost(host string) bool {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", host)
	output, err := cmd.CombinedOutput()
	if err != nil || !strings.Contains(string(output), "1 received") {
		return false
	}
	return true
}

func scanNetwork(baseIP string) []string {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	availableHosts := []string{}

	for i := 1; i <= maxHosts; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			ip := fmt.Sprintf("%s.%d", baseIP, i)
			if pingHost(ip) {
				mutex.Lock()
				availableHosts = append(availableHosts, ip)
				mutex.Unlock()
			}
		}(i)
	}

	wg.Wait()
	return availableHosts
}

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String(), nil
		}
	}

	return "", fmt.Errorf("не вдалося отримати локальну IP-адресу")
}