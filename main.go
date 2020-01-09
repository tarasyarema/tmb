package main

import "github.com/gin-gonic/gin"

func main() {
	loadEnv()

	r := gin.Default()

	r.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.GET("/lbs", getTimesStops)
	
	r.Run()
}