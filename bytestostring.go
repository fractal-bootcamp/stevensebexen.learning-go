package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	s := []string{}
	for i := 0; i < 4; i++ {
		s = append(s, fmt.Sprintf("%d", ip[i]))
	}
	return strings.Join(s, ".")
}

func BytesToString() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
