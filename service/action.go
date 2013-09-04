package main

import (
	"github.com/xiaojiong/memcachep"
)

var data map[string]string = make(map[string]string)

//初始化绑定处理程序
func init() {
	memcachep.BindAction(memcachep.GET, GetAction)
	memcachep.BindAction(memcachep.SET, SetAction)
}

func GetAction(req *memcachep.MCRequest, res *memcachep.MCResponse) {
	res.Fatal = false
	res.Value = []byte(data[req.Key])
}

func SetAction(req *memcachep.MCRequest, res *memcachep.MCResponse) {
	res.Fatal = false
	data[req.Key] = string(req.Value)
	res.Status = memcachep.STORED
}
