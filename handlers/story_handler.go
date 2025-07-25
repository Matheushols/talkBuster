package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"talkBuster/models"
	"talkBuster/services"
)

var Storys = make([]models.Story, 0)
var WordsData = make([]models.WordData, 0)

// CreateStoryHandler is used to create a new story
func CreateStoryHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var story models.Story
	err = json.Unmarshal(body, &story)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validLanguage := services.LanguageValidation(story.Language)
	if !validLanguage {
		http.Error(w, fmt.Sprintf("Language %s is not valid", story.Language), http.StatusBadRequest)
		return
	}

	story = services.CreateUuid(story)
	Storys = append(Storys, story)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(story)
}
