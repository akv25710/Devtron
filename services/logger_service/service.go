package logger_service

import (
	"fetchlogger/conf"
	"time"
)

type LoggerService struct {
	CloudService CloudLoggerInterface
	Bucket       string
}

type CloudProvider string

const (
	GCP CloudProvider = "gcp"
)

func getCloudLogger(cloud CloudProvider) CloudLoggerInterface {
	if cloud == GCP {
		return GCPLogger{}
	}

	return GCPLogger{}
}

func InitLoggerService(conf conf.LoggerConfiguration) LoggerService {
	cloud := getCloudLogger(conf.Cloud)

	return LoggerService{
		CloudService: cloud,
		Bucket:       conf.Bucket,
	}
}

func (l LoggerService) GetLogs(text string, start, end time.Time) ([]string, error) {
	return l.CloudService.SearchLogs(text, start, end)
}
