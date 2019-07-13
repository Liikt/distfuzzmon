package clientHandeling

import (
	"fmt"
	"net/http"

	"distfuzzmon/server/types"
)

func RegisterClient(w http.ResponseWriter, r *http.Request) {
	if _, ok := types.RegisteredClients[r.RemoteAddr]; ok {
		fmt.Printf("[-] Client from %s is already registered\n", r.RemoteAddr)
	} else {
		types.RegisteredClients[r.RemoteAddr] = true
		fmt.Println("[+] Registered new clients on", r.RemoteAddr)
	}
}

func DeregisterClient(w http.ResponseWriter, r *http.Request) {
	if _, ok := types.RegisteredClients[r.RemoteAddr]; ok {
		delete(types.RegisteredClients, r.RemoteAddr)
		fmt.Println("[+] Successfully deregistered", r.RemoteAddr)
	} else {
		fmt.Printf("[-] Client from %s was never registered\n", r.RemoteAddr)
	}
}

func EnableClient(w http.ResponseWriter, r *http.Request) {
	if data, ok := types.RegisteredClients[r.RemoteAddr]; !ok {
		fmt.Printf("[-] Client from %s was never registered\n", r.RemoteAddr)
	} else if data {
		types.RegisteredClients[r.RemoteAddr] = true
	}
	fmt.Println("[+] Successfully enabled", r.RemoteAddr)
}

func DisableClient(w http.ResponseWriter, r *http.Request) {
	if data, ok := types.RegisteredClients[r.RemoteAddr]; !ok {
		fmt.Printf("[-] Client from %s was never registered\n", r.RemoteAddr)
	} else if data {
		types.RegisteredClients[r.RemoteAddr] = false
	}
	fmt.Println("[+] Successfully disabled", r.RemoteAddr)
}
