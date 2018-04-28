package main

import (
	"./rpclib"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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
			reply.Error = "wrong Params"
		}

		err = handler.calc.Create(&rpclib.Pair{obj.Params[0], obj.Params[1]})

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = "create varialbe success"
		}

	case "updating":
		if len(obj.Params) != 2 {
			reply.Error = "wrong Params"
		}

		err = handler.calc.Update(&rpclib.Pair{obj.Params[0], obj.Params[1]})

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = "update varialbe success"
		}
	case "delete":
		if len(obj.Params) != 1 {
			reply.Error = "wrong Params"
		}

		err = handler.calc.Delete(obj.Params[0])

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = "delete varialbe success"
		}
	case "addition":
		if len(obj.Params) != 2 {
			reply.Error = "wrong Params"
		}

		value, err := handler.calc.DoCal(&rpclib.Pair{obj.Params[0], obj.Params[1]}, "add")

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = value.String()
		}
	case "subtraction":
		if len(obj.Params) != 2 {
			reply.Error = "wrong Params"
		}

		value, err := handler.calc.DoCal(&rpclib.Pair{obj.Params[0], obj.Params[1]}, "sub")

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = value.String()
		}
	case "multiplication":
		if len(obj.Params) != 2 {
			reply.Error = "wrong Params"
		}

		value, err := handler.calc.DoCal(&rpclib.Pair{obj.Params[0], obj.Params[1]}, "mul")

		if err != nil {
			reply.Error = err.Error()
		} else {
			reply.Result = value.String()
		}
	case "division":
		if len(obj.Params) != 2 {
			reply.Error = "wrong Params"
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

	handler := new(Handler)
	handler.calc.Init()
	server := rpc.NewServer()

	server.Register(handler)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	l, e := net.Listen("tcp", ":1234")

	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
