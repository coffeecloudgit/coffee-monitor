package test

import (
	"fmt"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	c := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		time.Sleep(time.Second)
		total += v
	}
	c <- total // send total to c
}
func TestChannel(t *testing.T) {
	a := []int{8, 8, 8, 8}

	b := []int{8, 8, 8, 8, 8, 8, 8}

	d := []int{8, 8, 8, 8, 8, 8, 8}

	c := make(chan int)
	go sum(a, c)
	go sum(b, c)
	go sum(d, c)

	x := <-c
	fmt.Println(time.Now(), ",", x)
	y := <-c
	fmt.Println(time.Now(), ",", y)

	z := <-c
	fmt.Println(time.Now(), ",", z)

	fmt.Println(x, y, x+y)
}

func Test04(t *testing.T) {
	// 原始切片
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 截取切片的起始索引和结束索引
	start := 3
	end := 7

	// 截取切片
	subSlice := originalSlice[start:end]

	fmt.Println("Original Slice:", originalSlice)
	fmt.Println("Sub Slice:", subSlice)
}
