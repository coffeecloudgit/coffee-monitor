package util

import (
	"errors"
	"fmt"
	"strconv"
)

func InterfaceToUnit64(data interface{}) (uint64, error) {
	if num, ok := data.(uint64); ok {
		fmt.Println("转换成功:", num)
		return num, nil
	} else {
		// 如果data不是uint64类型，尝试将其转换为字符串然后解析为uint64
		if str, ok := data.(string); ok {
			if n, err := strconv.ParseUint(str, 10, 64); err == nil {
				fmt.Println("解析字符串转换成功:", n)
				return num, nil
			} else {
				fmt.Println("转换失败:", err)
				return 0, err
			}
		} else {
			fmt.Println("转换失败: 不是uint64类型也不是字符串")
			return 0, errors.New("转换失败: 不是uint64类型也不是字符串")
		}
	}
}
