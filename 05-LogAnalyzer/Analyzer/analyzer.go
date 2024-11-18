package analyzer

import (
	"errors"
	"strings"
)

func SearchToken(line string) (string, error) {

	for _, word := range strings.Split(line, " ") {
		if word == "ERROR" || word == "INFO" || word == "WARN" {
			return word, nil
		}
	}
	return "", errors.New("type of log not found")
}
