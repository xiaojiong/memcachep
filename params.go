package memcachep

import (
	"fmt"
)

const VERSION = "0.0.1"

//request command type
type CommandCode string

//response status type
type Status int

//status result map
var StatusRes map[Status]string

//response 状态
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

//request command
const (
	GET     CommandCode = "get"
	SET     CommandCode = "set"
	ADD     CommandCode = "add"
	REPLACE CommandCode = "replace"
	DELETE  CommandCode = "delete"
	STATS   CommandCode = "stats"
)

//初始化response状态的返回结果
func init() {
	StatusRes = make(map[Status]string)
	StatusRes[ERROR] = "ERROR"
	StatusRes[STORED] = "STORED"
	StatusRes[NOT_STORED] = "NOT_STORED"
	StatusRes[END] = "END"
	StatusRes[DELETED] = "DELETED"
	StatusRes[NOT_FOUND] = "NOT_FOUND"
}

//status to string
func (s *Status) ToString() string {
	rv := StatusRes[*s]
	if rv == "" {
		rv = fmt.Sprintf("%s\r\n", StatusRes[NOT_FOUND])
	} else {
		rv = fmt.Sprintf("%s\r\n", rv)
	}
	return rv
}
