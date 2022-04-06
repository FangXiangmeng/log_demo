package main

import (
	"fmt"
	"path"
	"runtime"
)


func f(){
	pc, file, line, ok := runtime.Caller(1) //runtime主要是记录堆栈信息。例如函数调用。这次使用主要是取函数名。
	//ok 如果能够去到ok就是true，file文件名，line runtime所在行号。
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(pc)
	fmt.Println(path.Base(file)) //main.go
	fmt.Println(line) //11
}

func f1(){
	f()
}

func main() {
	f1()
}
