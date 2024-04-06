package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/rohanhonnakatti/go-nats-basic/config"
)

func JetStreamInit() (nats.JetStreamContext, error) {

	// connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %v", err)
	}

	// create a jet stream context
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, fmt.Errorf("failed to create jet stream context: %v", err)
	}

	err = CreateStream(js)
	if err != nil {
		return nil, err
	}

	return js, nil
}

func CreateStream(jetStream nats.JetStreamContext) error {
	stream, _ := jetStream.StreamInfo(config.StreamName)

	if stream == nil {
		log.Printf("Creating stream: %s\n", config.StreamName)

		_, err := jetStream.AddStream(&nats.StreamConfig{
			Name:     config.StreamName,
			Subjects: []string{config.StreamSubjects},
		})

		if err != nil {
			return fmt.Errorf("error in createStream function: %v", err)
		}
	}

	return nil
}
