package memcachep

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type Stats map[string]fmt.Stringer

//动态函数 status
type FuncStat struct {
	Callable func() string
}

func (f *FuncStat) String() string {
	return f.Callable()
}

//静态数据 status
type StaticStat struct {
	Value string
}

func (s *StaticStat) String() string {
	return s.Value
}

//线程安全计数器 status
type CounterStat struct {
	Count        int
	calculations chan int
}

func (c *CounterStat) Increment(num int) {
	c.calculations <- num
}

func (c *CounterStat) SetCount(num int) {
	c.Count = num
}

func (c *CounterStat) Decrement(num int) {
	c.calculations <- -num
}

func (c *CounterStat) String() string {
	return strconv.Itoa(c.Count)
}

func (c *CounterStat) work() {
	for num := range c.calculations {
		c.Count = c.Count + num
	}
}

func NewCounterStat() *CounterStat {
	c := &CounterStat{}
	c.calculations = make(chan int, 100)
	go c.work()
	return c
}

func NewStats() Stats {
	s := make(Stats)
	s["pid"] = &StaticStat{strconv.Itoa(os.Getpid())}
	s["version"] = &StaticStat{VERSION}
	s["golang"] = &StaticStat{runtime.Version()}
	//运行go程 数量
	s["goroutines"] = &FuncStat{func() string { return strconv.Itoa(runtime.NumGoroutine()) }}
	s["cpu_num"] = &StaticStat{strconv.Itoa(runtime.NumCPU())}
	s["total_connections"] = NewCounterStat()
	s["curr_connections"] = NewCounterStat()
	s["cmd_get"] = NewCounterStat()
	return s
}
