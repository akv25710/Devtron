package logger_service

import "time"

func GetLogDirectories(start, end time.Time) []string {
	var dates []string

	for current := start; current.Before(end) || current.Equal(end); current = current.AddDate(0, 0, 1) {
		dates = append(dates, current.Format("2006-01-02"))
	}

	return dates
}
