package database

import (
	"awordaday/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InsertAudit to insert a new sentence
func (d *Database) InsertAudit(audit *models.Audit) (err error) {
	c := d.Client.Create(audit)
	if c.Error != nil {
		return c.Error
	}
	return nil
}
