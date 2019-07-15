package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"distfuzzmon/client/globals"
	"distfuzzmon/client/types"
)

func startFuzzjob(fj types.Fuzzjob) {
	if _, ok := globals.ActiveFuzzers[fj.Target]; ok {
		fmt.Printf("[!] Target %s is already being fuzzed.\n", fj.Target)
		return
	}

	createTargetFolder(fj)

	fj.Disable = make(chan struct{})
	fj.Stop = make(chan struct{})

	globals.ActiveFuzzers[fj.Target] = &fj

	ticker := time.NewTicker(30 * time.Minute)
	for {
		select {
		case <-ticker.C:
			sendFilesToServer(fj.Target)
		}
	}
}

func createTargetFolder(fj types.Fuzzjob) {
	os.MkdirAll(filepath.Join(globals.Conf.BaseFolder, fj.Target, "in"), 0755)
	os.MkdirAll(filepath.Join(globals.Conf.BaseFolder, fj.Target, "out"), 0755)

	for i, content := range fj.Seeds {
		content, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			fmt.Printf("[!] Error decoding %s\n", content)
			continue
		}
		seedName := strconv.Itoa(i)
		ioutil.WriteFile(filepath.Join(globals.Conf.BaseFolder, fj.Target, "in", seedName), content, 0644)
	}
}

func sendFilesToServer(target string) {
	path := filepath.Join(globals.Conf.BaseFolder, target)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("[!] Something went wrong while listing the fuzzers directory. Error:", err)
		return
	}

	for _, file := range files {
		fileContent, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			fmt.Println("[!] Something went wrong while reading a file. Error:", err)
			continue
		}

		dropReq := types.DropFileRequest{Filename: file.Name(), Content: base64.StdEncoding.EncodeToString(fileContent)}
		jsonValue, _ := json.Marshal(dropReq)

		var clientResp types.ServerRespose
		r, _ := http.Post(fmt.Sprintf("http://%s:31337/api/dropfile/%s/", globals.Conf.ServerIP, target), "application/json", bytes.NewBuffer(jsonValue))
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&clientResp)
		if err != nil {
			fmt.Println("[!] Error parsing body from response")
			continue
		}
		if clientResp.Message != "Ok!" {
			fmt.Println("[!] Client returned:", clientResp.Message)
		}
	}
}
