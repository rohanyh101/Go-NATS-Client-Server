package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rohanhonnakatti/go-nats-basic/config"
	"github.com/rohanhonnakatti/go-nats-basic/models"
)

func publishReviews(js nats.JetStreamContext) {
	reviews, err := getReviews()
	if err != nil {
		log.Println(err)
		return
	}

	for _, review := range reviews {
		r := rand.Intn(1500)
		time.Sleep(time.Duration(r) * time.Millisecond)

		reviewString, err := json.Marshal(review)
		if err != nil {
			log.Printf("failed to marshal review: %v", err)
			continue
		}

		// publish review to REVIEWS.CreateRate
		_, err = js.Publish(config.SubjectNameCreateReview, reviewString)
		if err != nil {
			log.Printf("Error while publishing review: %v", err)
			continue
		}
		log.Printf("Publisher  =>  Message: %s\n", review.Text)
	}
}

func getReviews() ([]models.Review, error) {
	rawReviews, err := os.ReadFile("./reviews.json")
	if err != nil {
		return nil, fmt.Errorf("error reading ./reviews.json: %v", err)
	}

	var reviewsObj []models.Review
	err = json.Unmarshal(rawReviews, &reviewsObj)

	return reviewsObj, err
}
