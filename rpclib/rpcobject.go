package rpclib

type RPCObj struct {
	Method  string   `json: "method"`
	JsonRPC string   `json: "jsonrpc"`
	Params  []string `json: "params"`
	ID      int      `json: "id"`
}

type ReplyObj struct {
	Result string `json: "result"`
	Error  string `json: "error"`
	ID     int    `json: "id"`
}
