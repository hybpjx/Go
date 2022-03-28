package mylog

import (
	"fmt"
	"os"
	"time"
)

// 往终端上打印日志

// ConsoleLogger是一个终端日志结构体
type ConsoleLogger struct {
	level Level

	//file *os.File // 直接打印就不用传一个文件了
}

// NewConsoleLogger 文件日志结构体的构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger {

	logLevel := parseLevelStr(levelStr)

	c1 := &ConsoleLogger{
		level: logLevel,
		//file: os.Stdout, // linux 下万物皆文件 所以可以改成os.Studout 这样后续就不用改了
	}
	return c1
}

// currencyLog (通用日志) 将共用的记录日志的方法 封装成一个函数
func (c *ConsoleLogger) currencyLog(loglevel Level, format string, args ...interface{}) {
	// 传入的log 等级
	if c.level > loglevel {
		return
	}

	LevelStr := getLevelStr(loglevel)

	msg := fmt.Sprintf(format, args...)
	// 利用fmt 包 将msg 写入文件中

	// 日志格式[2022-01][文件格式：行号] [函数名] [日志级别] 日志信息

	nowStr := time.Now().Format("2006-01-02 15:04:05.000")

	fileName, line, funcName := getCallerInfo(3)

	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", nowStr, fileName, line, LevelStr, funcName, msg)

	//fmt.Fprintln(c.file, logMsg) // 利用fmt的fprintf将logmsg 写入 文件中  哪怕是console

	fmt.Fprintln(os.Stdout, logMsg)

}

// Debug 方法
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.currencyLog(DebugLevel, format, args...)
}

// 	Trace 方法
func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	c.currencyLog(TraceLevel, format, args...)
}

// Info 方法
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.currencyLog(InfoLevel, format, args...)
}

// Warn 方法
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.currencyLog(WarningLevel, format, args...)

}

// Error 方法
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.currencyLog(ErrorLevel, format, args...)

}

// Fatal 方法
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.currencyLog(FatalLevel, format, args...)

}

// Close 方法
func (c *ConsoleLogger) Close() {
}
