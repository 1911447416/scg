package main

import (
	"fmt"
	util "sourcecode/Util"
)

func main() {
	fmt.Println("start...")
	ArgsModel, _ := util.Parse()
	// fmt.Println(ArgsModel.ClassName)
	if ArgsModel.ClassName != "" {
		util.FileForEach(ArgsModel)
	} else {
		fmt.Println("add -help to know how to use.")
	}

}
