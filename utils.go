package goutils

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"time"
)

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func CheckError(e error) {
	if e != nil {
		log.Fatalf("ERROR: %#v", e)
	}
}

func GetRegexSubgroups(exp *regexp.Regexp, text string) map[string]string {
	match := exp.FindStringSubmatch(text)
	resultMap := make(map[string]string)
	for i, name := range exp.SubexpNames() {
		if i != 0 && name != "" {
			resultMap[name] = match[i]
		}
	}
	return resultMap
}

func StringsContain(strings []string, match string) bool {
	for _, ele := range strings {
		if ele == match {
			return true
		}
	}
	return false
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ReadFileToLines(file string) []string {
	openFile, err := os.Open(file)
	defer CloseFile(openFile)
	CheckError(err)

	var lines []string
	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func CloseFile(file *os.File) {
	err := file.Close()
	CheckError(err)
}
