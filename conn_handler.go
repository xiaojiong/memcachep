package memcachep

import (
	"bufio"
	"io"
	"log"
)

//请求句柄
type RequestHandler struct {
	request  chan *MCRequest
	response chan *MCResponse
}

//init请求句柄结构
func NewRequestHandler() *RequestHandler {
	return &RequestHandler{request: make(chan *MCRequest), response: make(chan *MCResponse)}
}

func HandleIo(s io.ReadWriteCloser, handler *RequestHandler) error {
	defer s.Close()
	var err error
	for err == nil {
		err = HandleMessage(bufio.NewReader(s), s, handler)
	}
	return err
}

func HandleMessage(r *bufio.Reader, w io.Writer, handler *RequestHandler) error {
	req, err := ReadPacket(r)

	if err != nil {
		return err
	}

	handler.request <- req
	log.Println("request: ", req.String())

	res := <-handler.response
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

func RunServer(handler *RequestHandler) {
	for {
		req := <-handler.request

		handler.response <- dispatch(req)
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
