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
func main() {
	n_errors := 0
	n_warns := 0
	n_info := 0
	n_line := 1
	sliceLogs := make([]string, 0)

	initFlags()
	logFile, err := os.Open("logs/log.txt")
	if err != nil {
		log.Panicf("ERROR	Failed to open logFile: %v", err)
	}
	defer logFile.Close()

	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		sliceLogs = append(sliceLogs, scanner.Text())
		typeLog, err := analyzer.SearchToken(scanner.Text())
		if err != nil {
			log.Printf("\n\nWARN    Warn reading the log file, Line %d: %v\n\n", n_line, err)
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
	if all {
		for _, log := range sliceLogs {
			fmt.Println(log)
		}
	}
	if eror {
		for _, log := range sliceLogs {
			if strings.Contains(log, "ERROR") {
				fmt.Println(log)
			}
		}
	}
	if warn {
		for _, log := range sliceLogs {
			if strings.Contains(log, "WARN") {
				fmt.Println(log)
			}
		}
	}

	if info {
		for _, log := range sliceLogs {
			if strings.Contains(log, "INFO") {
				fmt.Println(log)
			}
		}
	}

	if count {
		fmt.Printf("Numer of ERRORS:%d Number of WARNS:%d Numer of INFO:%d \n", n_errors, n_warns, n_info)
	}

}
