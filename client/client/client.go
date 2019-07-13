package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"distfuzzmon/client/globals"
	"distfuzzmon/client/types"
)

// BasicAPIRequest will make a basic API request to the server
func BasicAPIRequest(path string, exit bool) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:31337/api/%s/", globals.Conf.ServerIP, path), nil)
	if err != nil {
		fmt.Println("[-] Couldn't create a new request to the server.", err)
		if exit {
			os.Exit(1)
		} else {
			return
		}
	}

	req.Header.Set("X-Real-IP", globals.IP)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("[-] Couldn't send the request to the server.", err)
		if exit {
			os.Exit(1)
		} else {
			return
		}
	}

	var serverResp types.ServerRespose
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&serverResp)
	if err != nil {
		fmt.Println("[!] Couldn't decode the response from the server.", err)
		if exit {
			os.Exit(1)
		}
	}

	if serverResp.Message != "Ok!" {
		fmt.Println("[!] The server had some error:", serverResp.Message)
		if exit {
			os.Exit(1)
		}
	}
}

// StartFuzzjob is the endpoint to start a new fuzzing job
func StartFuzzjob(w http.ResponseWriter, r *http.Request) {

}

// StartClient will start the client activities
func StartClient() {
	ticker := time.NewTicker(30 * time.Minute)
	for {
		select {
		case <-ticker.C:
			// TODO: Stop the fuzzer
			sendFilesToServer()
		}
	}
}

func sendFilesToServer() {

}
