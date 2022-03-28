package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 日志文件结构体
type FileLogger struct {
	level    Level // 日志等级
	filePath string
	fileName string
	// os 的文件
	file    *os.File
	errFile *os.File
	maxSize int64
	// 定义一个专门用来存放日志信息的通道
	logDataChan chan *logData
}

type logData struct {
	TimeStr  string
	FileName string
	LineNum  int
	LogLevel string
	FuncName string
	Message  string
}

// Logger 定义一个 log结构
type Logger interface {
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

// NewFileLogger 文件日志结构体的构造函数
func NewFileLogger(levelStr, filePath, fileName string) *FileLogger {

	logLevel := parseLevelStr(levelStr)

	fileObj := &FileLogger{
		level:    logLevel,
		filePath: filePath,
		fileName: fileName,
		maxSize:  10 * 1024 * 1024,
		// 初始化通道
		logDataChan: make(chan *logData, 50000),
	}
	// 根据上面的文件路径和文件名打开日志文件，把文件句柄赋值给结构体对应的字段
	fileObj.initFile()

	return fileObj
}

// 方法
func (f *FileLogger) initFile() {
	filePathName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(filePathName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(fmt.Errorf("打开文件%s失败  失败原因为:%s", filePathName, err))
	}
	f.file = fileObj

	errFilePathName := fmt.Sprintf("Error%s", filePathName)

	errFileObj, err := os.OpenFile(errFilePathName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(fmt.Errorf("打开文件%s失败  失败原因为:%s", errFilePathName, err))
	}
	f.errFile = errFileObj

	// 开启goroutine 写日志
	go f.writeLogBackgroud()
}

// 检查是否要拆分
func (f *FileLogger) checkIsSplit(file *os.File) bool {
	// 拿到文件下详细信息
	fileInfo, _ := file.Stat()
	//拿到文件的大小是int64的
	fileSize := fileInfo.Size()
	// 当传进来的日志大小超过maxsize 就返回true
	return fileSize >= f.maxSize

}

// 封装一个 切分日志文件的方法
func (f *FileLogger) splitLogFile(file *os.File) *os.File {
	// 拆分文件

	// 拿到当前的名字
	fileName := file.Name()

	backupName := fmt.Sprintf("%s_%v.bak", fileName, time.Now().Unix())

	// 1. 关闭当前文件
	file.Close()
	// 2. 备份当前文件
	os.Rename(fileName, backupName)
	// 3. 新建文件  名字就用当前的好了
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(fmt.Errorf("打开文件%s失败  失败原因为:%s", fileName, err))
	}
	return fileObj
}

// currencyLog (通用日志) 将共用的记录日志的方法 封装成一个函数
func (f *FileLogger) currencyLog(loglevel Level, format string, args ...interface{}) {
	// 传入的log 等级
	if f.level > loglevel {
		return
	}

	LevelStr := getLevelStr(loglevel)

	msg := fmt.Sprintf(format, args...)
	// 利用fmt 包 将msg 写入文件中

	// 日志格式[2022-01][文件格式：行号] [函数名] [日志级别] 日志信息

	nowStr := time.Now().Format("2006-01-02 15:04:05.000")

	fileName, line, funcName := getCallerInfo(3)

	// 将日志发送到通道中，以备写日记的goroutine去接受值
	// 1. 构造日志结构体

	logDataStruct := &logData{
		TimeStr:  nowStr,
		FileName: fileName,
		LineNum:  line,
		LogLevel: LevelStr,
		FuncName: funcName,
		Message:  msg,
	}

	// 不写日志了 会有什么问题 50000 放满了 怎么办	当通道满了, 会产生阻塞
	select {
	case f.logDataChan <- logDataStruct:
	default:
		<-f.logDataChan                // 丢弃最早的一条
		f.logDataChan <- logDataStruct // 把最新的发进去
	}

}

// 真正往文件里写日志的函数
func (f *FileLogger) writeLogBackgroud() {

	for ch1 := range f.logDataChan {

		logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", ch1.TimeStr, ch1.FileName, ch1.LineNum, ch1.LogLevel, ch1.FuncName, ch1.Message)

		// 往文件写之前下需要进行一个检查  如果 大于最大大小 则拆分
		if f.checkIsSplit(f.file) {
			f.file = f.splitLogFile(f.file)
		}

		level := parseLevelStr(ch1.LogLevel)
		// 利用fmt的fprintf将logmsg 写入
		fmt.Fprintln(f.file, logMsg)

		if level >= ErrorLevel {

			if f.checkIsSplit(f.file) {
				f.errFile = f.splitLogFile(f.errFile)
			}
			// 如果是error级别 或者以上的日志还要额外记录到errorFile中
			fmt.Fprintln(f.errFile, logMsg)
		}

	}

}

// Debug 方法
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.currencyLog(DebugLevel, format, args...)
}

// 	Trace 方法
func (f *FileLogger) Trace(format string, args ...interface{}) {
	f.currencyLog(TraceLevel, format, args...)
}

// Info 方法
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.currencyLog(InfoLevel, format, args...)
}

// Warn 方法
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.currencyLog(WarningLevel, format, args...)

}

// Error 方法
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.currencyLog(ErrorLevel, format, args...)

}

// Fatal 方法
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.currencyLog(FatalLevel, format, args...)

}

// Fatal 方法
func (f *FileLogger) Close() {
	f.file.Close()
	f.errFile.Close()

}
