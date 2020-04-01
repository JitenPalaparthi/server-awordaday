package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"awordaday/models"

	"github.com/golang/glog"

	nats "github.com/nats-io/nats.go"
)

// Message is a message content, ideally for a channel
type Message struct {
	Data    []byte
	Subject string
}

var (
	chanMessage  chan Message
	NC           *nats.Conn
	ChanWord     chan models.Word
	BotToken     string = "1129175375:AAHiJ0FMhGOZmj-nIKg488jNMtAEqoyIUHY"
	BotChatId    int64  = -402080650
	BotUrlstring        = "https://api.telegram.org/bot" + BotToken + "/sendMessage" //"https://api.telegram.org/bot1129175375:AAHiJ0FMhGOZmj-nIKg488jNMtAEqoyIUHY/sendMessage"
)

type TelgramBotData struct {
	ChatID    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

// Inidiate the channel at the beginning of the handler usage
func Init(nc *nats.Conn) {
	if chanMessage == nil {
		chanMessage = make(chan Message, 20)
		glog.Info("----------------> Handler Init")
		go ProcessMessage(nc)
	}
	if ChanWord == nil {
		ChanWord = make(chan models.Word, 20)
		glog.Info("----------------> Handler Init for Word")

		//go ProduceWord(nc)
		go SubscribeWord(nc)
	}
}

func ProcessMessage(nc *nats.Conn) {
	if nc != nil {
		for {
			select {
			case msg := <-chanMessage:
				nc.Publish(msg.Subject, msg.Data)
			case word := <-ChanWord:
				byteWord, err := json.Marshal(word)
				if err != nil {
					glog.Error(err)

				}
				err = nc.Publish("ATipADayBot", byteWord)
				if err != nil {
					glog.Error(err)

				}
			default:
			}
		}
	}

}
func SubscribeWord(nc *nats.Conn) {
	if nc != nil {
		nc.Subscribe("ATipADayBot", func(m *nats.Msg) {
			var _word models.Word
			if err := json.Unmarshal(m.Data, &_word); err != nil {
				glog.Error(err)
			}
			fmt.Println(_word)
			//today := time.Now().Format("Mon Jan _2 15:04:05 2006")
			today := time.Now().Format("Mon Jan _2 2006")
			//ConstructText := "*Word On " + today + "*\n*Word:        " + _word.Word + "*\n*Meaning:  " + _word.Meaning + "*\n*Type:        " + _word.Type + "*\n*Sentences:*\n*1 " + _word.Sentences[0].Sentence + "*\n*2 " + _word.Sentences[1].Sentence + "*\n*3 " + _word.Sentences[2].Sentence + "*"
			//
			ConstructText := "<h3>Word On " + today + "</h3><br><h3>Word:        " + _word.Word + "</h3><br><h3>Meaning:  " + _word.Meaning + "</h3><br><h4>Type:        " + _word.Type + "</h4><br><h3>Sentences:</h3><br><h4>1 " + _word.Sentences[0].Sentence + "</h4><br><h4>2 " + _word.Sentences[1].Sentence + "</h4><br><h4>3 " + _word.Sentences[2].Sentence + "</h4>"

			reqBody := &TelgramBotData{
				ChatID:    BotChatId,
				Text:      ConstructText,
				ParseMode: "html",
			}

			reqBytes, err := json.Marshal(reqBody)
			if err != nil {
				glog.Error(err)
			}

			// Create the request body struct

			// Create the JSON body from the struct

			// Send a post request with your token
			res, err := http.Post(BotUrlstring, "application/json", bytes.NewBuffer(reqBytes))
			if err != nil {
				glog.Error(err)
			}
			if res.StatusCode != http.StatusOK {
				glog.Error("unexpected status" + res.Status)
			}

			//http.Post(BotUrlstring, "application/json")

		})
	}

}
