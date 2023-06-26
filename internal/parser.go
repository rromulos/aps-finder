package core

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rromulos/aps-finder/internal/messages"
	"github.com/rromulos/aps-finder/pkg/logger"
	"github.com/rromulos/aps-finder/pkg/report"
)

const appSettingPattern = "AppSettingManager::get"

var (
	verboseMode string
	qtySuccess  int
	qtyWarning  int
	qtyError    int
)

// PerformAnalysis starts parsing for app_settings.
// pVerboseMode represents whether verbose mode should be considered or not.
func PerformAnalysis(pVerboseMode string) {
	godotenv.Load()
	verboseMode = pVerboseMode
	start := time.Now()
	findAllFilesByExtension(os.Getenv("APP_PATH"), ".php")
	finished := time.Since(start)
	println("=====================================================================")
	log.Printf("Execution took %s", finished)
}

// findAllFilesByExtension searches for all files with the PHP extension.
func findAllFilesByExtension(targetFolder, ext string) []string {
	var count int

	qtySuccess = 0
	qtyWarning = 0
	qtyError = 0

	var a []string
	filepath.WalkDir(targetFolder, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(d.Name()) == ext && !strings.Contains(filePath, "app/protected/vendor/") {
			count++

			if verboseMode == "y" {
				println(messages.ANALYZING_FILE + filePath)
			}

			logger.Log(logger.INFO, messages.ANALYZING_FILE+filePath, logger.EXECUTION_FILE_NAME)
			qtySuccess, qtyWarning, qtyError = searchForAppSettingInFile(filePath)
			a = append(a, filePath)
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

// searchForAppSettingInFile parses the given file looking for AppSettings.
func searchForAppSettingInFile(file string) (int, int, int) {
	contentBytes, err := ioutil.ReadFile(file)
	if err != nil {
		if verboseMode == "y" {
			println(messages.CANT_OPEN_FILE + file)
		}
		logger.Log(logger.ERROR, messages.CANT_OPEN_FILE+file, logger.EXECUTION_FILE_NAME)
		qtyError++
	}

	content := string(contentBytes)
	count := strings.Count(content, appSettingPattern)
	logger.Log(logger.INFO, messages.NUMBER_OF_OCCURRENCES+strconv.Itoa(count), logger.EXECUTION_FILE_NAME)

	if verboseMode == "y" {
		println(messages.NUMBER_OF_OCCURRENCES, count)
	}

	for strings.Index(content, appSettingPattern) != -1 {
		idxFind := strings.Index(content, appSettingPattern)
		left := strings.LastIndex(content[:idxFind], "\n")

		if left == -1 {
			if verboseMode == "y" {
				println(messages.COULD_NOT_GET_LAST_INDEX)
			}
			logger.Log(logger.WARN, messages.COULD_NOT_GET_LAST_INDEX, logger.EXECUTION_FILE_NAME)
			left = 1
			qtyWarning++
		}

		right := strings.Index(content[idxFind:], "\n")
		occurrence := content[left : idxFind+right]
		result, _ := regexp.Compile(`\([^)]+\)|get\(''\)`)
		cleanedString := removeByPattern(appSettingPattern, content[left:idxFind+right])

		for _, match := range result.FindStringSubmatch(cleanedString) {
			if len(match) == 0 {
				if verboseMode == "y" {
					println(messages.EMPTY_VALUE)
				}
				qtyWarning++
				logger.Log(logger.WARN, messages.EMPTY_VALUE, logger.EXECUTION_FILE_NAME)
				continue
			}

			appSetting := getValueBetweenSingleQuotes(match)
			containsInvalidChars := checkAppSettingIsValid(appSetting)

			if !report.CheckAppSettingAlreadyExists(appSetting) {
				if !containsInvalidChars {
					if verboseMode == "y" {
						println("["+match+"] ", messages.CANT_READ_VALUE_FROM_PHP_VARIABLE)
					}
					report.AddToOutputReport(report.OUTPUT_WARNING_FILE_NAME, match)
					logger.Log(logger.WARN, "["+match+"] "+messages.CANT_READ_VALUE_FROM_PHP_VARIABLE, logger.EXECUTION_FILE_NAME)
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

// removeByPattern removes content before the given string.
func removeByPattern(pattern, appSetting string) string {
	_, appSettingCleaned, ok := strings.Cut(appSetting, pattern)
	if !ok {
		return ""
	}
	return appSettingCleaned
}

// checkAppSettingIsValid checks if the given app_setting is valid.
func checkAppSettingIsValid(appSetting string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9._-]*$`)
	return regex.MatchString(appSetting)
}

// getValueBetweenSingleQuotes gets value between single quotes.
func getValueBetweenSingleQuotes(appSetting string) string {
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
