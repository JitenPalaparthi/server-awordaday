package handler

import (
	"fmt"

	"github.com/golang/glog"

	nats "github.com/nats-io/nats.go"
)

// Message is a message content, ideally for a channel
type Message struct {
	Data    []byte
	Subject string
}

var (
	chanMessage chan Message
	NC          *nats.Conn
)

// Inidiate the channel at the beginning of the handler usage
func Init(nc *nats.Conn) {
	if chanMessage == nil {
		chanMessage = make(chan Message, 20)
		glog.Info("----------------> Handler Init")
		go ProcessMessage(nc)
	}
}

func ProcessMessage(nc *nats.Conn) {
	for msg := range chanMessage {
		if nc != nil {
			nc.Publish(msg.Subject, msg.Data)
			nc.Subscribe("tokenizer", func(m *nats.Msg) {
				fmt.Println("Received a message: %s\n", string(m.Data))
			})
		}
		//glog.Info("Nats Connection has been expired")
	}
}
