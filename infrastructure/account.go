package infrastructure

import (
	"github.com/f-person/analytics-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID           primitive.ObjectID `json:"id"    bson:"_id"`
	EmailAddress string             `json:"email" bson:"email"`
	Apps         []App              `json:"apps"  bson:"apps"`
}

func (a *Account) ToDomain() domain.Account {
	apps := make([]domain.App, len(a.Apps))
	for _, app := range a.Apps {
		apps = append(apps, app.ToDomain())
	}

	return domain.Account{
		ID:           a.ID.String(),
		EmailAddress: a.EmailAddress,
		Apps:         apps,
	}
}

func AccountFromDomain(account domain.Account) Account {
	apps := make([]App, len(account.Apps))
	for _, app := range account.Apps {
		apps = append(apps, AppFromDomain(app))
	}

	return Account{
		ID:           objectIDFromString(account.ID),
		EmailAddress: account.EmailAddress,
		Apps:         apps,
	}
}
