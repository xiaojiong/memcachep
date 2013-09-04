package main

import (
	"github.com/xiaojiong/memcachep"

	"fmt"
	"log"
	"net"
)

func main() {
	port := 11345
	ls, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if e != nil {
		log.Fatalf("Got an error:  %s", e)
	}
	memcachep.Listen(ls)
}
