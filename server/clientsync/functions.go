package clientsync

import (
	"bytes"
	"distfuzzmon/server/types"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"distfuzzmon/server/utils"

	"github.com/go-chi/chi"
)

// DropFile will save a file that the master recieved from a client
func DropFile(w http.ResponseWriter, r *http.Request) {
	fuzzer := chi.URLParam(r, "fuzzer")
	target := chi.URLParam(r, "target")
	var reqContent types.DropFileRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqContent)
	if err != nil {
		fmt.Printf("[!] Error parsing body from fuzzer: %s target: %s\n", fuzzer, target)
		w.WriteHeader(500)
		w.Write([]byte("{\"msg\": \"Couldn't parse the body of a file.\"}"))
		return
	}
	content, err := base64.StdEncoding.DecodeString(reqContent.Content)
	if err != nil {
		fmt.Printf("[!] Error decoding %s\n", reqContent.Content)
		w.WriteHeader(500)
		w.Write([]byte("{\"msg\": \"Couldn't decode the body of a file.\"}"))
		return
	}

	fuzzerDir := filepath.Join("sync_dir", target, fuzzer)
	if exists, err := utils.PathExists(fuzzerDir); err != nil {
		fmt.Println("[-] Something went wrong while stating the sync dir")
		w.WriteHeader(500)
		w.Write([]byte("{\"msg\": \"Couldn't stat the sync dir.\"}"))
		return
	} else if !exists {
		os.MkdirAll(fuzzerDir, 0755)
	}

	ioutil.WriteFile(filepath.Join(fuzzerDir, reqContent.Filename), content, 0644)
	w.Write([]byte("{\"msg\": \"Ok!\"}"))
}

// Sync is an endpoint the client calls to get back a synced version of it's target binary
func Sync(w http.ResponseWriter, r *http.Request) {
	target := chi.URLParam(r, "target")
	ip := r.RemoteAddr
	targetPath := filepath.Join("sync_dir", target)

	fuzzers, err := ioutil.ReadDir(targetPath)
	if err != nil {
		fmt.Println("[-] Something went wrong while listing the directory. Error:", err)
		w.WriteHeader(500)
		w.Write([]byte("{\"msg\": \"List the directory.\"}"))
		return
	}

	w.Write([]byte("{\"msg\": \"Ok!\"}"))

	for _, fuzzerDir := range fuzzers {
		// Maybe turn this into goroutines later
		sendFiles(filepath.Join(targetPath, fuzzerDir.Name()), ip, target)
	}
}

func sendFiles(path, ip, target string) {
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

		var clientResp types.ClientResponse
		r, _ := http.Post(fmt.Sprintf("http://%s:31338/api/dropfile/%s/", ip, target), "application/json", bytes.NewBuffer(jsonValue))
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
