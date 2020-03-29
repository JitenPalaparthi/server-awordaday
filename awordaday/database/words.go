package database

import (
	"awordaday/models"
	"errors"
	"fmt"
	"strings"
	"time"
)

//FindAllWords is to find all words from the database
func (d *Database) FindAllWords() (words []models.Word, err error) {
	//words := []models.Word{}
	err = d.Client.Preload("Sentences").Order("last_updated DESC").Find(&words).Error
	if err != nil {
		return nil, err
	}
	return words, nil
}

//FindAllWords is to find all words from the database
func (d *Database) FindWords(skip, limit int64) (words []models.Word, err error) {
	//words := []models.Word{}
	err = d.Client.Preload("Sentences").Order("last_updated DESC").Limit(limit).Offset(skip).Find(&words).Error
	if err != nil {
		return nil, err
	}
	return words, nil
}

// FindWordByWord is to find a Word by a actual word
func (d *Database) FindWordByWord(_word string) *models.Word {
	word := &models.Word{}
	if d.Client.Where("Upper(word)=?", strings.ToUpper(_word)).First(word).Error != nil {
		return nil
	}
	return word
}

// FindMagicWord is to find a Word by a actual word
func (d *Database) FindMagicWord() (word *models.Word, err error) {
	if d.Client != nil {
		word = &models.Word{}
		err := d.Client.Where("status='Active'").Preload("Sentences").Order("last_updated DESC").First(&word).Error
		if err != nil {
			return nil, err
		}
		if word != nil {
			return word, nil
		}
		if word == nil {
			return nil, errors.New("no records")
		}

	}
	return nil, errors.New("no records")
}

//InsertWord is to insert a nee word
func (d *Database) InsertWord(word *models.Word) (err error) {
	fmt.Println(word)
	if d.FindWordByWord(word.Word) == nil {
		word.LastUpdated = time.Now()
		c := d.Client.Create(word)
		if c.Error != nil {
			return c.Error
		}
		return nil
	}
	fmt.Println(word)

	return errors.New("Word already existed")
}

//InsertWord is to insert a nee word
func (d *Database) InsertRequestedWord(requestedWord *models.RequestWord) (err error) {
	//fmt.Println(requestedWord)
	if d.FindWordByWord(requestedWord.Word) == nil {
		requestedWord.LastUpdated = time.Now()
		c := d.Client.Create(requestedWord)
		if c.Error != nil {
			return c.Error
		}
		return nil
	}
	return errors.New("Word already existed")
}

// UpdateWord is to update a word based on values that are by map
func (d *Database) UpdateWord(id string, values map[string]interface{}) (err error) {
	if len(values) < 1 {
		return errors.New("no values provided to update")
	}
	values["last_updated"] = time.Now()
	word := &models.Word{}
	word.ID = id
	c := d.Client.Model(word).Updates(values)
	if c.Error != nil {
		return c.Error
	}
	return nil
}

func (d *Database) DeleteWord(_word string) (err error) {
	word := &models.Word{}
	c := d.Client.Where("word = ?", _word).First(word)
	if c.Error != nil {
		fmt.Println(c.Error)
		return c.Error
	}
	sentence := &models.Sentence{}
	c1 := d.Client.Where("word_id = ?", word.ID).Delete(sentence)
	if c1.Error != nil {
		fmt.Println(c.Error)
		return c.Error
	}

	c2 := d.Client.Where("word = ?", _word).Delete(word)
	if c2.Error != nil {
		fmt.Println(c.Error)
		return c.Error
	}
	return nil
}
