package models

import (
	"time"
)

type Audit struct {
	ID       string    `gorm:"primary_key" json:"id"`
	Data     string    `gorm:"data" json:"data"`
	IP       string    `gorm:"ip" json:"ip"`
	Device   string    `gorm:"device" json:"device"`
	URLPath  string    `gorm:"url_path" json:"urlPath"`
	Headers  string    `gorm:"headers" json:"headers"` //map[string][]string `gorm:"sentence" json:"sentence"`
	DateTime time.Time `gorm:"date_time" json:"dateTime"`
}
