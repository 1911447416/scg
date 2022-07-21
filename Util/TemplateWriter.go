package util

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 模板替换
func strreplace(s string, new string) string {
	newstring := strings.Replace(s, "<%ClassName%>", new, -1)
	return newstring
}

// 通用的写文本代码
func Writefile(path string, content string) {
	// 打开输出的路径
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("writefile. patherr failed,err:", err, "or there is a file with the same name")
	}
	// 输出
	writer := bufio.NewWriter(file)
	_, w_err := writer.WriteString(content)
	writer.Flush()
	if w_err != nil {
		fmt.Println("w_err failed. err:", w_err)
	}
}

// 替换每个文件
func FileForEach(ArgsModel ArgsModel) []fs.FileInfo {
	files, err := ioutil.ReadDir(ArgsModel.TemplatePath)
	if err != nil {
		fmt.Println("readdir err,err:", err)
	}
	var myFile []fs.FileInfo
	for _, file := range files {
		// fmt.Println(file.Name())
		// 循环替换

		// 跳过目录
		if file.IsDir() {
			continue
		}
		// 跳过指定后缀名的文件
		suffix := filepath.Ext(file.Name())
		if suffix != "" {
			suffix = strings.ToLower(suffix[1:])
		}
		fmt.Println(suffix)
		if suffix == "mp3" || suffix == "mp4" || suffix == "wav" {
			continue
		}
		// fmt.Println("suffix", suffix)

		editfile := TemplateRead(ArgsModel.TemplatePath + "/" + file.Name())
		editfile = strreplace(editfile, ArgsModel.ClassName)
		filename := strreplace(file.Name(), ArgsModel.ClassName)
		fmt.Println("new file, filename:", filename)

		// 判断输出目录是否存在
		_, err := os.Stat(ArgsModel.OutputPath)
		if os.IsNotExist(err) {
			os.Mkdir(ArgsModel.OutputPath, os.ModeDir|os.ModePerm)
		}
		Writefile(ArgsModel.OutputPath+filename, editfile)

		if file.IsDir() {
			continue
		}
		myFile = append(myFile, file)
	}
	fmt.Println("New files is in outputpath : ", ArgsModel.OutputPath)

	return myFile
}
