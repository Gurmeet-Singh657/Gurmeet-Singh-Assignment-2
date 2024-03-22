package controllers

import (
	"Assignment/service"
	"log"
	"time"
)

const fetchInterval = 10 * time.Second

func CronJob(query string) error {
	for {
		if err := service.StoreYoutubeDataAsynchronously(query); err != nil {
			log.Printf("Error adding videos: %v", err)
		} else {
			log.Print("Videos Added in Mongodb successfully")
		}
		time.Sleep(fetchInterval)
	}
}