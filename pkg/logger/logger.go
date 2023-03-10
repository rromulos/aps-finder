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
const APP_FINDER_LOG = "appfinder"

//Remove all log files.
func destroyLogFiles() {
	os.Remove("logs/execution.log")
	os.Remove("logs/output.log")
	os.Remove("logs/output.log")
}

//@TODO check why the log files are not being created in this method.
func InitLogs() {
	destroyLogFiles()
}

//Adds the given content to the given log file.
//level represents the level of the entry, accepted values (INFO, DEBUG, WARN and ERROR).
//content represents the content that needs to be logged.
//targetFile represents the log file where the content should be logged, accept values (EXECUTION_FILE_NAME, APP_FINDER_LOG).
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
