package infrastructure

import (
	"github.com/f-person/analytics-service/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type eventsDB struct {
	collection *mongo.Collection
}

func NewEventsDB(mongoClient *mongo.Client) eventsDB {
	collection := mongoClient.Database(dbName).Collection(eventsCollectionName)

	return eventsDB{
		collection: collection,
	}
}

func (e eventsDB) CreateEvent(event domain.Event) (domain.Event, error) {

	return domain.Event{}, nil
}
