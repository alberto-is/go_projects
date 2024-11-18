package analyzer

import (
	"errors"
	"strings"
)

func SearchToken(line string) (string, error) {

	if strings.Contains(line, "ERROR") {
		return "ERROR", nil
	} else if strings.Contains(line, "INFO") {
		return "INFO", nil
	} else if strings.Contains(line, "WARN") {
		return "WARN", nil
	}
	return "", errors.New("type of log not found")
}
