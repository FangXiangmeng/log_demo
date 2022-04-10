日志库需求分析：
```
1.支持往不同的地方输出日志
2.日志分级别：
- Debug
- Trace
- Info
- Warning
- Error
- Fatal
3.日志要支持开关控制。比如说开发的时候什么日志级别都能输出，但是上线之后只有INFO级别往下的日志才能输出。
4.完整的日志记录要包含有时间、行号、文件名、日志级别、日志信息 
5.日志文件要切割
- 按照文件大小切割，每次记录日志之前都判断一下当前写的这个文件的大小。
- 按照日期切割
```



代码分析
```
日志库调用流程
Console版本日志采集
1.main.go首先会调用console中的NewLog这个构造函数。
2.构造函数内部调用parseLogLevel函数来将传进去的字符串进行大小写转换，并返回大写的日志级别和err
3.main.go中继续调用log.Debug这个方法。然后传入字符串。
4.Debug这个方法内部调用c.log这个方法。并传入日志级别，字符串和空接口。
5.log这个方法接受日志级别 字符串 可变长空接口。
6.log这个方法内部调用c.enable并将INFO传进去。
7.enable这个方法主要是用来做日志级别的开关。
8.enable这个方法内部传进来的LogLevel和c.Level进行对比。如果在NewLog这个方法 传进去的是INFO 然后调用log.Debug就不会打印日志。因为Debug=1 Info=3. 当 LogLevel >= c.Level大于等于c.Level才会打印日志。
9.如果LogLevel大于等于c.Level就会打印日志，并走到getInfo这个函数中。
10.getInfo这个函数主要是调用runtime.Caller这个方法。runtime主要是记录堆栈信息。例如函数调用。这次使用主要是取函数名。传skip int并返回funcName, fileName string, lineNo int
11.log这个方法内部还会调用getLogString这个函数打印日志级别

方法调用流程： - NewLog -- parseLogLevel -- log.Debug -- log -- enable --  getInfo -- 



file版本日志采集(带日志切割功能)
1.通过在main.go函数中，调用mylogger这个包里面的NewFileLogger这个构造函数。
2.NewFileLogger这个函数接受日志级别，文件路径，文件名字，和文件大小
2.在NewFileLogger这个函数内部中调用parseLogLevel这个函数。
3.parseLogLevel主要是将大小写进行转换。将传入的参数转换为小写。并retrun 日志级别。
4.NewFileLogger这个函数调用initFile这个方法。
5.initFile这个方法主要是用来创建文件。并给FileLogger这个结构体赋值。
6.在main.go中调用Debug这个方法。
7.Debug这个方法调用log这个方法并传入日志级别，日志，还有interface
8.log这个方法又去调用enable这个方法来判断日志级别。
9.如果日志级别LogLevel >= f.Level就开始写日志。
10.写之前会调用f.checkSize这个方法。来判断文件大小。
11.checekSize这个方法中会嗲用file.Stat这个方法来获取文件状态。
12.然后返回fileInfo.Size() > f.maxFileSize这个判断。返回值是bool值。如果大于就切割。小与就不走这个方法
13.再往下周会调用splitFile用来切割文件。
14.splitFile这个方法主要是将文件进行rename然后重新在创建个新的文件。并将fileObj返回。
15.然后log这个方法中在对f.fileObj这个名字进行重新赋值。
16.往f.fileObj这个文件中开始打入新的日志。
17.错误日志也是上面相同的方法进行判断。
18.到此就结束了。

方法调用流程: - NewFileLogger -- parseLogLevel -- initFile -- log.Debug -- log -- enable --checkSize -- splitFile
```