package infrastructure

import (
	"github.com/f-person/analytics-service/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type eventsDB struct {
	mongoClient *mongo.Client
}

func NewEventsDB(mongoClient *mongo.Client) eventsDB {
	return eventsDB{
		mongoClient: mongoClient,
	}
}

func (e eventsDB) CreateEvent(event domain.Event) (domain.Event, error) {
	return domain.Event{}, nil
}
