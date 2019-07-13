package client

import (
	"distfuzzmon/client/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"distfuzzmon/client/globals"
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
