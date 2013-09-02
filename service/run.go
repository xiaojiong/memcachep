package main

import (
	"fmt"

	"github.com/xiaojiong/memcachep"
	"log"
	"net"
)

func main() {
	port := 11345
	ls, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if e != nil {
		log.Fatalf("Got an error:  %s", e)
	}

	waitForConnections(ls)
}

func waitForConnections(ls net.Listener) {
	handler := memcachep.NewRequestHandler()
	go memcachep.RunServer(handler)

	for {
		s, e := ls.Accept()
		if e == nil {
			log.Printf("Got a connection from %v", s.RemoteAddr())
			go memcachep.HandleIo(s, handler)
		} else {
			log.Printf("Error accepting from %s", ls)
		}
	}
}
