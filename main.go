package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	loadEnv()

	r := gin.Default()

	r.GET("/pool", getTimesStopsPool)
	r.GET("/routines", getTimesStopsRoutines)

	portStr := fmt.Sprintf(":%v", port)

	r.Run(portStr)
}
