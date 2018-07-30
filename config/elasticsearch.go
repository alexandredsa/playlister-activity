package config

import (
	"errors"
	"log"
	"os"
	"sync"

	"github.com/olivere/elastic"
)

var once sync.Once
var client *elastic.Client

func StartClient(url string) error {
	var err error
	var c *elastic.Client
	once.Do(func() {
		c, err = elastic.NewSimpleClient(elastic.SetURL(url), elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)))
		client = c
	})
	return err
}

func StopClient() error {
	if client == nil {
		return createErrorClientNotStarted()
	}
	defer client.Stop()

	return nil
}

func GetClient() (*elastic.Client, error) {
	if client == nil {
		return nil, createErrorClientNotStarted()
	}

	return client, nil
}

func createErrorClientNotStarted() error {
	return errors.New("Client Not Started")
}

func CheckBulkResponse(bulkResponse *elastic.BulkResponse) error {
	failures := bulkResponse.Failed()
	if len(failures) == 0 {
		return nil
	}

	msgError := ""
	for i, failure := range failures {
		if i > 0 {
			msgError += ";"
		}
		msgError += failure.Error.Reason
	}

	return errors.New(msgError)
}
