package controllers

import (
	"Assignment/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const defaultPageSize = 10

func GetVideosController(c *gin.Context) {

	// To Handle the case when page query is not given
	// 1. First Page Query is taken from the Request
	// 2. Converting the String to Integer
	// 3. From the page number will come and that will be used for request handling
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		page = 1
	}

	// Setting Default Page Size as 10
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize <= 0 {
		pageSize = defaultPageSize
	}

	// Calling the GetPaginated Videos Function to retrieve the videos on particular page and with page Size
	videos, err := service.GetVideosHandler(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, videos)
}