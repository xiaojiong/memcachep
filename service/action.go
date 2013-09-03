package main

import (
	"fmt"
	"github.com/xiaojiong/memcachep"
)

var data map[string]string

//初始化绑定处理程序
func init() {
	memcachep.BindAction(memcachep.STATS, StatsAction)
	memcachep.BindAction(memcachep.GET, GetAction)
	memcachep.BindAction(memcachep.SET, SetAction)

}

func GetAction(req *memcachep.MCRequest, res *memcachep.MCResponse) {
	res.Fatal = false
	//res.Value = []byte("aaabbbccc" + req.Key)

	res.Value = []byte(data[req.Key])

}

func StatsAction(req *memcachep.MCRequest, res *memcachep.MCResponse) {
	res.Fatal = false

	res.Value = []byte(fmt.Sprintf("STAT %s 4\r\nEND/r/n", "num"))
}

func SetAction(req *memcachep.MCRequest, res *memcachep.MCResponse) {

	res.Fatal = false

	res.Status = memcachep.STORED
}
