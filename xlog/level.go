package xlog

const (
	XLogLevelDebug = iota
	XLogLevelTrace
	XLogLevelInfo
	XLogLevelWarn
	XLogLevelError
	XLogLevelFatal
)

const (
	XLogTypeFile    = iota //输出到文件
	XLogTypeConsole        //输出到控制台
)

//该函数为将常量转化为可读的文本或字符串(私有（函数名小写开头表示私有，仅可以内部调用）)
func getLevelStr(level int) string {
	switch level {
	case XLogLevelDebug:
		return "DEBUG"
	case XLogLevelTrace:
		return "TRACE"
	case XLogLevelInfo:
		return "INFO"
	case XLogLevelWarn:
		return "WARN"
	case XLogLevelError:
		return "ERROR"
	case XLogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"

	}
}
