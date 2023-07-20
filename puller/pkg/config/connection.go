package config

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func ConnectGCS(cfg Config) (*storage.Client, error) {
	ctx := context.Background()
	gcsClient, err := storage.NewClient(ctx, option.WithCredentialsFile(cfg.GoogleCloud.ServiceAccountLocation))
	if err != nil {
		return nil, err
	}

	return gcsClient, nil
}
