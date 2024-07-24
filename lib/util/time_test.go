package util

import (
	"fmt"
	"testing"
	"time"
)

func TestStringToTime(t *testing.T) {
	timeString := "2024-07-23T18:32:40.022+0800"
	myTime, err := StringToTime(timeString)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(myTime)
	secondSub := time.Now().Unix() - myTime.Unix()

	fmt.Println(secondSub)
}
