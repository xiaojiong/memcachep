package main

import (
	"github.com/xiaojiong/memcachep"
)

//初始化绑定处理程序
func init() {
	memcachep.BindHandler(memcachep.GET, handleGet)
}

func handleGet(req *memcachep.MCRequest, res *memcachep.MCResponse) {
	res.Fatal = false
	res.Value = []byte("aaabbbccc" + req.Key)
}
