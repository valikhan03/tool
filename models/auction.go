package tool_models

import (
	"github.com/google/uuid"
)

type Auction struct {
	ID              uuid.UUID `json:"id"`
	Slug            string    `json:"slug"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	MaxParticipants int       `json:"max_participants"`
	ParticipantsNum int       `json:"participants_num"`
	Status          string    `json:"status"`
	Type			int		  `json:"type"`
}
