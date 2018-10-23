package xlog

import (
	"fmt"
)

type XFile struct {
	level    int    //日志级别
	filename string //文件名
	module   string //模块
}

func NewXFile(level int, filename, module string) XLog {
	logger := &XFile{ //生成（new）一个文件的实例，因为文件实例已经实现了XLog接口的规范，所以接口变量logger就可以存储这个文件实例
		level:    level,
		filename: filename,
		module:   module,
	}
	return logger
}

func (c *XFile) LogDebug(format string, args ...interface{}) {
	fmt.Printf("log debug of file\n")
}

func (c *XFile) LogTrace(format string, args ...interface{}) {
	fmt.Printf("log trace of file\n")
}

func (c *XFile) LogInfo(format string, args ...interface{}) {
	fmt.Printf("log info of file\n")
}

func (c *XFile) LogWarn(format string, args ...interface{}) {
	fmt.Printf("log warn of file\n")
}

func (c *XFile) LogError(format string, args ...interface{}) {
	fmt.Printf("log error of file\n")
}

func (c *XFile) LogFatal(format string, args ...interface{}) {
	fmt.Printf("log fatal of file\n")
}

func (c *XFile) SetLevel(level int) {
	c.level = level
}
