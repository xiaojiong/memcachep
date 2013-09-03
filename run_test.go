package memcachep

import (
	"fmt"

	"log"
	"net"
	"testing"
)

func Test_t1(t *testing.T) {
	port := 11345
	ls, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if e != nil {
		log.Fatalf("Got an error:  %s", e)
	}
	Listen(ls)
}

var data map[string]string = make(map[string]string)

//初始化绑定处理程序
func init() {
	BindAction(STATS, StatsAction)
	BindAction(GET, GetAction)
	BindAction(SET, SetAction)
}

func GetAction(req *MCRequest, res *MCResponse) {
	res.Fatal = false
	res.Value = []byte(data[req.Key])
}

func SetAction(req *MCRequest, res *MCResponse) {
	res.Fatal = false
	data[req.Key] = string(req.Value)
	res.Status = STORED
}
