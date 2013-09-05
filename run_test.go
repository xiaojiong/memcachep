package memcachep

import (
	"fmt"

	"log"
	"net"
	"testing"
	"time"
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

	time.Sleep(1 * time.Second)

	res.Value = []byte("{\"ResList\":[{\"File\":\"csdn.txt\",\"Strings\":[\"qq290747680 # 123456789 # 290747680@qq.com\\r\"],\"Num\":1,\"EndTime\":0,\"StartTime\":0,\"RunTime\":0}],\"Counter\":{\"Num\":1}}")
}

func SetAction(req *MCRequest, res *MCResponse) {
	res.Fatal = false
	data[req.Key] = string(req.Value)
	res.Status = STORED
}
