package util

import (
	"os"
	"path/filepath"
)

//CWD 获取程序当前目录
func CWD() string {
	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(path)
}

//Exist 判断文件是否存在
func Exist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}