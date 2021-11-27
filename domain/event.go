package domain

type Event struct {
	ID    string
	Name  string
	Data  map[string]string
	AppID string
}
