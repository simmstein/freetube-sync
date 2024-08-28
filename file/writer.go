package file

import (
	"os"
	"strings"
)

func WriteDatabase(file string, data []string) error {
	content := []byte(strings.Join(data, "\n") + "\n")

	return os.WriteFile(file, content, 0644)
}
