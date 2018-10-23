package xlog

import (
	"runtime"
)

func getLineInfo(skip int) (fileName, funcName string, lineNo int) { //要传递skip（调用深度进去）
	pc, file, line, ok := runtime.Caller(skip) //获取当前调用栈的一些信息（返回:1、pc：程序计数器（获取到当前执行指令所在行的名字）2、文件名字 3、行号 4、ok（是否调用成功））
	if ok {                                    //如果调用成功，将pc换成函数名字
		fun := runtime.FuncForPC(pc)
		funcName = fun.Name()
	}

	fileName = file
	lineNo = line
	return
} //返回文件名、函数名、行号
