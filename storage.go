package main

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
)

// Client & Context Google Cloud
type Client struct {
	CTX context.Context
	GCS *storage.Client
}

// NewClient Google Cloud
func NewClient() (Client, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return Client{}, err
	}

	return Client{
		CTX: ctx,
		GCS: client,
	}, nil
}

// Write content in object GCS
func (c Client) Write(bucket, object string, content io.Reader) error {
	wc := c.GCS.Bucket(bucket).Object(object).NewWriter(c.CTX)
	if _, err := io.Copy(wc, content); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

// https://stackoverflow.com/questions/20876780/how-to-append-write-to-google-cloud-storage-file-from-app-engine/20876882
// Compose(..) {   c.GCS.Bucket(bucket).Object(object).ComposerFrom().Run(c.CTX) }
