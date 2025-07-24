package models

type WordData struct {
	Word            string `json:"word"`
	Translation     string `json:"translation"`
	SentenceContext string `json:"sentence_context"`
}

type Story struct {
	Uuid  string     `json:"uuid"`
	Title string     `json:"title"`
	Text  string     `json:"text"`
	Words []WordData `json:"words"`
}
