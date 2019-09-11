package gcsexport

import (
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
)

func Export(inputReader io.Reader, gcsObject string) (int64, error) {
	gcsURL, err := url.Parse(gcsObject)
	if err != nil {
		return 0, err
	}
	if gcsURL.Scheme != "gs" {
		return 0, errors.New("URL should start with gs://")
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return 0, err
	}

	objectName := gcsURL.Path[1:]

	fmt.Printf("uploading to bucket %s and object %s\n", gcsURL.Host, objectName)

	writer := client.Bucket(gcsURL.Host).Object(objectName).NewWriter(ctx)
	defer writer.Close() // make sure we close
	bytesCopied, err := io.Copy(writer, inputReader)
	if err != nil {
		return 0, err
	}
	if err := writer.Close(); err != nil {
		return 0, err
	}
	return bytesCopied, nil
}
