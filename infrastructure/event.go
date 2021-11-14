package infrastructure

import (
	"github.com/f-person/analytics-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID    primitive.ObjectID `json:"id"              bson:"_id"`
	Name  string             `json:"name"            bson:"name"`
	Data  map[string]string  `json:"data,omitempty"  bson:"data,omitempty"`
	AppID string             `json:"appId,omitempty" bson:"appId,omitempty"`
}

func (e *Event) ToDomain() domain.Event {
	return domain.Event{
		Name: e.Name,
		Data: e.Data,
	}
}

func EventFromDomain(event domain.Event) Event {
	return Event{
		Name: event.Name,
		Data: event.Data,
	}
}
