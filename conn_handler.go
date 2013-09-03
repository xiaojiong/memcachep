package memcachep

import (
	"bufio"
	"io"
	"log"
	"net"
)

//一个 请求产生 一个 chanReq
type chanReq struct {
	request  *MCRequest
	response chan *MCResponse
}

//请求句柄
type reqHandler struct {
	ch chan chanReq
}

var RunStats Stats

func Listen(ls net.Listener) {
	reqChannel := make(chan chanReq)

	RunStats := NewStats()

	go waitDispatch(reqChannel)
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
	log.Println(req.String())

	cr := chanReq{
		req,
		make(chan *MCResponse),
	}
	handler.ch <- cr

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

type action func(req *MCRequest, res *MCResponse)

var actions = map[CommandCode]action{}

//等待分发处理
func waitDispatch(rc chan chanReq) {
	for {
		input := <-rc

		input.response <- dispatch(input.request)
	}
}

//分发请求到响应的action操作函数上去
func dispatch(req *MCRequest) (res *MCResponse) {
	if h, ok := actions[req.Opcode]; ok {
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

//给request绑定上处理程序
func BindAction(opcode CommandCode, h action) {
	actions[opcode] = h
}
