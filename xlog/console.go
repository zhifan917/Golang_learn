package xlog

import (
	"fmt"
	"path/filepath"
	"time"
)

type XConsole struct {
	level  int    //控制台并不需要文件需要的是日志级别
	module string //模块
}

func NewXConsole(level int, module string) XLog {
	logger := &XConsole{ //生成（new）一个文件的实例，因为控制台实例已经实现了XLog接口的规范，所以接口变量logger就可以存储这个控制台实例
		level:  level,
		module: module,
	}
	return logger
}

func (c *XConsole) LogDebug(format string, args ...interface{}) {

	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05.000") //日期
	levelStr := getLevelStr(XLogLevelDebug)          //日志级别
	module := c.module                               //模块名
	fileName, funcName, lineNo := getLineInfo(2)     //文件名、函数名、行号
	fileName = filepath.Base(fileName)               //只获取文件名而不是连带文件路径一起获取了
	data := fmt.Sprintf(format, args...)             //日志body格式化

	fmt.Printf("%s %s %s (%s:%s:%d) %s\n", timeStr, levelStr, module, fileName, funcName, lineNo, data)
}

func (c *XConsole) LogTrace(format string, args ...interface{}) {
	fmt.Printf("log trace of Console\n")
}

func (c *XConsole) LogInfo(format string, args ...interface{}) {
	fmt.Printf("log info of Console\n")
}

func (c *XConsole) LogWarn(format string, args ...interface{}) {
	fmt.Printf("log warn of Console\n")
}

func (c *XConsole) LogError(format string, args ...interface{}) {
	fmt.Printf("log error of Console\n")
}

func (c *XConsole) LogFatal(format string, args ...interface{}) {
	fmt.Printf("log fatal of Console\n")
}

func (c *XConsole) SetLevel(level int) {
	c.level = level
}
