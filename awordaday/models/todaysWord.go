package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type TodaysWord struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Today     time.Time `json:today`
	Status    string    `json:"status"`
	UpdatedBy string    `gorm:"updated_by" json:"updated_by`
}
