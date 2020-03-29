package handler

import (
	"awordaday/database"
	"awordaday/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMagicWord(d *database.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			word, err := d.FindMagicWord()
			fmt.Println(err)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, word)
		}
	}
}

func GetAllWords(d *database.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			words, err := d.FindAllWords()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, words)
		}
	}
}

func GetWords(d *database.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			skip := c.Param("skip")
			limit := c.Param("limit")

			if skip == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": "skip parameter has not been provieded",
				})
				c.Abort()
				return
			}

			if limit == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": "limit parameter has not been provieded",
				})
				c.Abort()
				return
			}

			iskip, err := strconv.ParseInt(skip, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err,
				})
				c.Abort()
				return
			}

			ilimit, err := strconv.ParseInt(limit, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err,
				})
				c.Abort()
				return
			}

			words, err := d.FindWords(iskip, ilimit)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, words)
		}
	}
}

// InsertWord is to insert a new Word
func InsertWord(d *database.Database) func(c *gin.Context) {
	var err error
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			var word models.Word
			word = models.Word{}
			word.Status = "NOT-ACTIVE"
			//word.LastUpdated = time.Now()
			err = json.NewDecoder(c.Request.Body).Decode(&word)
			fmt.Println(word.Word)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}

			err = d.InsertWord(&word)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			//c.BindJSON(&u)
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Word Successfully Created",
			})
			c.Abort()
			return
		}
	}
}

// InsertWord is to insert a new Word
func InsertRequestedWord(d *database.Database) func(c *gin.Context) {
	var err error
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			var requestWord models.RequestWord
			requestWord = models.RequestWord{}
			requestWord.Status = "NOT-ACTIVE"
			//word.LastUpdated = time.Now()
			err = json.NewDecoder(c.Request.Body).Decode(&requestWord)
			fmt.Println(requestWord.Word)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}

			err = d.InsertRequestedWord(&requestWord)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			//c.BindJSON(&u)
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Requested word successfully added",
			})
			c.Abort()
			return
		}
	}
}

func DeleteWord(d *database.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == "DELETE" {
			word := c.Param("word")
			if word == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": "word parameter has not been provieded",
				})
				c.Abort()
				return
			}
			err := d.DeleteWord(word)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Requested word successfully deleted",
			})
			c.Abort()
			return
		}
	}
}

// InsertSentence is to insert a new Word
func InsertSentence(d *database.Database) func(c *gin.Context) {
	var err error
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			var sentence models.Sentence
			sentence = models.Sentence{}
			sentence.Status = "NOT-ACTIVE"
			err = json.NewDecoder(c.Request.Body).Decode(&sentence)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}

			err = d.InsertSentence(&sentence)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			//c.BindJSON(&u)
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Sentence Successfully Created",
			})
			c.Abort()
			return
		}
	}
}

func UpdateWord(d *database.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == "PUT" {
			id := c.Param("id")
			if id == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": "Id parameter has not been provieded",
				})
				c.Abort()
				return
			}

			jsonMap := make(map[string]interface{})
			err := json.NewDecoder(c.Request.Body).Decode(&jsonMap)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "body seems to be wrong json format",
					"message": err.Error(),
				})
				c.Abort()
				return
			}

			err = d.UpdateWord(id, jsonMap)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Word successfulluy updated",
			})
			c.Abort()
			return
		}
	}
}
