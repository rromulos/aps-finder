package helpers

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
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
			println("buscando no arquivo = ", s)
			searchForAppSettingInFile(s)
			a = append(a, s)
		}
		return nil
	})
	return a
}

func searchForAppSettingInFile(file string) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	s := string(b)

	count := strings.Count(s, "AppSettingManager::get")
	fmt.Println("count ", count)

	for strings.Index(s, "AppSettingManager::get") != -1 {
		idxFind := strings.Index(s, "AppSettingManager::get")
		left := strings.LastIndex(s[:idxFind], "\n")
		right := strings.Index(s[idxFind:], "\n")
		fmt.Println("left ", left)
		fmt.Println("right ", right)
		occurrence := s[left : idxFind+right]
		fmt.Println(s[left : idxFind+right])
		s = strings.Replace(s, occurrence, "", -1)
	}
}
