package memcached

import (
	"fmt"

	"log"
	"net"
	"testing"
)

func waitForConnections(ls net.Listener) {
	handler := NewRequestHandler()
	go RunServer(handler)

	for {
		s, e := ls.Accept()
		if e == nil {
			log.Printf("Got a connection from %v", s.RemoteAddr())

			go HandleIo(s, handler)
		} else {
			log.Printf("Error accepting from %s", ls)
		}
	}
}

func Test_T1(t *testing.T) {
	port := 11345
	ls, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if e != nil {
		log.Fatalf("Got an error:  %s", e)
	}

	waitForConnections(ls)
}

//memcache 协议后的函数处理

type handler func(req *MCRequest) *MCResponse

var handlers = map[CommandCode]handler{
	GET: handleGet,
}

func RunServer(handler *RequestHandler) {
	for {
		req := <-handler.request

		handler.response <- dispatch(req)
	}
}

func dispatch(req *MCRequest) (res *MCResponse) {
	if h, ok := handlers[req.Opcode]; ok {
		res = h(req)
	} else {
		return notFound(req)
	}
	return
}

func handleGet(req *MCRequest) (res *MCResponse) {
	res = &MCResponse{}
	res.Fatal = false
	res.Value = []byte("aaabbbccc" + req.Key)

	return
}

func notFound(req *MCRequest) *MCResponse {
	var response MCResponse
	response.Status = UNKNOWN_COMMAND
	return &response
}
