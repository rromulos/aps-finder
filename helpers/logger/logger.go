package logger

import (
	"fmt"
	"os"
	"time"
)

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
	l, err := f.WriteString(logContent + "\n")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "log written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
