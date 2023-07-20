package adapter

import (
	"errors"
	"fmt"

	config "github.com/thnkrn/comet/puller/pkg/config"
	domain "github.com/thnkrn/comet/puller/pkg/domain"
	client "github.com/thnkrn/comet/puller/pkg/driver/client"
	adapterClient "github.com/thnkrn/comet/puller/pkg/driver/client/adapter"
)

type RocksdbAdmin struct {
	CometAPIUrl string
	httpClient  client.HTTP
}

func NewRocksdbAdmin(cfg config.Config, httpClient client.HTTP) *RocksdbAdmin {
	return &RocksdbAdmin{cfg.App.APIUrl, httpClient}
}

func (r *RocksdbAdmin) Ingest(authorization, db, directory string) error {
	var response domain.IngestResponse
	url := fmt.Sprintf("%s/admin/databases/%s/ingests/%s", r.CometAPIUrl, db, directory)
	if err := r.httpClient.Put(url, nil, &response, adapterClient.WithAuth("Bearer "+authorization)); err != nil {
		return err
	}
	if response.Error != "" {
		return errors.New(response.Error)
	}

	return nil
}

func (r *RocksdbAdmin) GetLastIngest(authorization, db string) (string, error) {
	var response domain.GetLastIngestResponse
	url := fmt.Sprintf("%s/admin/databases/%s/ingests/last", r.CometAPIUrl, db)
	if err := r.httpClient.Get(url, nil, &response, adapterClient.WithAuth("Bearer "+authorization)); err != nil {
		return "", err
	}
	if response.StatusCode == 404 {
		return "", nil
	}
	if response.Error != "" {
		return "", errors.New(response.Error)
	}
	return response.Result, nil
}
