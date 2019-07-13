package client

import (
	"distfuzzmon/client/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"distfuzzmon/client/globals"
)

// RegisterClient will start the client activities
func RegisterClient() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:31337/api/registerclient/", globals.Conf.ServerIP), nil)
	if err != nil {
		fmt.Println("[-] Couldn't create a new request to the server.", err)
		os.Exit(1)
	}

	req.Header.Set("X-Real-IP", globals.IP)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("[-] Couldn't send the request to the server.", err)
		os.Exit(1)
	}

	var serverResp types.ServerRespose
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&serverResp)
	if err != nil {
		fmt.Println("[!] Couldn't decode the response from the server.", err)
	}

	if serverResp.Message != "Ok!" {
		fmt.Println("[!] The server had some error:", serverResp.Message)
	}
}

// DeregisterClient will deregister the client when Ctrl+C was hit
func DeregisterClient() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:31337/api/deregisterclient/", globals.Conf.ServerIP), nil)
	if err != nil {
		fmt.Println("[-] Couldn't create a new request to the server.", err)
		os.Exit(1)
	}

	req.Header.Set("X-Real-IP", globals.IP)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("[-] Couldn't send the request to the server.", err)
		os.Exit(1)
	}

	var serverResp types.ServerRespose
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&serverResp)
	if err != nil {
		fmt.Println("[-] Couldn't decode the response from the server.", err)
		os.Exit(1)
	}

	if serverResp.Message != "Ok!" {
		fmt.Println("[!] The server had some error:", serverResp.Message)
	}
}

// EnableClient will reenable the client if it was disabled
func EnableClient() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:31337/api/enableclient/", globals.Conf.ServerIP), nil)
	if err != nil {
		fmt.Println("[-] Couldn't create a new request to the server.", err)
		return
	}

	req.Header.Set("X-Real-IP", globals.IP)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("[-] Couldn't send the request to the server.", err)
		return
	}

	var serverResp types.ServerRespose
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&serverResp)
	if err != nil {
		fmt.Println("[-] Couldn't decode the response from the server.", err)
		return
	}

	if serverResp.Message != "Ok!" {
		fmt.Println("[!] The server had some error:", serverResp.Message)
	}
}

// DisableClient will reenable the client if it was disabled
func DisableClient() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:31337/api/disableclient/", globals.Conf.ServerIP), nil)
	if err != nil {
		fmt.Println("[-] Couldn't create a new request to the server.", err)
		return
	}

	req.Header.Set("X-Real-IP", globals.IP)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("[-] Couldn't send the request to the server.", err)
		return
	}

	var serverResp types.ServerRespose
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&serverResp)
	if err != nil {
		fmt.Println("[-] Couldn't decode the response from the server.", err)
		return
	}

	if serverResp.Message != "Ok!" {
		fmt.Println("[!] The server had some error:", serverResp.Message)
	}
}
