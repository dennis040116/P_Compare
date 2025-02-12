package bootstrap

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticSearchClient(env *Env) *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s:%s", env.ESHost, env.ESPort),
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the elasticsearch client: %s", err)
	}

	return es
}
