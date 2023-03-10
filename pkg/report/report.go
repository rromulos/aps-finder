package report

import (
	"fmt"
	"os"
	"strings"

	"github.com/rromulos/aps-finder/pkg/logger"
)

const OUTPUT_SUCCESS_FILE_NAME string = "output"
const OUTPUT_WARNING_FILE_NAME string = "output_warning"
const OUTPUT_FILE_NAME_EXTENSION string = ".txt"

//Adds the app_setting informed in the output report.
//The filename defines whether the app_setting will be inserted in the warning or success report.
//The OUTPUT_SUCCESS_FILE_NAME constant represents the success report.
//The OUTPUT_WARNING_FILE_NAME constant represents the report containing warnings.
//@TODO add logging for error situations (checks where err is not null).
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

//Remove all report files.
//This method is invoked every time the option referring to the search for app_settings gets triggered.
func destroyReportFiles() {
	os.Remove("output/" + OUTPUT_SUCCESS_FILE_NAME + OUTPUT_FILE_NAME_EXTENSION)
	os.Remove("output/" + OUTPUT_WARNING_FILE_NAME + OUTPUT_FILE_NAME_EXTENSION)
}

//Checks if the given app_setting already exists in the success report.
//Returns true if so.
//Returns false if not.
//@TODO add logging for error situations (checks where err is not null).
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

//Creates all blank report files.
func InitReports() {
	destroyReportFiles()
	_, err := os.Create("output/" + OUTPUT_SUCCESS_FILE_NAME + OUTPUT_FILE_NAME_EXTENSION)

	if err != nil {
		logger.Log(logger.ERROR, "Can't create "+OUTPUT_SUCCESS_FILE_NAME+OUTPUT_FILE_NAME_EXTENSION, logger.APP_FINDER_LOG)
	}

	_, err = os.Create("output/" + OUTPUT_WARNING_FILE_NAME + OUTPUT_FILE_NAME_EXTENSION)

	if err != nil {
		logger.Log(logger.ERROR, "Can't create "+OUTPUT_WARNING_FILE_NAME+OUTPUT_FILE_NAME_EXTENSION, logger.APP_FINDER_LOG)
	}
}
