package database

import (
	"awordaday/models"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//FindSentences is to find all sentences in the database
func (d *Database) FindSentences() []models.Sentence {
	sentences := []models.Sentence{}
	if d.Client.Find(&sentences).Error != nil {
		return nil
	}
	return sentences
}

// FindSentencesByWord is to find sentences based on a word
func (d *Database) FindSentencesByWord(_word string) []models.Sentence {
	sentences := []models.Sentence{}
	word := d.FindWordByWord(_word)
	if word != nil {
		if d.Client.Where("word_id=?", word.ID).Find(&sentences).Error != nil {
			return nil
		}
		return sentences
	}
	return nil
}

// InsertSentence to insert a new sentence
func (d *Database) InsertSentence(sentense *models.Sentence) (err error) {
	sentense.LastUpdated = time.Now()
	c := d.Client.Create(sentense)

	if c.Error != nil {
		return c.Error
	}
	return nil
}

// UpdateSentence is to update a sentence based on values that is based on map
func (d *Database) UpdateSentence(sentence *models.Sentence, values map[string]interface{}) (err error) {
	c := d.Client.Model(sentence).Updates(values)
	if c.Error != nil {
		return c.Error
	}
	return nil
}

// DeleteSentence is to delete the sentence from the database.It is a hard delete
func (d *Database) DeleteSentence(sentence *models.Sentence) (err error) {
	c := d.Client.Delete(sentence)
	if c.Error != nil {
		return c.Error
	}
	return nil
}
