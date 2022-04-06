package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj,err := os.OpenFile("xx.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Println("文件打开失败")
	}

	log.SetOutput(fileObj) //将日志输出到文件
	for {
		log.Println("这是一条测试日志")
		time.Sleep(time.Second * 3)
	}
}
