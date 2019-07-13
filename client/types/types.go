package types

// ServerRespose is a json object that represents a response from the server
type ServerRespose struct {
	Message string `json:"msg"`
}

type Fuzzjob struct {
	Target      string `json:"target"`
	Fuzzer      string `json:"fuzzer"`
	FullCommand string `json:"command"`
}
