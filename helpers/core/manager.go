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

func PerformAnalysis(targetFolder, ext string, pVerboseMode string) {
	verboseMode = pVerboseMode
	findAllFilesByExtension(targetFolder, ext)
}

func findAllFilesByExtension(targetFolder, ext string) []string {
	var count = 0
	qtySuccess := 0
	qtyWarning := 0
	qtyError := 0
	fQtySuccess := 0
	fQtyWarning := 0
	fQtyError := 0
	var a []string
	filepath.WalkDir(targetFolder, func(s string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if filepath.Ext(d.Name()) == ext {
			if !strings.Contains(s, "app/protected/vendor/") {
				count++

				if verboseMode == "y" {
					println("Analyzing file " + s)
				}

				logger.Log(logger.INFO, "Analyzing file "+s, logger.EXECUTION_FILE_NAME)
				qtySuccess, qtyWarning, qtyError = searchForAppSettingInFile(s)
				fQtySuccess += qtySuccess
				fQtyWarning += qtyWarning
				fQtyError += qtyError
				a = append(a, s)
			}
		}

		return nil
	})

	if fQtyWarning > 0 || fQtyError > 0 {
		println("=====================================================================")
		println("Finished, but needs attention. [", count, "] files were analyzed.")
		println(strconv.Itoa(fQtyWarning) + " WARNING(s) found during the analysis")
		println(strconv.Itoa(fQtyError) + " ERROR(s) found during the analysis")
		println(strconv.Itoa(fQtySuccess) + " App Settings successfully found")
	} else {
		println("=====================================================================")
		println("Successfully Finished. [", count, "] files were analyzed.")
		println("App_settings successfully found = ", fQtySuccess)
	}

	return a
}

func searchForAppSettingInFile(file string) (int, int, int) {
	qtyError := 0
	qtyWarning := 0
	qtySuccess := 0
	b, err := ioutil.ReadFile(file)

	if err != nil {

		if verboseMode == "y" {
			println("Can't open the file " + file)
		}

		logger.Log(logger.ERROR, "Can't open the file "+file, logger.EXECUTION_FILE_NAME)
		qtyError++
	}

	s := string(b)

	count := strings.Count(s, APP_SETTING_PATTERN)
	logger.Log(logger.INFO, "Number of occurrences "+strconv.Itoa(count), logger.EXECUTION_FILE_NAME)

	if verboseMode == "y" {
		println("Number of occurrences ", count)
	}

	for strings.Index(s, APP_SETTING_PATTERN) != -1 {
		idxFind := strings.Index(s, APP_SETTING_PATTERN)
		left := strings.LastIndex(s[:idxFind], "\n")

		if left == -1 {

			if verboseMode == "y" {
				println("Couldn't get LastIndex, setting 1 ")
			}

			logger.Log(logger.WARN, "Couldn't get LastIndex, setting 1 ", logger.EXECUTION_FILE_NAME)
			left = 1
			qtyWarning++
		}

		right := strings.Index(s[idxFind:], "\n")
		occurrence := s[left : idxFind+right]
		r, _ := regexp.Compile(`\([^()]*\)`)
		cleanedString := removeFromPattern(APP_SETTING_PATTERN, s[left:idxFind+right])

		for _, match := range r.FindStringSubmatch(cleanedString) {
			appSetting := removeUnnecessaryChars(match)

			if len(match) == 0 {

				if verboseMode == "y" {
					println("Empty value found")
				}

				qtyWarning++
				logger.Log(logger.WARN, "Empty value found", logger.EXECUTION_FILE_NAME)
				continue
			}

			checkReturn := checkContentContainsInvalidChars(appSetting)

			if !checkReturn {
				if !report.CheckAppSettingAlreadyExists(appSetting) {

					if verboseMode == "y" {
						println("[" + match + "] contains values ​​in variables that cannot be read")
					}

					report.AddToOutputReport(report.OUTPUT_WARNING_FILE_NAME, appSetting)
					logger.Log(logger.WARN, "["+match+"] contains values ​​in variables that cannot be read", "execution")
					qtyWarning++
				}
			} else {
				if !report.CheckAppSettingAlreadyExists(appSetting) {
					report.AddToOutputReport(report.OUTPUT_SUCCESS_FILE_NAME, appSetting)
					qtySuccess++
				}

			}
		}

		s = strings.Replace(s, occurrence, "", -1)
	}
	return qtySuccess, qtyWarning, qtyError
}

func removeFromPattern(p, ms string) string {
	_, a, ok := strings.Cut(ms, p)

	if !ok {
		return ""
	}

	return a
}

func removeUnnecessaryChars(s string) string {
	s1 := strings.Replace(s, "(", "", -1)
	s2 := strings.Replace(s1, ")", "", -1)
	s3 := strings.Replace(s2, "'", "", -1)
	s4 := strings.Trim(s3, "\"")

	return s4
}

func checkContentContainsInvalidChars(s string) bool {
	r, _ := regexp.Compile(`^[a-zA-Z0-9_/s/.]+[/s]*$`)

	if r.MatchString(s) {
		return true
	}

	return false
}
