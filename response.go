package memcachep

import (
	"fmt"
	"io"
)

type MCResponse struct {
	Opcoed CommandCode
	Status Status
	Key    string
	Value  []byte
	Flags  int
	Fatal  bool
}

func (res *MCResponse) Transmit(w io.Writer) (err error) {
	if res.Status == SUCCESS {
		switch res.Opcoed {
		case GET:
			if res.Status == SUCCESS {
				rs := fmt.Sprintf("VALUE %s %d %d\r\n%s\r\nEND\r\n", res.Key, res.Flags, len(res.Value), res.Value)
				_, err = w.Write([]byte(rs))
			} else {
				_, err = w.Write([]byte(res.Status.ToString()))
			}
		case SET, REPLACE:
			_, err = w.Write([]byte(res.Status.ToString()))
		}

	}
	return
}
