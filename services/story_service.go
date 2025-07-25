package services

import (
	"github.com/google/uuid"
	"talkBuster/models"
)

// CreateUuid is used to make a new Uuid for Story
func CreateUuid(story models.Story) models.Story {
	id := uuid.New()
	story.Uuid = id.String()
	return story
}

// Valids languages to use
var validLanguages = map[string]bool{
	"pt-br": true,
	"en":    true,
	"es":    true,
	"fr":    true,
}

// LanguageValidation is used to check whether the specified language is valid
func LanguageValidation(language string) bool {
	return validLanguages[language]
}
