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

//服务运行状态
var RunStats Stats

func Listen(ls net.Listener) {
	reqChannel := make(chan chanReq)

	RunStats = NewStats()

	go waitDispatch(reqChannel)
	handler := &reqHandler{reqChannel}
	for {
		s, e := ls.Accept()
		if e == nil {
			//链接数统计
			RunStats["total_connections"].(*CounterStat).Increment(1)
			RunStats["curr_connections"].(*CounterStat).Increment(1)

			go HandleIo(s, handler)
		} else {
			log.Println(RunStats)
			log.Printf("Error accepting from %s", ls)
		}
	}
}

func HandleIo(s io.ReadWriteCloser, handler *reqHandler) error {
	defer func() {
		RunStats["curr_connections"].(*CounterStat).Decrement(1)
		s.Close()
	}()

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
