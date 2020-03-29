package models

import (
	"time"
)

type Sentence struct {
	ID          string    `gorm:"primary_key" json:"id"`
	WordID      string    `gorm:"word_id" json:"wordId"`
	Sentence    string    `gorm:"sentence" json:"sentence"`
	Status      string    `json:"status"`
	UpdatedBy   string    `gorm:"updated_by" json:"updatedBy"`
	LastUpdated time.Time `gorm:"last_updated" json:"lastUpdated"`
}

//ValidateForInsert is to insert the record into the database
func (s *Sentence) ValidateForInsert() error {

	if s.WordID == "" {
		return GetFieldValidationError("WordID")
	}
	if s.Sentence == "" {
		return GetFieldValidationError("Sentence")
	}
	return nil
}
