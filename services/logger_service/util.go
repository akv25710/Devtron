package logger_service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func GetLogDirectories(start, end time.Time) []string {
	var dates []string

	for current := start; current.Before(end) || current.Equal(end); current = current.AddDate(0, 0, 1) {
		dates = append(dates, current.Format("2006-01-02"))
	}

	return dates
}

func GetLogFiles(path string, folders []string, start, end time.Time) []string {
	var result []string

	for _, folder := range folders {
		dest := fmt.Sprintf("%s/%s", path, folder)

		files, err := os.ReadDir(dest)
		if err != nil {
			logrus.Error("Error reading directory:", err)
			continue
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			fileDate, _ := GetDateFromDirectoryAndFile(folder, file.Name())

			if fileDate.After(start) && fileDate.Before(end) {
				result = append(result, fmt.Sprintf("%s/%s", dest, file.Name()))
			}
		}
	}

	return result
}

func GetDateFromDirectoryAndFile(folder, file string) (time.Time, error) {
	fileName := strings.Split(file, ".")[0]

	dateTimeStr := fmt.Sprintf("%sT%s:00:00Z00:00", folder, fileName)

	return time.Parse(time.RFC3339, dateTimeStr)
}
