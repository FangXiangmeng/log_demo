package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写文件相关代码

type FileLogger struct {
	Level       LogLevel
	filePath    string //日志文件保存路径
	fileName    string //日志文件保存的文件名
	fileObj     *os.File //
	errFileObj  *os.File
	errFileName string
	maxFileSize int64
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开和创建
	if err != nil {
		panic(err)
	}
	return fl
}

//判断文件是否存在，如果不存在就创建，如果存在就追加
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)  // ./和fangxiangmeng.log 加在一起就是./fangxiangmeng.log  Join主要是路径拼接
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //创建gxiangmeng.log.err.log文件
	if err != nil {
		fmt.Printf("open err log file failed,err:%v", err)
		return err
	}

	//日志文件都已经打开了，
	f.fileObj = fileObj //给f这个结构体中的fileObj对象赋值
	f.errFileObj = errFileObj  //给f这个结构体中的errFileObj对象赋值
	return nil //返回一个空
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed ,err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())      //拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) //拼接一个日志备份的名字
	// 1.关闭当前的日志文件
	file.Close()
	// 2.备份一下 rename xx.log --> xx.log.bak202204051114
	os.Rename(logName, newLogName)
	// 3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open new log file failed,err:%v", err)
		return nil, err
	}
	// 4.将打开的新的日志文件赋值给 f.fileObj
	return fileObj, err
}

func (f FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)  //获取文件名字和文件行号
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj) //用来切割文件大小。
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			//如果要记录的日志大于等于ERROR级别，我还要在err日志文件中在记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed ,err:%v\n", err)
		return false
	}
	//如果当前文件大小 大于等于 日志文件的最大值 就应该返回true 表示切割
	return fileInfo.Size() > f.maxFileSize
}

func (f FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

func (f FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

func (f FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
