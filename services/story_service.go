package services

import (
	"github.com/google/uuid"
	"talkBuster/models"
)

func CreateUuid(story models.Story) models.Story {
	id := uuid.New()
	story.Uuid = id.String()
	return story
}
