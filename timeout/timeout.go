package main

import (
	"fmt"
	"os/exec"
	"time"
)

func ping(host string, timeout time.Duration) error {
	cmd := exec.Command("ping", "-c 1", "-W", fmt.Sprintf("%d", int(timeout.Seconds())), host)
	err := cmd.Run()
	return err
}

func main() {
	host := "google.com"
	timeout := 2 * time.Second // Таймаут в 2 секунды

	err := ping(host, timeout)
	if err != nil {
		fmt.Printf("Ping %s failed: %v\n", host, err)
	} else {
		fmt.Printf("Ping to %s succeeded!\n", host)
	}
}
