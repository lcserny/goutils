package goutils

import (
	"bytes"
	"encoding/gob"
	"os"
	"regexp"
	"strings"
	"time"
)

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
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

func GetLinesFromString(content string) []string {
	return strings.Split(content, "\n")
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

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GetRealPath(path string) (string, os.FileInfo) {
	simLinkInfo, err := os.Lstat(path)
	LogError(err)

	// TODO: don't understand this if, but it works...
	if simLinkInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		realPath, err := os.Readlink(path)
		LogError(err)
		path = realPath
	}

	fileInfo, err := os.Stat(path)
	LogError(err)

	if path[len(path)-1:] == string(os.PathSeparator) {
		path = path[:len(path)-1]
	}

	return path, fileInfo
}
