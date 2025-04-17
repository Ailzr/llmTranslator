package logHelper

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

const (
	dateLayout = "2006-01-02"
	timeLayout = "2006-01-02 15:04:05"
	redPre     = "31"
	greenPre   = "32"
	yellowPre  = "33"
	bluePre    = "34"
	blackBack  = "40"
)

func init() {
	//创建logs文件夹
	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		os.Mkdir("logs", 0644)
	}
}

func cmdOutput(outType, log, preground, background string, a ...any) {
	currentTime := time.Now().Format(timeLayout)
	log = fmt.Sprintf(log, a...)
	fmt.Println(fmt.Sprintf("\033[1;%s;%sm %s | %s | %s\033[0m", preground, background, outType, currentTime, log))
}

func Info(log string, a ...any) {
	cmdOutput("INFO", log, greenPre, blackBack, a...)
}

func Debug(log string, a ...any) {
	cmdOutput("DEBUG", log, bluePre, blackBack, a...)
}

func Warn(log string, a ...any) {
	cmdOutput("WARN", log, yellowPre, blackBack, a...)
}

func Error(log string, a ...any) {
	cmdOutput("ERROR", log, redPre, blackBack, a...)
	WriteLog(log, a...)
}

func WriteLog(log string, a ...any) {

	log = fmt.Sprintf(log, a...)

	// 获取报错位置
	_, callFile, line, _ := runtime.Caller(1)

	// 获取当前日期
	currentDate := time.Now().Format(dateLayout)

	// 创建logFile
	file, err := os.OpenFile(fmt.Sprintf("logs/%s logfile.log", currentDate), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 获取当前时间
	currentTime := time.Now().Format(timeLayout)
	_, err = fmt.Fprintf(file, "ERROR | %s | %s:%d | error info: %s\n", currentTime, callFile, line, log)
}
