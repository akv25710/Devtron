package logger_service

import "time"

type CloudLoggerInterface interface {
	SearchLogs(text string, start, end time.Time) ([]string, error)
}
