package goutils

import (
	"bufio"
	"github.com/juju/errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ConfigProperties map[string]string

func (p *ConfigProperties) HasProperty(propertyName string) bool {
	_, exists := (*p)[propertyName]
	return exists
}

func (p *ConfigProperties) GetPropertyAsInt(propertyName string) int {
	val := (*p)[propertyName]
	i, err := strconv.ParseInt(val, 0, 32)
	CheckError(err)
	return int(i)
}

func (p *ConfigProperties) GetPropertyAsString(propertyName string) string {
	return (*p)[propertyName]
}

func ReadPropertiesFile(filename string) *ConfigProperties {
	file, err := os.Open(filename)
	CheckError(err)
	defer CloseFile(file)

	config := ConfigProperties{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal, comment := strings.Index(line, "="), strings.Index(line, "#"); equal >= 0 && comment == -1 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	CheckError(scanner.Err())
	return &config
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func CheckError(e error) {
	if e != nil {
		log.Fatalf("ERROR: %#v", errors.Trace(e))
	}
}

func CheckErrorWithMessage(e error, message string) {
	if e != nil {
		log.Fatalf("ERROR: %s\n%#v", errors.Trace(e))
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
	CheckError(err)
	defer CloseFile(openFile)

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

func LogError(message string, err error) {
	log.Printf("ERROR: %s: %#v\n", message, errors.Trace(err))
}

func LogInfo(message string) {
	log.Printf("INFO: %s\n", message)
}
