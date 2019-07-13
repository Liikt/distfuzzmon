package clienthandeling

import (
	"fmt"
	"net/http"

	"distfuzzmon/server/types"
)

// RegisterClient will register a client matching the X-Real-IP header
func RegisterClient(w http.ResponseWriter, r *http.Request) {
	if _, ok := types.RegisteredClients[r.RemoteAddr]; ok {
		fmt.Printf("[~] Client from %s is already registered\n", r.RemoteAddr)
		w.Write([]byte("{\"msg\": \"Client is already registered\"}"))
	} else {
		types.RegisteredClients[r.RemoteAddr] = true
		fmt.Println("[+] Registered new clients on", r.RemoteAddr)
		w.Write([]byte("{\"msg\": \"Ok!\"}"))
	}
}

// DeregisterClient will remove the client matching the X-Real-IP header
func DeregisterClient(w http.ResponseWriter, r *http.Request) {
	if _, ok := types.RegisteredClients[r.RemoteAddr]; ok {
		delete(types.RegisteredClients, r.RemoteAddr)
		fmt.Println("[+] Successfully deregistered", r.RemoteAddr)
		w.Write([]byte("{\"msg\": \"Ok!\"}"))
	} else {
		fmt.Printf("[~] Client from %s was never registered\n", r.RemoteAddr)
		w.Write([]byte("{\"msg\": \"Client never registered\"}"))
	}
}

//EnableClient will enable the client matching the X-Real-IP header
func EnableClient(w http.ResponseWriter, r *http.Request) {
	if data, ok := types.RegisteredClients[r.RemoteAddr]; !ok {
		fmt.Printf("[~] Client from %s was never registered\n", r.RemoteAddr)
	} else if data {
		types.RegisteredClients[r.RemoteAddr] = true
	}
	fmt.Println("[+] Successfully enabled", r.RemoteAddr)
	w.Write([]byte("{\"msg\": \"Ok!\"}"))
}

//DisableClient will disable the client matching the X-Real-IP header
func DisableClient(w http.ResponseWriter, r *http.Request) {
	if data, ok := types.RegisteredClients[r.RemoteAddr]; !ok {
		fmt.Printf("[~] Client from %s was never registered\n", r.RemoteAddr)
	} else if data {
		types.RegisteredClients[r.RemoteAddr] = false
	}
	fmt.Println("[+] Successfully disabled", r.RemoteAddr)
	w.Write([]byte("{\"msg\": \"Ok!\"}"))
}
