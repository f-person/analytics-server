package domain

type Account struct {
	ID           string
	EmailAddress string
	Apps         []App
}
