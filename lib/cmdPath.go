// 获取可执行文件的绝对路径

package lib

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// 获取可执行文件的绝对路径
func GetAbsPath() string {
	execpath, _ := os.Executable()
	execpath = filepath.Dir(execpath) // 获得程序路径
	execpath = strings.Replace(execpath, "\\", "/", -1)
	return execpath
}

// 获取当前执行文件绝对路径（go run）
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(path.Dir(filename))
	}
	return abPath
}
