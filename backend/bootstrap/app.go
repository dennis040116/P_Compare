package bootstrap

import (
	"log"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"github.com/elastic/go-elasticsearch/v8"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
	ES    *elasticsearch.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	log.Println("Connected to MongoDB!")
	app.ES = NewElasticSearchClient(app.Env)
	log.Println("Connected to ElasticSearch!")
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
