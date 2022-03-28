package mylog

import "strings"

type Level uint16

const (
	DebugLevel Level = iota
	TraceLevel
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

// 得到返回日志级别的字符串函数
func getLevelStr(level Level) string {
	// 根据输入的值返回相应字符串
	switch level {
	case DebugLevel:
		return "Debug"
	case TraceLevel:
		return "Trace"
	case InfoLevel:
		return "Info"
	case WarningLevel:
		return "Warning"
	case ErrorLevel:
		return "Error"
	case FatalLevel:
		return "Fatal"
	default:
		return "Debug"
	}
}

// 根据用户 传入的字符串 判断日志的级别
func parseLevelStr(levelStr string) Level {
	// 将字符串全大写
	levelStr = strings.ToUpper(levelStr)
	switch levelStr {
	case "DEBUG":
		return DebugLevel
	case "TRACE":
		return TraceLevel
	case "INFO":
		return InfoLevel
	case "WARNING":
		return WarningLevel
	case "ERROR":
		return ErrorLevel
	case "FATAL":
		return FatalLevel
	default:
		return DebugLevel
	}
}
