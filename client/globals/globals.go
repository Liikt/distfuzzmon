package globals

import (
	"distfuzzmon/client/config"
	"distfuzzmon/client/utils"
)

var (
	// Conf is the global config file for this client
	Conf config.Config
	// IP is the ip address the client is bound to
	IP string
)

// SetupGlobals initializes all global variables
func SetupGlobals(configPath string) {
	IP = utils.FindIP()
	Conf = config.LoadConfig(configPath)
}
