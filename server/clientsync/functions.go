package clientsync

import (
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
		w.Write([]byte("{\"error\": \"Couldn't parse the body of a file.\"}"))
		return
	}
	content, err := base64.StdEncoding.DecodeString(reqContent.Content)
	if err != nil {
		fmt.Printf("[!] Error decoding %s\n", reqContent.Content)
		w.WriteHeader(500)
		w.Write([]byte("{\"error\": \"Couldn't decode the body of a file.\"}"))
		return
	}

	fuzzerDir := filepath.Join("sync_dir", target, fuzzer)
	if exists, err := utils.PathExists(fuzzerDir); err != nil {
		fmt.Println("[-] Something went wrong while stating the sync dir")
		os.Exit(1)
	} else if !exists {
		os.MkdirAll(fuzzerDir, 0755)
	}

	ioutil.WriteFile(filepath.Join(fuzzerDir, reqContent.Filename), content, 0644)
}
