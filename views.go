package main

import (
	_ "fmt"
	"strings"
	"strconv"
	_ "encoding/json"
	"github.com/gin-gonic/gin"
)

/*
 *
 * Get remaining times
 *
 */

func getTimesStops(c *gin.Context) {
	data := strings.Split(c.Query("data"), ",")
	data_len := len(data)

	if data_len % 2 != 0 {
		c.JSON(400, gin.H{
			"message": "missing data in the query string",
		})

		return
	}

	pair_data := make([]Pair, data_len / 2)

	for i := 0; i < data_len; i += 2 {
		pair_data[i / 2].Line, _ = strconv.Atoi(data[i])
		pair_data[i / 2].Stop, _ = strconv.Atoi(data[i + 1])
	}

	time_data := fetchLineStopPairs(pair_data)

	c.JSON(200, gin.H{
		"message": "OK",
		"data": time_data,
	})
}