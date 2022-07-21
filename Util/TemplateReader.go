package util

import (
	"fmt"
	"os"
)

// 读取模板 返回文件的内容
func TemplateRead(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("openfile failed,err:", err)
		return "nil"
	}
	files := string(file)
	return files
}
