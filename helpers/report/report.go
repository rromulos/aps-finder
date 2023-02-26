package report

import (
	"fmt"
	"os"
	"strings"

	"github.com/rromulos/aps-finder/helpers/logger"
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
	os.Remove("output/" + OUTPUT_SUCCESS_FILE_NAME + OUTPUT_FILE_NAME_EXTENSION)
	os.Remove("output/" + OUTPUT_WARNING_FILE_NAME + OUTPUT_FILE_NAME_EXTENSION)
}

func CheckAppSettingAlreadyExists(appSetting string) bool {
	data, err := os.ReadFile("output/" + OUTPUT_SUCCESS_FILE_NAME + OUTPUT_FILE_NAME_EXTENSION)

	if err != nil {
		panic(err)
	}

	s := string(data)

	if strings.Contains(s, appSetting) {
		return true
	} else {
		return false
	}
}

func InitReports() {
	destroyReportFiles()
	_, err := os.Create("output/" + OUTPUT_SUCCESS_FILE_NAME + OUTPUT_FILE_NAME_EXTENSION)
	if err != nil {
		logger.Log(logger.ERROR, "Can't create "+OUTPUT_SUCCESS_FILE_NAME+OUTPUT_FILE_NAME_EXTENSION, logger.SYSTEM_FILE_NAME)
	}
	_, err2 := os.Create("output/" + OUTPUT_WARNING_FILE_NAME + OUTPUT_FILE_NAME_EXTENSION)
	if err2 != nil {
		logger.Log(logger.ERROR, "Can't create "+OUTPUT_WARNING_FILE_NAME+OUTPUT_FILE_NAME_EXTENSION, logger.SYSTEM_FILE_NAME)
	}

}
