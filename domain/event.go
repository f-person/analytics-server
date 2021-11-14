package domain

type Event struct {
	Name string
	Data map[string]string `json:"data,omitempty"`
}
