package infrastructure

import (
	"github.com/f-person/analytics-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type App struct {
	ID    primitive.ObjectID `json:"id"              bson:"_id"`
	Name  string             `json:"name"            bson:"name"`
	Token string             `json:"token,omitempty" bson:"token,omitempty"`
}

func (a *App) ToDomain() domain.App {
	return domain.App{
		ID:    a.ID.String(),
		Name:  a.Name,
		Token: a.Token,
	}
}

func AppFromDomain(app domain.App) App {
	return App{
		ID:    objectIDFromString(app.ID),
		Name:  app.Name,
		Token: app.Token,
	}

}
