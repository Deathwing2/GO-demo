package main

import (
	"fmt"
	"net"
	"time"
)

func Ping(host string) (time.Duration, error) {
	start := time.Now()
	_, err := net.LookupHost(host)
	if err != nil {
		return 0, err
	}
	return time.Since(start), nil
}

func main() {
	time, err := Ping("google.com")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Ping time: %v\n", time)
}
