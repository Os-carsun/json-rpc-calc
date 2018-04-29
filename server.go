package main

import (
	"./rpclib"
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
)

type Handler struct {
	calc rpclib.Calculator
}

func (handler *Handler) Call(obj *rpclib.RPCObj, reply *rpclib.ReplyObj) error {

	reply.ID = obj.ID
	var err error
	switch obj.Method {
	case "create":
		if len(obj.Params) != 2 {
			reply.Error = "create fail, wrong Params, need 2 Params"
			break
		}

		err = handler.calc.Create(&rpclib.Pair{obj.Params[0], obj.Params[1]})

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = "create varialbe success"
		}

	case "set":
		if len(obj.Params) != 2 {
			reply.Error = "set fail, wrong Params, need 2 Params"
			break
		}

		err = handler.calc.Update(&rpclib.Pair{obj.Params[0], obj.Params[1]})

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = "update varialbe success"
		}
	case "delete":
		if len(obj.Params) != 1 {
			reply.Error = "delete fail, wrong Params, need 1 Params"
			break
		}

		err = handler.calc.Delete(obj.Params[0])

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = "delete varialbe success"
		}
	case "plus":
		if len(obj.Params) != 2 {
			reply.Error = "plus fail, wrong Params, need 2 Params"
			break
		}

		value, err := handler.calc.DoCal(&rpclib.Pair{obj.Params[0], obj.Params[1]}, "add")

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = value.String()
		}
	case "minus":
		if len(obj.Params) != 2 {
			reply.Error = "minus fail, wrong Params, need 2 Params"
			break
		}

		value, err := handler.calc.DoCal(&rpclib.Pair{obj.Params[0], obj.Params[1]}, "sub")

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = value.String()
		}
	case "multiply":
		if len(obj.Params) != 2 {
			reply.Error = "multiply fail, wrong Params, need 2 Params"
			break
		}

		value, err := handler.calc.DoCal(&rpclib.Pair{obj.Params[0], obj.Params[1]}, "mul")

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = value.String()
		}
	case "divide":
		if len(obj.Params) != 2 {
			reply.Error = "divide fail, wrong Paramsm need 2 Params"
			break
		}

		value, err := handler.calc.DoCal(&rpclib.Pair{obj.Params[0], obj.Params[1]}, "div")

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = value.String()
		}
	default:
		reply.Error = "no such mehtod"
	}
	return nil
}

func main() {

	portNum := flag.Int("port", 1234, "server port, defualt is 1234")
	mutex := sync.Mutex{}

	flag.Parse()

	if *portNum < 1000 || *portNum > 65535 {
		fmt.Println("invliad port, use default port 1234")
		*portNum = 1234
	}

	fmt.Printf("server is run in port: %d\n", *portNum)

	handler := new(Handler)
	handler.calc.Init()

	server := rpc.NewServer()
	server.Register(handler)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	l, e := net.Listen("tcp", fmt.Sprintf(":%d", *portNum))

	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}

		go func() {
			mutex.Lock()
			server.ServeCodec(jsonrpc.NewServerCodec(conn))
			mutex.Unlock()
		}()
	}
}
