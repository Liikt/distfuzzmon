package globals

import (
	"fmt"
	"net"
	"os"

	"distfuzzmon/client/config"
	"distfuzzmon/client/utils"
)

var (
	Conf config.Config
	IP   string
)

func SetupGlobals(configPath string) {
	IP = findIP()
	Conf = config.LoadConfig(configPath)
}

func findIP() string {
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
			if ipstr != "127.0.0.1" && utils.Instr(ipstr, ".") {
				return ipstr
			}
		}
	}
	return ""
}
