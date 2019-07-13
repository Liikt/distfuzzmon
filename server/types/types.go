package types

type clientMap map[string]bool

type DropFileRequest struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

var (
	// RegisteredClients is a map that holds the information of all registered clients
	RegisteredClients clientMap
)
