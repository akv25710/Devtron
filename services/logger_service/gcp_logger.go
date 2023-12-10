package logger_service

import (
	"bufio"
	"fmt"
	"lfetchogger/core/gcp"
	"os"
	"strings"
	"time"
)

type GCPLogger struct {
	Parallelism chan int
	BucketName  string
}

func (g GCPLogger) SearchLogs(text string, start, end time.Time) ([]string, error) {
	folders := GetLogDirectories(start, end)

	err := g.DownloadFolders(folders, "/tmp")
	if err != nil {
		return nil, err
	}

	return g.searchFiles(text, folders), nil
}

func (g GCPLogger) searchFiles(text string, folders []string) []string {
	var result []string

	for _, folder := range folders {
		for i := 0; i < 24; i++ {
			dest := fmt.Sprintf("/tmp/%s/%02d.txt", folder, i)
			result = append(result, g.searchFile(text, dest)...)
		}
	}

	return result
}

func (g GCPLogger) searchFile(text, path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, text) {
			result = append(result, line)
		}
	}

	return result
}

func (g GCPLogger) DownloadFolders(folders []string, dest string) error {
	for _, folder := range folders {
		g.Parallelism <- 1

		folder := folder
		go func() {
			defer func() {
				<-g.Parallelism
			}()

			_ = gcp.DownloadFolder(g.BucketName, folder, fmt.Sprintf("%s/%s", dest, folder))
		}()
	}

	return nil
}
