package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"distfuzzmon/server/clientHandeling"
	"distfuzzmon/server/types"

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

	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("[-] Nothing to see here. Please move on."))
		})
		r.Get("/register_client/", clientHandeling.RegisterClient)
		r.Get("/deregister_client/", clientHandeling.DeregisterClient)
		r.Get("/enable_client/", clientHandeling.EnableClient)
		r.Get("/disable_client/", clientHandeling.DisableClient)
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
