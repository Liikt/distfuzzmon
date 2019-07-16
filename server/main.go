package main

import (
	"bytes"
	"encoding/json"
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
	r.Use(middleware.Logger)

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
		r.Get("/sync/{target}/", clientsync.Sync)
		r.Post("/dropfile/{target}/{fuzzer}/", clientsync.DropFile)
	})

	go test()

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

func test() {
	ips := []string{"192.168.0.6", "192.168.0.6"}
	for _, ip := range ips {
		fj := types.Fuzzjob{
			FullCommand: "test",
			Target:      "testfuzz",
			Fuzzer:      "afl",
			FuzzerCount: 3,
			Seeds: []string{
				"Cg==",
				"aW5kZXgK",
				"aW1hZ2VzCg==",
				"ZG93bmxvYWQK",
				"MjAwNgo=",
				"bmV3cwo=",
				"Y3JhY2sK",
				"c2VyaWFsCg==",
				"d2FyZXoK",
				"ZnVsbAo=",
				"MTIK",
				"Y29udGFjdAo=",
				"YWJvdXQK",
				"c2VhcmNoCg==",
				"c3BhY2VyCg==",
				"cHJpdmFjeQo=",
				"MTEK",
				"bG9nbwo=",
				"YmxvZwo=",
				"bmV3Cg==",
				"MTAK",
				"Y2dpLWJpbgo=",
				"ZmFxCg==",
				"cnNzCg==",
				"aG9tZQo=",
				"aW1nCg==",
				"ZGVmYXVsdAo=",
				"MjAwNQo=",
				"cHJvZHVjdHMK",
				"c2l0ZW1hcAo=",
			},
		}

		var clientResp types.ClientResponse
		jsonValue, _ := json.Marshal(fj)
		r, _ := http.Post(fmt.Sprintf("http://%s:31337/api/dropfile/%s/", ip, fj.Target), "application/json", bytes.NewBuffer(jsonValue))
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&clientResp)
		if err != nil {
			fmt.Println("[!] Error parsing body from response")
			continue
		}
		if clientResp.Message != "Ok!" {
			fmt.Println("[!] Client returned:", clientResp.Message)
		}
	}
}
