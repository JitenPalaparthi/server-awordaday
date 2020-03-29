package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Word struct {
	ID          string     `gorm:"primary_key" json:"id"`
	Word        string     `json:"word"`
	Meaning     string     `json:"meaning"`
	Type        string     `json:"type"`
	Sentences   []Sentence `json:"sentences" gorm:"foreignkey:word_id"`
	Status      string     `json:"status"`
	UpdatedBy   string     `gorm:"updated_by" json:"updatedBy`
	LastUpdated time.Time  `gorm:"last_updated" json:"lastUpdated"`
}

//ValidateForInsert is to insert the record into the database
func (w *Word) ValidateForInsert() error {
	if w.Word == "" {
		return GetFieldValidationError("Word")
	}
	if w.Meaning == "" {
		return GetFieldValidationError("Meaning")
	}
	if w.Type == "" {
		return GetFieldValidationError("Type")
	}
	return nil
}

type RequestWord struct {
	ID          string    `gorm:"primary_key" json:"id"`
	Word        string    `json:"word"`
	Status      string    `json:"status"`
	RequestedBy string    `gorm:"requested_by" json:"requestedBy`
	LastUpdated time.Time `gorm:"last_updated" json:"lastUpdated"`
}
