package main

import (
	"flag"
	"xlog"
)

func logic(logger xlog.XLog) {
	logger.LogDebug("shshshshshshshshsh")
	logger.LogDebug("shshshshshshshshsh")
	logger.LogDebug("shshshshshshshshsh")
	logger.LogDebug("shshshshshshshshsh")
	logger.LogDebug("shshshshshshshshsh")
	logger.LogDebug("shshshshshshshshsh")
	logger.LogDebug("shshshshshshshshsh")
}

func main() {

	var logTypeStr string
	flag.StringVar(&logTypeStr, "type", "file", "please input logger type") //借助flag包，动态将日志类型传递进来
	flag.Parse()

	var logType int
	if logTypeStr == "file" {
		logType = xlog.XLogTypeFile
	} else {
		logType = xlog.XLogTypeConsole
	}

	logger := xlog.NewXLog(logType, xlog.XLogLevelDebug, "", "xlog_example")
	logic(logger)

}
