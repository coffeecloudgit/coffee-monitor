package util

import (
	"fmt"
	"strconv"
	"testing"
)

func TestQueue_Get(t *testing.T) {
	q := Queue{
		Content: []Object{},
		Timeout: 0,
		MaxSize: 4,
	}
	go func() {
		for i := 0; i < 50; i++ {
			st := "hello" + strconv.Itoa(i)
			q.Put(st)
			fmt.Println(st)
		}
	}()
	for i := 0; i < 50; i++ {
		fmt.Println("read: ", q.Get(), ",size:", len(q.Content))
	}
}
