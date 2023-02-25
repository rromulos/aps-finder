package helpers

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/rromulos/aps-finder/helpers/logger"
)

func PerformAnalysis(targetFolder, ext string, prefix string) {
	findAllFilesByExtension(targetFolder, ext)
}

func findAllFilesByExtension(targetFolder, ext string) []string {
	var a []string
	filepath.WalkDir(targetFolder, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(d.Name()) == ext {
			if !strings.Contains(s, "app/protected/vendor/") {
				logger.Log("INFO", "Analyzing file "+s, "execution")
				searchForAppSettingInFile(s)
				a = append(a, s)
			}
		}
		return nil
	})
	return a
}

func searchForAppSettingInFile(file string) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Log("ERROR", "Can't open the file "+file, "execution")
	}
	s := string(b)

	count := strings.Count(s, "AppSettingManager::get")
	fmt.Println("count ", count)

	for strings.Index(s, "AppSettingManager::get") != -1 {
		idxFind := strings.Index(s, "AppSettingManager::get")
		left := strings.LastIndex(s[:idxFind], "\n")
		if left == -1 {
			left = 1
		}
		right := strings.Index(s[idxFind:], "\n")
		fmt.Println("file ", file)
		occurrence := s[left : idxFind+right]
		fmt.Println(s[left : idxFind+right])
		s = strings.Replace(s, occurrence, "", -1)
	}
}
