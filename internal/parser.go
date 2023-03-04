package core

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/rromulos/aps-finder/pkg/logger"
	"github.com/rromulos/aps-finder/pkg/report"
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
	filepath.WalkDir(targetFolder, func(filePath string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if filepath.Ext(d.Name()) == ext {

			//excludes the vendor folder
			if !strings.Contains(filePath, "app/protected/vendor/") {

				fileExtension := filepath.Ext(filePath)

				//ignore all files that do not have a php extension
				if fileExtension == ".php" {
					count++

					if verboseMode == "y" {
						println(ANALYZING_FILE + filePath)
					}

					logger.Log(logger.INFO, ANALYZING_FILE+filePath, logger.EXECUTION_FILE_NAME)
					qtySuccess, qtyWarning, qtyError = searchForAppSettingInFile(filePath)
					a = append(a, filePath)
				}

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
	contentBytes, err := ioutil.ReadFile(file)

	if err != nil {

		if verboseMode == "y" {
			println(CANT_OPEN_FILE + file)
		}

		logger.Log(logger.ERROR, CANT_OPEN_FILE+file, logger.EXECUTION_FILE_NAME)
		qtyError++
	}

	content := string(contentBytes)
	count := strings.Count(content, APP_SETTING_PATTERN)
	logger.Log(logger.INFO, NUMBER_OF_OCCURRENCES+strconv.Itoa(count), logger.EXECUTION_FILE_NAME)

	if verboseMode == "y" {
		println(NUMBER_OF_OCCURRENCES, count)
	}

	for strings.Index(content, APP_SETTING_PATTERN) != -1 {

		idxFind := strings.Index(content, APP_SETTING_PATTERN)
		left := strings.LastIndex(content[:idxFind], "\n")

		if left == -1 {

			if verboseMode == "y" {
				println(COULD_NOT_GET_LAST_INDEX)
			}

			logger.Log(logger.WARN, COULD_NOT_GET_LAST_INDEX, logger.EXECUTION_FILE_NAME)
			left = 1
			qtyWarning++
		}

		right := strings.Index(content[idxFind:], "\n")
		occurrence := content[left : idxFind+right]
		//considers only the value after the get method inside parentheses and single quotes
		result, _ := regexp.Compile(`\([^)]+\)|get\(''\)`)
		cleanedString := removeByPattern(APP_SETTING_PATTERN, content[left:idxFind+right])

		for _, match := range result.FindStringSubmatch(cleanedString) {

			if len(match) == 0 {

				if verboseMode == "y" {
					println(EMPTY_VALUE)
				}

				qtyWarning++
				logger.Log(logger.WARN, EMPTY_VALUE, logger.EXECUTION_FILE_NAME)
				continue
			}

			appSetting := getValueBetweenSingleQuotes(match)
			containsInvalidChars := checkAppSettingIsValid(appSetting)

			if !report.CheckAppSettingAlreadyExists(appSetting) {

				if !containsInvalidChars {

					if verboseMode == "y" {
						println("["+match+"] ", CANT_READ_VALUE_FROM_PHP_VARIABLE)
					}

					report.AddToOutputReport(report.OUTPUT_WARNING_FILE_NAME, match)
					logger.Log(logger.WARN, "["+match+"] "+CANT_READ_VALUE_FROM_PHP_VARIABLE, logger.EXECUTION_FILE_NAME)
					qtyWarning++

				} else {
					report.AddToOutputReport(report.OUTPUT_SUCCESS_FILE_NAME, appSetting)
					qtySuccess++
				}
			}
		}

		content = strings.Replace(content, occurrence, "", -1)
	}
	return qtySuccess, qtyWarning, qtyError
}

//Removes content before the given string
//Return string
func removeByPattern(pattern, appSetting string) string {
	_, appSettingCleaned, ok := strings.Cut(appSetting, pattern)

	if !ok {
		return ""
	}

	return appSettingCleaned
}

//Check if the found app_setting is valid
//accepts all numbers, letters hyphen, underscore and dot
//but the string can't start with hyphen underscore and dot
func checkAppSettingIsValid(appSetting string) bool {
	result, _ := regexp.Compile(`^[a-zA-Z0-9][a-zA-Z0-9._-]*$`)
	return result.MatchString(appSetting)
}

//Get value between single quotes
//This method also replaces double quotes with single quotes
func getValueBetweenSingleQuotes(appSetting string) string {
	//using replace because go has problem with regex that uses backreference (\1)
	appSetting = strings.Replace(appSetting, "\"", "'", -1)
	re := regexp.MustCompile(`'([^']*)'`)
	matches := re.FindAllStringSubmatch(appSetting, -1)

	if len(matches) > 0 {

		for _, match := range matches {
			return match[1]
		}

	}

	return " "
}
