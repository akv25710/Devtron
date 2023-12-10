package logger_service

import "time"

func GetLogDirectories(startDate, endDate time.Time) []string {
	var dates []string

	for current := startDate; current.Before(endDate) || current.Equal(endDate); current = current.AddDate(0, 0, 1) {
		dates = append(dates, current.Format("2006-01-02"))
	}

	return dates
}
