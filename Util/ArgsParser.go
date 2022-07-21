package util

import (
	"flag"
	"fmt"
)

// var cfg *ini.File

type ArgsModel struct {
	ClassName    string
	TemplatePath string
	OutputPath   string
}

// 读配置文件
// func Readconfig() *ini.File {
// cfg, err := ini.Load("./conf.ini")
// if err != nil {
// 	fmt.Println("fail to raad config,err:", err)
// 	os.Exit(1)
// }
// return cfg
// }

func IsPath(OutputPath *string, TemplatePath *string, ClassName *string) bool {
	// 使用静态配置
	if *OutputPath == "" {
		// *OutputPath = cfg.Section("").Key("OutputPath").String()
		fmt.Println("you ought to add OutputPath. usage: -o ./newfiles/")
		return false
	}

	if *TemplatePath == "" {
		// *TemplatePath = cfg.Section("").Key("TemplatePath").String()
		fmt.Println("you ought to add TemplatePath. usage: -t ./TemplatePath/")
		return false
	}

	if *ClassName == "" {
		// *TemplatePath = cfg.Section("").Key("TemplatePath").String()
		fmt.Println("you ought to add ClassName. usage: -c ClassName or Enter ClassName at the end of the sentence")
		return false
	}
	O := *OutputPath
	T := *TemplatePath
	if (O[len(O)-1:]) != "/" {
		*OutputPath = O + "/"
	}
	if (T[len(T)-1:]) != "/" {
		*TemplatePath = T + "/"
	}
	return true

}

// 解释命令行参数 返回ArgsModel
func Parse() (ArgsModel, error) {
	var ClassName = flag.String("c", "", "class name")
	var TemplatePath = flag.String("t", "", "template path")
	var OutputPath = flag.String("o", "", "output path")
	// cfg := Readconfig()
	//解析参数
	flag.Parse()
	if len(flag.Args()) != 0 { // 有没设flag
		fmt.Println("--->", flag.Args())
		ClassName = &flag.Args()[0]
	}
	if !IsPath(OutputPath, TemplatePath, ClassName) {
		return ArgsModel{}, nil
	}
	return ArgsModel{
		ClassName:    *ClassName,
		TemplatePath: *TemplatePath,
		OutputPath:   *OutputPath,
	}, nil
}
