package logger

import (
	"fmt"
	"os"
	"time"
)

const INFO string = "INFO"
const DEBUG string = "DEBUG"
const WARN string = "WARN"
const ERROR string = "ERROR"
const EXECUTION_FILE_NAME = "execution"
const SYSTEM_FILE_NAME = "sytem"

func destroyLogFiles() {
	os.Remove("logs/execution.log")
	os.Remove("logs/output.log")
}

func InitLogs() {
	destroyLogFiles()
}

func Log(level string, content string, targetFile string) {
	f, err := os.OpenFile("logs/"+targetFile+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	logContent := time.Now().Format("2006-01-02 15:04:05") + " | " + level + " | " + content
	_, err = f.WriteString(logContent + "\n")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
