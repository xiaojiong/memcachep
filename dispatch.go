package memcachep

import (
	"fmt"
)

type action func(req *MCRequest, res *MCResponse)

var actions = map[CommandCode]action{
	STATS:StatsAction
}

//等待分发处理
func waitDispatch(rc chan chanReq) {
	for {
		input := <-rc

		input.response <- dispatch(input.request)
	}
}

//分发请求到响应的action操作函数上去
func dispatch(req *MCRequest) (res *MCResponse) {
	if h, ok := actions[req.Opcode]; ok {
		res = &MCResponse{}
		h(req, res)
	} else {
		return notFound(req)
	}
	return
}

//未支持命令
func notFound(req *MCRequest) *MCResponse {
	var response MCResponse
	response.Status = UNKNOWN_COMMAND
	return &response
}

//给request绑定上处理程序
func BindAction(opcode CommandCode, h action) {
	actions[opcode] = h
}

//stats
func StatsAction(req *MCRequest, res *MCResponse) {
	res.Fatal = false

	stats := ""
	for key, value := range RunStats {
		stats += fmt.Sprintf("STAT %s %s\r\n", key, value)
	}
	stats += "END\r\n"
	res.Value = []byte(stats)
}
