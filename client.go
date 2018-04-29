package main

import (
	"./rpclib"
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func asyncReq(serverIP *string, serverPort *int) {
	client, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", *serverIP, *serverPort), 1000*1000*1000*30)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	c := jsonrpc.NewClient(client)

	endChan := make(chan int, 15)

	method := []string{"create", "set", "delete", "plus", "minus", "multiply", "divide"}
	for i := 1; i <= 15; i++ {
		request := &rpclib.RPCObj{
			method[i%len(method)],
			"1.0",
			[]string{"test", "1234"},
			i}

		log.Println("client\t-", "call create method")

		asyncCall := c.Go("Handler.Call", request, &rpclib.ReplyObj{}, nil)

		go func(num int) {
			reply := <- asyncCall.Done
			obj := (reply.Reply).(*rpclib.ReplyObj)
			log.Println("clent\t-", "recive response:{ ID: ", obj.ID, ", result: ", obj.Result, ", error:", obj.Error, "} id should be :", num)
			endChan <- num
		}(i)
	}

	for i := 1; i <= 15; i++ {
		_ = <-endChan
	}
}

func main() {

	serverIP := flag.String("IP", "127.0.0.1", "rpc serverIP, default is 127.0.0.1")
	serverPort := flag.Int("port", 1234, "rpc port, default is 1234")
	asyncReq(serverIP, serverPort)
}
