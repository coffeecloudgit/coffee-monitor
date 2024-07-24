package util

import (
	"sync"
	"time"
)

// Object 数据结构
type Object interface{}

type Queue struct {
	Content []Object
	Timeout int // timeout为0为无限延时， 小于0为不延时， 大于0为延时timeout秒
	MaxSize int // 队列容量, 小于或等于0为不限量，不限量时延时无效，大于0且到达上限时则开始延时
}

var lock = sync.Mutex{}

// Put 超过设定延时时间后， 元素会被抛弃
func (q *Queue) Put(msg Object) {
	lock.Lock()
	closeSingle := make(chan bool)
	successSingle := make(chan bool)
	go func(close chan bool, success chan bool) {
		var t1 *time.Timer
		t1 = time.NewTimer(time.Second * time.Duration(q.Timeout))
		for {
			select {
			case <-t1.C:
				if q.Timeout == 0 {
					t1 = time.NewTimer(time.Second * time.Duration(q.Timeout))
					continue
				} else {
					success <- true
					return
				}
			default:
				if q.MaxSize != 0 && q.MaxSize == len(q.Content) {
					continue
				} else {
					q.Content = append(q.Content, msg)
					success <- true
					return
				}
			}
		}
	}(closeSingle, successSingle)

	for {
		select {
		case <-successSingle:
			lock.Unlock()
			return
		}
	}
}

// Get 超过延时时间时会返回空字符串
func (q *Queue) Get() Object {
	closeSingle := make(chan bool)
	successSingle := make(chan Object)
	go func(close chan bool, output chan Object) {
		var t1 *time.Timer
		t1 = time.NewTimer(time.Second * time.Duration(q.Timeout))
		for {
			select {
			case <-t1.C:
				if q.Timeout == 0 {
					t1 = time.NewTimer(time.Second * time.Duration(q.Timeout))
					continue
				} else {
					output <- ""
					return
				}
			default:
				if 0 == len(q.Content) {
					continue
				} else {
					msg := q.Content[0]
					q.Content = q.Content[1:]
					output <- msg
					return
				}
			}
		}
	}(closeSingle, successSingle)

	for {
		select {
		case res := <-successSingle:
			return res
		}
	}
}
