package types

type clientMap map[string]bool

var (
	// RegisteredClients is a map that holds the information of all registered clients
	RegisteredClients clientMap
)
