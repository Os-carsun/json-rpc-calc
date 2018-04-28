package main

import (
	"./rpclib"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Handler struct {
	calc rpclib.Calculator
}

func (handler *Handler) Call(obj *rpclib.RPCObj, replay *rpclib.ReplyObj) error {

	replay.ID = obj.ID

	switch obj.Method {
	case "create":
		if len(obj.Params) != 2 {
			return errors.New("wrong Params")
		}
		replay.Result = "create varialbe success"
		return handler.calc.Create(&rpclib.Pair{obj.Params[0], obj.Params[1]})
	case "updating":
		if len(obj.Params) != 2 {
			return errors.New("wrong Params")
		}
		replay.Result = "update varialbe success"
		return handler.calc.Update(&rpclib.Pair{obj.Params[0], obj.Params[1]})
	case "delete":
		if len(obj.Params) != 1 {
			return errors.New("wrong Params")
		}
		replay.Result = "delete varialbe success"
		return handler.calc.Delete(obj.Params[0])
	case "addition":
	case "subtraction":
	case "multiplication":
	case "division":
	default:
		return errors.New("no such method")
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
