package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config is a struct that holds all config options available for
type Config struct {
	BaseFolder string `json:"base_folder"`
	ServerIP   string `json:"serverip"`
}

// LoadConfig takes in a path to a config file and returns a config object
func LoadConfig(configPath string) Config {
	var conf Config
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("[-] Couldn't read the config file.")
		os.Exit(1)
	}
	err = json.Unmarshal(content, &conf)
	if err != nil {
		fmt.Println("[-] Couldn't parse the config file.", err)
		os.Exit(1)
	}

	return conf
}
