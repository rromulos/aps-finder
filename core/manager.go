package core

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/rromulos/aps-finder/helpers/logger"
	"github.com/rromulos/aps-finder/helpers/report"
)

const APP_SETTING_PATTERN string = "AppSettingManager::get"

var verboseMode = ""
var qtyError = 0
var qtyWarning = 0
var qtySuccess = 0

//Method that will start the analysis
func PerformAnalysis(targetFolder, ext string, pVerboseMode string) {
	verboseMode = pVerboseMode
	findAllFilesByExtension(targetFolder, ext)
}

//Searches for all files with the PHP extension
func findAllFilesByExtension(targetFolder, ext string) []string {
	var count = 0
	qtySuccess := 0
	qtyWarning := 0
	qtyError := 0

	var a []string
	filepath.WalkDir(targetFolder, func(s string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if filepath.Ext(d.Name()) == ext {
			if !strings.Contains(s, "app/protected/vendor/") {
				count++

				if verboseMode == "y" {
					println(ANALYZING_FILE + s)
				}

				logger.Log(logger.INFO, ANALYZING_FILE+s, logger.EXECUTION_FILE_NAME)
				qtySuccess, qtyWarning, qtyError = searchForAppSettingInFile(s)
				a = append(a, s)
			}
		}

		return nil
	})

	if qtyWarning > 0 || qtyError > 0 {
		println("=====================================================================")
		println("Finished, but needs attention. [", count, "] files were analyzed.")
		println(strconv.Itoa(qtyWarning) + " WARNING(s) found during the analysis")
		println(strconv.Itoa(qtyError) + " ERROR(s) found during the analysis")
		println(strconv.Itoa(qtySuccess) + " App Settings successfully found")
	} else {
		println("=====================================================================")
		println("Successfully Finished. [", count, "] files were analyzed.")
		println("App_settings successfully found = ", qtySuccess)
	}

	return a
}

//Reads the given file looking for AppSettings
func searchForAppSettingInFile(file string) (int, int, int) {

	b, err := ioutil.ReadFile(file)

	if err != nil {

		if verboseMode == "y" {
			println(CANT_OPEN_FILE + file)
		}

		logger.Log(logger.ERROR, CANT_OPEN_FILE+file, logger.EXECUTION_FILE_NAME)
		qtyError++
	}

	s := string(b)

	count := strings.Count(s, APP_SETTING_PATTERN)
	logger.Log(logger.INFO, NUMBER_OF_OCCURRENCES+strconv.Itoa(count), logger.EXECUTION_FILE_NAME)

	if verboseMode == "y" {
		println(NUMBER_OF_OCCURRENCES, count)
	}

	for strings.Index(s, APP_SETTING_PATTERN) != -1 {
		idxFind := strings.Index(s, APP_SETTING_PATTERN)
		left := strings.LastIndex(s[:idxFind], "\n")

		if left == -1 {

			if verboseMode == "y" {
				println(COULD_NOT_GET_LAST_INDEX)
			}

			logger.Log(logger.WARN, COULD_NOT_GET_LAST_INDEX, logger.EXECUTION_FILE_NAME)
			left = 1
			qtyWarning++
		}

		right := strings.Index(s[idxFind:], "\n")
		occurrence := s[left : idxFind+right]
		r, _ := regexp.Compile(`\([^)]+\)|get\(''\)`)
		cleanedString := removeFromPattern(APP_SETTING_PATTERN, s[left:idxFind+right])

		for _, match := range r.FindStringSubmatch(cleanedString) {

			appSetting := getValueBetweenSingleQuotes(match)

			if len(match) == 0 {

				if verboseMode == "y" {
					println(EMPTY_VALUE)
				}

				qtyWarning++
				logger.Log(logger.WARN, EMPTY_VALUE, logger.EXECUTION_FILE_NAME)
				continue
			}

			containsInvalidChars := checkContentContainsInvalidChars(appSetting)
			addContentToOutputReport(containsInvalidChars, match, appSetting)
		}

		s = strings.Replace(s, occurrence, "", -1)
	}
	return qtySuccess, qtyWarning, qtyError
}

//Removes content before the given string
//Return string
func removeFromPattern(p, s string) string {
	_, a, ok := strings.Cut(s, p)

	if !ok {
		return ""
	}

	return a
}

//Check if the string contains invalid characters
func checkContentContainsInvalidChars(s string) bool {
	r, _ := regexp.Compile(`^[a-zA-Z0-9_/s/.]+[/s]*$`)

	if r.MatchString(s) {
		return true
	}

	return false
}

//Invokes the report in order to add the outputs
func addContentToOutputReport(containsInvalidChars bool, match string, appSetting string) {
	if !containsInvalidChars {
		if !report.CheckAppSettingAlreadyExists(appSetting) {

			if verboseMode == "y" {
				println("["+match+"] ", CANT_READ_VALUE_FROM_PHP_VARIABLE)
			}

			report.AddToOutputReport(report.OUTPUT_WARNING_FILE_NAME, appSetting)
			logger.Log(logger.WARN, "["+match+"] "+CANT_READ_VALUE_FROM_PHP_VARIABLE, logger.EXECUTION_FILE_NAME)
			qtyWarning++
			println("warning -> ", qtyWarning)
		}
	} else {
		if !report.CheckAppSettingAlreadyExists(appSetting) {
			report.AddToOutputReport(report.OUTPUT_SUCCESS_FILE_NAME, appSetting)
			println("success -> ", qtySuccess)
			qtySuccess++
		}

	}
}

func getValueBetweenSingleQuotes(s string) string {
	re := regexp.MustCompile(`'([^']*)'`)
	matches := re.FindAllStringSubmatch(s, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			return match[1]
		}
	}
	return " "
}
