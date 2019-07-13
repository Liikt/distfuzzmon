package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"distfuzzmon/client/client"
	"distfuzzmon/client/globals"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	globals.SetupGlobals("./dfm.conf")

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("[!] There will be a website here soon."))
	})
	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("{\"msg\": \"Nothing to see here. Please move on\"}"))
		})
	})

	client.BasicAPIRequest("registerclient", true)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n[~] Caught Ctrl+C. Cleaning up the client.")
		client.BasicAPIRequest("deregisterclient", true)
		os.Exit(0)
	}()
	fmt.Println("[+] Starting server on 31338")
	http.ListenAndServe(":31338", r)
}
