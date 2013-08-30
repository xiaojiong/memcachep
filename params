package memcached

import (
	"fmt"
)

type CommandCode string
type Status int

var StatusRes map[Status]string

const (
	SUCCESS Status = iota
	ERROR
	STORED
	NOT_STORED
	END
	DELETED
	NOT_FOUND
	UNKNOWN_COMMAND
)

const (
	GET     CommandCode = "get"
	SET     CommandCode = "set"
	ADD     CommandCode = "add"
	REPLACE CommandCode = "replace"
	DELETE  CommandCode = "delete"
	STATS   CommandCode = "stats"
)

func init() {
	StatusRes = make(map[Status]string)
	StatusRes[ERROR] = "ERROR"
	StatusRes[STORED] = "STORED"
	StatusRes[NOT_STORED] = "NOT_STORED"
	StatusRes[END] = "END"
	StatusRes[DELETED] = "DELETED"
	StatusRes[NOT_FOUND] = "NOT_FOUND"
}

func (s *Status) ToString() string {
	rv := StatusRes[*s]
	if rv == "" {
		rv = fmt.Sprintf("%s\r\n", StatusRes[NOT_FOUND])
	} else {
		rv = fmt.Sprintf("%s\r\n", rv)
	}
	return rv
}
