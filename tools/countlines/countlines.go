package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

const (
	GoFile = ".go"
	MdFile = ".md"
)

func main() {
	result := make(map[string]int)

	countCache(result)
	countLib(result)
	countDocs(result)
	countTools(result)

	var total int
	for key, value := range result {
		fmt.Printf("committer: %s, total: %d\n", key, value)
		total += value
	}
	fmt.Printf("total code line: %d\n", total)
}

func countCache(countResult map[string]int) map[string]int {
	return countMod("../../cache", GoFile, countResult)
}

func countLib(countResult map[string]int) map[string]int {
	return countMod("../../lib", GoFile, countResult)
}

func countDocs(countResult map[string]int) map[string]int {
	return countMod("../../docs", MdFile, countResult)
}

func countTools(countResult map[string]int) map[string]int {
	return countMod("../", GoFile, countResult)
}

func countMod(modPath string, fileType string, countResult map[string]int) map[string]int {
	res := scanDirectory(modPath)

	for _, r := range res {
		countSingleFile(r, fileType, countResult)
	}

	return countResult
}

func scanDirectory(dir string) []string {
	fileInfos, err := ioutil.ReadDir(dir)
	var dirList []string

	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInfos {
		nextFileInfo := path.Join(dir, fileInfo.Name())
		if fileInfo.IsDir() {
			dirList = append(dirList, scanDirectory(nextFileInfo)...)
		} else {
			dirList = append(dirList, nextFileInfo)
		}
	}

	return dirList
}

func countSingleFile(file string, fileType string, result map[string]int) map[string]int {
	if !strings.Contains(file, fileType) {
		return result
	}

	cmd := exec.Command("git", "blame", file)
	output, _ := cmd.CombinedOutput()

	for _, line := range strings.Split(string(output), "\n") {
		reg := regexp.MustCompile(`\((\w+)\s`)
		committer := reg.FindStringSubmatch(line)

		if len(committer) > 1 {
			if _, ok := result[committer[1]]; ok {
				result[committer[1]] += 1
			} else {
				result[committer[1]] = 1
			}
		}
	}

	return result
}
