package memcachep

import (
	"fmt"
	"io"
)

type MCResponse struct {
	//命令
	Opcoed CommandCode
	//返回状态
	Status Status
	//key
	Key string
	//返回内容
	Value []byte
	//返回标识
	Flags int
	//错误
	Fatal bool
}

//解析response 并把返回结果写入socket链接
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
