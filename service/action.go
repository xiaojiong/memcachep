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

	v := ""
	for key, value := range a {
		v += fmt.Sprintf("STAT %s %s\r\n", key, value.String())
	}
	v += "END/r/n"

	res.Value = []byte("STAT pid 1876\r\nSTAT bytes_written 410\r\nSTAT limit_maxbytes 67108864\r\nSTAT threads 1\r\nEND\r\n")
}

func SetAction(req *memcachep.MCRequest, res *memcachep.MCResponse) {

	res.Fatal = false

	res.Status = memcachep.STORED
}
