package utils

import (
	"fmt"
	"net"
	"os"
)

// Instr returns true if the character char is in the string str. Otherwise false
func Instr(str, char string) bool {
	for _, c := range str {
		if char == string(c) {
			return true
		}
	}
	return false
}

// FindIP finds the first ipv4 address of an interface that is not localhost
func FindIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("[-] Coulnd't find interfaces. Reason:", err)
		os.Exit(1)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println("[-] Coulnd't get address of iface. Reason:", err)
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ipstr := fmt.Sprintf("%v", ip)
			if ipstr != "127.0.0.1" && Instr(ipstr, ".") {
				return ipstr
			}
		}
	}
	return ""
}
