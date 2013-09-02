package memcachep

import (
	"fmt"

	"log"
	"net"
	"testing"
)


func Test_T1(t *testing.T) {
	port := 11345
	ls, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if e != nil {
		log.Fatalf("Got an error:  %s", e)
	}

	waitForConnections(ls)
}
