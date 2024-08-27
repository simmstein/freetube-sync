package file

import (
	"io/ioutil"
	"strings"
)

func GetLines(file string) []string {
	bytesRead, _ := ioutil.ReadFile(file)
	fileContent := string(bytesRead)

	return strings.Split(fileContent, "\n")
}
