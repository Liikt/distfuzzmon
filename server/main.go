package main

import (
	"distfuzzmon/server/types"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func registerClient(w http.ResponseWriter, r *http.Request) {
	if _, ok := types.RegisteredClients[r.RemoteAddr]; ok {
		fmt.Printf("[~] Client from %s is already registered\n", r.RemoteAddr)
	} else {
		types.RegisteredClients[r.RemoteAddr] = true
		fmt.Println("[+] Registered new clients on", r.RemoteAddr)
	}
}

func deregisterClient(w http.ResponseWriter, r *http.Request) {
	if _, ok := types.RegisteredClients[r.RemoteAddr]; ok {
		delete(types.RegisteredClients, r.RemoteAddr)
		fmt.Println("[+] Successfully deregistered", r.RemoteAddr)
	} else {
		fmt.Printf("[~] Client from %s was never registered\n", r.RemoteAddr)
	}
}

func main() {
	r := chi.NewRouter()
	types.SetupGlobals()

	r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("[!] There will be a webserver"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("[-] Nothing to see here. Please move on."))
		})
		r.Get("/register_client/", registerClient)
		r.Get("/deregister_client/", deregisterClient)
	})

	fmt.Println("[+] Starting server on 31337")
	http.ListenAndServe(":31337", r)
}
