package types

// SetupGlobals sets up the global variables
func SetupGlobals() {
	RegisteredClients = make(clientMap)
}
