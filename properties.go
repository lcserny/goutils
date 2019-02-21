package goutils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type ConfigProperties map[string]string

func (p *ConfigProperties) HasProperty(propertyName string) bool {
	_, exists := (*p)[propertyName]
	return exists
}

func (p *ConfigProperties) GetPropertyAsInt(propertyName string) int {
	i, err := strconv.ParseInt((*p)[propertyName], 0, 32)
	LogFatal(err)
	return int(i)
}

func (p *ConfigProperties) GetPropertyAsString(propertyName string) string {
	return (*p)[propertyName]
}

func (p *ConfigProperties) GetPropertyAsBool(propertyName string) bool {
	return strings.ToLower((*p)[propertyName]) == "true"
}

func ReadPropertiesFile(filename string) *ConfigProperties {
	file, err := os.Open(filename)
	LogFatal(err)
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

	LogFatal(scanner.Err())
	return &config
}
