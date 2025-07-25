package models

type WordData struct {
	Word            string `json:"word"`
	Translation     string `json:"translation"`
	SentenceContext string `json:"sentence_context"`
	Language        string `json:"language"`
}

type Story struct {
	Uuid     string     `json:"uuid"`
	Title    string     `json:"title"`
	Text     string     `json:"text"`
	Language string     `json:"language"`
	Words    []WordData `json:"words"`
}
