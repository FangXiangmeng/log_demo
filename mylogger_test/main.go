package main

import (
	"time"

	"captain.github.com/studygo/job/mylogger"
)

// 测试我们自己写的日志库
func main() {
	log := mylogger.NewLog("INFO") //这个是输出到控制台
	//log := mylogger.NewFileLogger("INFO", "./", "fangxiangmeng.log", 10*1024*1024)  //这个是输出到文件
	//	for {
	log.Debug("这是一条Debug日志")
//	log.Info("这是一条Info日志")
//	log.Warning("这是一条Warning日志")
//	id := 10010
//	name := "fxm"
//	log.Error("这是一条Error日志,id:%d,name:%s", id, name) //使用interface传多个值
//	log.Fatal("这是一条Fatal日志")
	time.Sleep(time.Second * 3)
	//	}
}
