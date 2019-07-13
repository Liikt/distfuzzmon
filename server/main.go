package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"distfuzzmon/server/clienthandeling"
	"distfuzzmon/server/clientsync"
	"distfuzzmon/server/types"
	"distfuzzmon/server/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	types.SetupGlobals()
	if exists, err := utils.PathExists("sync_dir"); err != nil {
		fmt.Println("[-] Something went wrong while stating the sync dir")
		os.Exit(1)
	} else if !exists {
		os.Mkdir("sync_dir", 0755)
	}

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("[!] There will be a webserver"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("[-] Nothing to see here. Please move on."))
		})
		r.Get("/registerclient/", clienthandeling.RegisterClient)
		r.Get("/deregisterclient/", clienthandeling.DeregisterClient)
		r.Get("/enableclient/", clienthandeling.EnableClient)
		r.Get("/disableclient/", clienthandeling.DisableClient)
		r.Post("/dropfile/{target}/{fuzzer}/", clientsync.DropFile)
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
