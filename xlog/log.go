package xlog

import (
	"fmt"
)

type XLog interface { //定义一个日志接口
	// 需求1：打印各个level日志
	LogDebug(fmt string, args ...interface{}) //格式化输出，并且是带多个可变参数，因为我们也不清楚要输入多少内容
	LogTrace(fmt string, args ...interface{})
	LogInfo(fmt string, args ...interface{})
	LogWarn(fmt string, args ...interface{})
	LogError(fmt string, args ...interface{})
	LogFatal(fmt string, args ...interface{})

	// 需求2：设置日志级别
	SetLevel(level int)
}

// 需求3：构造函数（go中是没有构造函数的，如果要生成一个日志对象，需要自己写一个构造函数）
func NewXLog(LogType, level int, filename, module string) XLog { //传入日志输出类型、日志级别、文件路径，模块名，返回给接口

	var logger XLog //定义一个接口类型的变量
	switch LogType {
	case XLogTypeFile: //输出到文件
		logger = NewXFile(level, filename, module)

	case XLogTypeConsole: //输出到控制台
		logger = NewXConsole(level, module)
	default:
		logger = NewXFile(level, filename, module)
	}
	return logger //返回接口实例（用户拿到的是一个接口类型的变量）
}

func Wea() {
	fmt.Println("a")
}
