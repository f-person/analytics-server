package infrastructure

import (
	"github.com/f-person/analytics-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID    primitive.ObjectID `json:"id"              bson:"_id"`
	Name  string             `json:"name"            bson:"name"`
	Data  map[string]string  `json:"data,omitempty"  bson:"data,omitempty"`
	AppID primitive.ObjectID `json:"appId,omitempty" bson:"appId,omitempty"`
}

func (e *Event) ToDomain() domain.Event {
	return domain.Event{
		ID:    e.ID.String(),
		Name:  e.Name,
		Data:  e.Data,
		AppID: e.AppID.String(),
	}
}

func EventFromDomain(event domain.Event) Event {
	return Event{
		ID:    objectIDFromString(event.ID),
		Name:  event.Name,
		Data:  event.Data,
		AppID: objectIDFromString(event.AppID),
	}
}
