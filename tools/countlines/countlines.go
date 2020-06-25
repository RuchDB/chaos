package main

import (
	"io/ioutil"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

const (
	ConfFile   = ".conf"
	DockerFile = "Dockerfile"
	GoFile     = ".go"
	JSonFile   = ".json"
	MakeFile   = "Makefile"
	MdFile     = ".md"
	TSFile     = ".ts"
	TSXFile    = ".tsx"
)

func countStats() map[string]int {
	result := make(map[string]int)

	countCache(result)
	countLib(result)
	countDocs(result)
	countTest(result)
	countTools(result)
	countRoot(result)

	//fmt.Printf("%#v\n", result)
	return result
}

func countCache(countResult map[string]int) map[string]int {
	return countMod("../../cache", countResult, GoFile)
}

func countLib(countResult map[string]int) map[string]int {
	return countMod("../../lib", countResult, GoFile)
}

func countDocs(countResult map[string]int) map[string]int {
	return countMod("../../docs", countResult, MdFile)
}

func countTest(countResult map[string]int) map[string]int {
	return countMod("../", countResult, GoFile, MdFile)
}

func countTools(countResult map[string]int) map[string]int {
	return countMod("../", countResult, GoFile, TSFile, TSXFile, JSonFile, ConfFile, DockerFile)
}

func countRoot(countResult map[string]int) map[string]int {
	return countMod("../..", countResult, MakeFile, MdFile)
}

func countMod(modPath string, countResult map[string]int, fileTypes ...string) map[string]int {
	res := scanDirectory(modPath)
	for _, r := range res {
		countSingleFile(r, countResult, fileTypes...)
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

func countSingleFile(file string, result map[string]int, fileTypes ...string) map[string]int {
	for _, fileType := range fileTypes {
		if strings.Contains(file, fileType) {
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
		}
	}

	return result
}
