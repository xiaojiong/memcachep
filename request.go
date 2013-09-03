package memcachep

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

//mc请求产生一个request对象
type MCRequest struct {
	//请求命令
	Opcode CommandCode
	//key
	Key string
	//请求内容
	Value []byte
	//请求标识
	Flags int
	//请求内容长度
	Length int
	//过期时间
	Expires int64
}

//request to string
func (req *MCRequest) String() string {
	return fmt.Sprintf("{MCRequest opcode=%s, bodylen=%d, key='%s'}",
		req.Opcode, len(req.Value), req.Key)
}

//将socket请求内容 解析为一个MCRequest对象
func (req *MCRequest) Receive(r *bufio.Reader) error {
	line, _, err := r.ReadLine()
	if err != nil || len(line) == 0 {
		return io.EOF
	}

	params := strings.Fields(string(line))

	command := CommandCode(params[0])

	switch command {
	case SET, ADD, REPLACE:
		req.Opcode = command
		req.Key = params[1]
		req.Length, _ = strconv.Atoi(params[4])

		value := make([]byte, req.Length+2)
		io.ReadFull(r, value)

		req.Value = make([]byte, req.Length)
		copy(req.Value, value)

	case GET:
		log.Println(string(line))
		req.Opcode = command
		req.Key = params[1]
	case STATS:
		req.Opcode = command
	}

	return err
}
