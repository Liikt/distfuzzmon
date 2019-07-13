package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"distfuzzmon/server/types"

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

func enableClient(w http.ResponseWriter, r *http.Request) {
	if data, ok := types.RegisteredClients[r.RemoteAddr]; !ok {
		fmt.Printf("[~] Client from %s was never registered\n", r.RemoteAddr)
	} else if data {
		types.RegisteredClients[r.RemoteAddr] = true
	}
	fmt.Println("[+] Successfully enabled", r.RemoteAddr)
}

func disableClient(w http.ResponseWriter, r *http.Request) {
	if data, ok := types.RegisteredClients[r.RemoteAddr]; !ok {
		fmt.Printf("[~] Client from %s was never registered\n", r.RemoteAddr)
	} else if data {
		types.RegisteredClients[r.RemoteAddr] = false
	}
	fmt.Println("[+] Successfully disabled", r.RemoteAddr)
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
		r.Get("/enable_client/", enableClient)
		r.Get("/disable_client/", disableClient)
	})

	fmt.Println("[+] Starting server on 31337")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n[~] Caught Ctrl+C. Exiting the server now.")
		os.Exit(0)
	}()
	http.ListenAndServe(":31337", r)
}
