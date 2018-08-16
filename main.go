package main

import (
	"fmt"
	"utils/file"
)

func main(){
	//文件夹遍历
	dir := "/Users/yuan/www/html/liebao/go/src/utils"
	files, err := file.ScanDir(dir)
	if err != nil{
		fmt.Printf("ScanDir: %s is not a directory\n", err.Error())
	}

	for _,file := range files{
		fmt.Printf("ScanDir: res: %s\n", file)
	}

}