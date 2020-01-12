package main

import "github.com/gin-gonic/gin"

func main() {
	loadEnv()

	r := gin.Default()

	r.GET("/pool", getTimesStopsPool)
	r.GET("/routines", getTimesStopsRoutines)

	r.Run()
}
