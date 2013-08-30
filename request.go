package memcached

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type MCRequest struct {
	Opcode  CommandCode
	Key     string
	Value   []byte
	Flags   int
	Length  int
	Expires int64
}

func (req *MCRequest) String() string {
	return fmt.Sprintf("{MCRequest opcode=%s, bodylen=%d, key='%s'}",
		req.Opcode, len(req.Value), req.Key)
}

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
