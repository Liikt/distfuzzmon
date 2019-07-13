package main

import (
	"distfuzzmon/server/types"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	types.SetupGlobals()

	r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("[!] There will be a webserver"))
	})
	r.Get("/register_client/", func(w http.ResponseWriter, r *http.Request) {
		if _, ok := types.RegisteredClients[r.RemoteAddr]; ok {
			fmt.Printf("[~] Client from %s is already registered\n", r.RemoteAddr)
		} else {
			types.RegisteredClients[r.RemoteAddr] = true
			fmt.Println("[+] Registered new clients on", r.RemoteAddr)
		}
	})

	fmt.Println("[+] Starting server on 31337")
	http.ListenAndServe(":31337", r)
}
