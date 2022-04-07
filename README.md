日志库需求分析：

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
