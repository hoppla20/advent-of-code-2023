package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func SetLogLevel() {
	level := slog.LevelInfo
	s, ok := os.LookupEnv("LOG_LEVEL")
	if ok {
		switch s {
		case "DEBUG":
			level = slog.LevelDebug
		case "INFO":
			level = slog.LevelInfo
		case "WARN":
			level = slog.LevelWarn
		case "ERROR":
			level = slog.LevelError
		default:
			slog.Warn(fmt.Sprintf("Unknown log level '%s'", s))
			level = slog.LevelInfo
		}
	}
	slog.SetLogLoggerLevel(level)
}

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func Lowest(items []int) (int, error) {
	if len(items) == 0 {
		return 0, errors.New("Can not calculate lowest item of an empty list!")
	}

	lowest := items[0]
	for _, item := range items[1:] {
		if item < lowest {
			lowest = item
		}
	}

	return lowest, nil
}
