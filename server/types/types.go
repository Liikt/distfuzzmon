package types

type clientMap map[string]bool

type DropFileRequest struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

type ClientResponse struct {
	Message string `json:"msg"`
}

// Fuzzjob is a struct that contains relevant data to a fuzzjob
type Fuzzjob struct {
	Target      string   `json:"target"`
	Fuzzer      string   `json:"fuzzer"`
	FullCommand string   `json:"command"`
	FuzzerCount int      `json:"count"`
	Seeds       []string `json:"seeds"`
}

var (
	// RegisteredClients is a map that holds the information of all registered clients
	RegisteredClients clientMap
)
