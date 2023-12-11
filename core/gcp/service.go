package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io"
	"os"
)

var (
	storageClient *storage.Client
	ctx           context.Context
)

const (
	MB = 1 << 20
)

func init() {
	storageClientKeyFile := os.Getenv("STORAGE_CLIENT_KEY")
	if storageClientKeyFile == "" {
		log.Error("[ERROR] STORAGE_CLIENT_KEY not provided")
		//os.Exit(1)
	}

	ctx = context.Background()

	var err error
	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile(storageClientKeyFile))
	if err != nil {
		log.Error("Unable to initialize storage client. Error: ", err.Error())
		os.Exit(1)
	}
}

func DownloadFolder(bucketName, folderPath, destinationPath string) error {

	bucket := storageClient.Bucket(bucketName)
	query := &storage.Query{
		Prefix:    folderPath,
		Delimiter: "/",
	}
	objects := bucket.Objects(ctx, query)

	for {
		objAttrs, err := objects.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Errorf(err.Error())
			return err
		}

		source := bucket.Object(objAttrs.Name)
		reader, err := source.NewReader(ctx)
		if err != nil {
			return err
		}
		defer reader.Close()

		destinationFile := fmt.Sprintf("%s/%s", destinationPath, objAttrs.Name)
		destination, err := os.Create(destinationFile)
		if err != nil {
			return err
		}
		defer destination.Close()

		if _, err := io.Copy(destination, reader); err != nil {
			return err
		}

	}

	return nil
}
