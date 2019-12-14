package main

import "fmt"

type IPAddr [4]byte

// Sprintf, format belirticisine göre formatlar ve elde edilen dizgiyi döndürür.
// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var s string
	for i := range ip {
		if i == 0 {
			s += fmt.Sprint(int(ip[i]))
		} else {
			s += "." + fmt.Sprint(int(ip[i]))
		}
	}
	return s
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
