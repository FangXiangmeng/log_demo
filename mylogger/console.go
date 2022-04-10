package mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type ConsoleLogger struct {
	Level LogLevel   //实际上是unit16只不过是使用了类型别名
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s) //将传进来的s参数转换为小写
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")  //使用字符串创建一个错误，这样如果代码执行不下去就报err
		return UNKNOWN, err
	}
}

func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err) //如果到这里err不等于nil直接退出程序
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	// str := LogLevel >= c.Level
	// fmt.Println(str)
	fmt.Println(LogLevel,c.Level) //1,3 返回false 就不会打印日志。
	return LogLevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)

}

func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
