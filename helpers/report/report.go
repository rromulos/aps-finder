package report

import (
	"fmt"
	"os"
)

const OUTPUT_SUCCESS_FILE_NAME string = "output"
const OUTPUT_WARNING_FILE_NAME string = "output_warning"
const OUTPUT_ERROR_FILE_NAME string = "output_error"
const OUTPUT_FILE_NAME_EXTENSION string = ".txt"

func AddToOutputReport(filename string, appSetting string) {
	f, err := os.OpenFile("output/"+filename+OUTPUT_FILE_NAME_EXTENSION, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.WriteString(appSetting + "\n")
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

func AddToOutputReportDetailed() {

}

func destroyReportFiles() {
	os.Remove("output/execution.log")
	os.Remove("output/output.log")
}

func InitLogs() {
	destroyReportFiles()
}
