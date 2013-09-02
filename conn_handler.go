package memcachep

import (
	"bufio"
	"io"
	"log"
	"net"
)


type chanReq struct {
	request *MCRequest
	response chan *MCResponse
}

type reqHandler struct {
	ch chan chanReq
}

func waitForConnections(ls net.Listener) {
	reqChannel := make(chan chanReq)
	go RunServer(reqChannel)
	handler := &reqHandler{reqChannel}
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

func HandleIo(s io.ReadWriteCloser, handler *reqHandler) error {
	defer s.Close()
	var err error
	for err == nil {
		err = HandleMessage(bufio.NewReader(s), s, handler)
	}
	return err
}

func HandleMessage(r *bufio.Reader, w io.Writer, handler *reqHandler) error {
	req, err := ReadPacket(r)

	if err != nil {
		return err
	}

	cr := chanReq{
		req,
		make(chan *MCResponse),
	}
	handler.chan <- cr

	res := <-cr.response

	if !res.Fatal {
		res.Opcoed = req.Opcode
		res.Key = req.Key

		err = res.Transmit(w)
		if err != nil {
			return err
		}
		return nil
	}

	return io.EOF
}

func ReadPacket(r *bufio.Reader) (*MCRequest, error) {
	req := &MCRequest{}
	err := req.Receive(r)
	return req, err
}

//协议service的action处理方法



type handler func(req *MCRequest, res *MCResponse)

var handlers = map[CommandCode]handler{}

func RunServer(handler *chanReq) {
	for {
		req := <-handler.request

		handler.response <- dispatch(req)s
	}
}

func dispatch(req *MCRequest) (res *MCResponse) {
	if h, ok := handlers[req.Opcode]; ok {
		res = &MCResponse{}
		h(req, res)
	} else {
		return notFound(req)
	}
	return
}

//未支持命令
func notFound(req *MCRequest) *MCResponse {
	var response MCResponse
	response.Status = UNKNOWN_COMMAND
	return &response
}

//给处理程序绑定上 handler
func BindHandler(opcode CommandCode, h handler) {
	handlers[opcode] = h
}
