package entities

type Message struct {
	ID      string `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Content string `json:"content,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
}
