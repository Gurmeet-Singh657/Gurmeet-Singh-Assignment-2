package main

import (
	"Assignment/controllers"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Get Videos
	router.GET("/api/getVideos", controllers.GetVideosController)

	// go routine for continously fetching the videos related to cricket and storing them continously in mongodb
	go func() {
		if err := controllers.CronJob("cricket"); err != nil {
			log.Fatalf("Error in FetchAndStoreVideos: %v", err)
		}
	}()

	// Running the Server at Localhost:3000
	router.Run(":3000")
}
