package main

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/rohanhonnakatti/go-nats-basic/config"
	"github.com/rohanhonnakatti/go-nats-basic/models"
)

func consumeReviews(js nats.JetStreamContext) {
	_, err := js.Subscribe(config.SubjectNameCreateReview, func(m *nats.Msg) {
		err := m.Ack()

		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}

		var review models.Review
		err = json.Unmarshal(m.Data, &review)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Consumer => \nSubject: %s\nID: %s\nAuthor: %s\nRating: %d\n\n", m.Subject, review.Id, review.Author, review.Rating)
	})

	if err != nil {
		log.Fatalf("Error while Subscring message: %v", err)
	}
}
