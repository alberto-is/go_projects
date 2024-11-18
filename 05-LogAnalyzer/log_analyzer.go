package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	analyzer "loganalyzer/Analyzer"
	"os"
	"strings"
)

var (
	count bool
	all   bool
	eror  bool
	warn  bool
	info  bool
)

func initFlags() {
	flag.BoolVar(&count, "count", false, "Show the total count of each log type")
	flag.BoolVar(&count, "c", false, "Show the total count of each log type (short)")

	flag.BoolVar(&all, "all", false, "Show all log entries")
	flag.BoolVar(&all, "a", false, "Show all log entries (short)")

	flag.BoolVar(&eror, "error", false, "Show only ERROR logs")
	flag.BoolVar(&eror, "e", false, "Show only ERROR logs (short)")

	flag.BoolVar(&warn, "warn", false, "Show only WARN logs")
	flag.BoolVar(&warn, "w", false, "Show only WARN logs (short)")

	flag.BoolVar(&info, "info", false, "Show only INFO logs")
	flag.BoolVar(&info, "i", false, "Show only INFO logs (short)")

	flag.Parse()
}

func processLogFile(filename string) ([]string, map[string]int) {
	logFile, err := os.Open("logs/log.txt")
	if err != nil {
		log.Panicf("ERROR	Failed to open logFile: %v", err)
	}
	defer logFile.Close()

	logs := make([]string, 0)
	counts := map[string]int{"ERROR": 0, "WARN": 0, "INFO": 0}

	n_line := 1
	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
		typeLog, err := analyzer.SearchToken(scanner.Text())
		if err != nil {
			log.Printf("\n\nWARN    Warn reading the log file, Line %d: %v\n\n", n_line, err)
			continue

		}
		counts[typeLog]++
		n_line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("ERROR: Failed to scan the log file: %v", err)
	}

	return logs, counts
}

func printFilteredLogs(logs []string, logType string) {
	fmt.Printf("\n%s Logs:\n", logType)
	for _, log := range logs {
		if strings.Contains(log, logType) {
			fmt.Println(log)
		}
	}
}

func main() {
	initFlags()

	logs, counts := processLogFile("logs/log.txt")

	if all {
		fmt.Println("All Logs:")
		for _, log := range logs {
			fmt.Println(log)
		}
	}
	if eror {
		printFilteredLogs(logs, "ERROR")
	}
	if warn {
		printFilteredLogs(logs, "WARN")
	}
	if info {
		printFilteredLogs(logs, "INFO")
	}

	if count {
		fmt.Printf("\nLog Counts:\nERROR: %d\nWARN: %d\nINFO: %d\n", counts["ERROR"], counts["WARN"], counts["INFO"])
	}

}
