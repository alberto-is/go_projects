package main

import (
	"bufio"
	"fmt"
	"log"
	analyzer "loganalyzer/Analyzer"
	"os"
)

func main() {
	n_errors := 0
	n_warns := 0
	n_info := 0
	logFile, err := os.Open("logs/log.txt")

	if err != nil {
		log.Panicf("ERROR	Failed to open logFile: %v", err)
	}

	defer logFile.Close()

	scanner := bufio.NewScanner(logFile)

	n_line := 1
	for scanner.Scan() {
		typeLog, err := analyzer.SearchToken(scanner.Text())

		if err != nil {
			log.Printf("WARN    Line %d: %v\n", n_line, err)
		} else {
			switch typeLog {
			case "ERROR":
				n_errors++
			case "WARN":
				n_warns++
			case "INFO":
				n_info++
			}
		}
		n_line++
	}
	fmt.Printf("Numer of ERRORS:%d Number of WARNS:%d Numer of INFO:%d \n", n_errors, n_warns, n_info)

}
