package main

import (
	"bytes"
	_ "encoding/json"
	_ "fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "INFO: ", log.Lshortfile)

	infof = func(info string) {
		logger.Output(2, info)
	}
)

func getTimesStopsPool(c *gin.Context) {
	defer timeTrack(time.Now(), "getTimesStopsPool")

	_t := time.Now()

	data := strings.Split(c.Query("data"), ",")
	dataLen := len(data)

	if dataLen%2 != 0 {
		c.JSON(400, gin.H{
			"message": "missing data in the query string",
		})

		return
	}

	pairData := make([]Pair, dataLen/2)

	for i := 0; i < dataLen; i += 2 {
		pairData[i/2].Line = data[i]
		pairData[i/2].Stop = data[i+1]
	}

	ret := gin.H{"message": "OK"}
	code := 200

	timeData, err := fetchLineStopPairs(pairData)

	if err != nil {
		code = 400
		ret["message"] = err
		ret["elapsed"] = time.Since(_t).Milliseconds()

		c.JSON(code, ret)
		return
	}

	if timeData != nil {
		ret["data"] = timeData
	}

	ret["elapsed"] = time.Since(_t).Milliseconds()

	c.JSON(code, ret)
}

func getTimesStopsRoutines(c *gin.Context) {
	defer timeTrack(time.Now(), "getTimesStopsRoutines")

	_t := time.Now()

	data := strings.Split(c.Query("data"), ",")
	dataLen := len(data)

	if dataLen%2 != 0 {
		c.JSON(400, gin.H{
			"message": "missing data in the query string",
		})

		return
	}

	pairData := make([]Pair, dataLen/2)

	for i := 0; i < dataLen; i += 2 {
		pairData[i/2].Line = data[i]
		pairData[i/2].Stop = data[i+1]
	}

	dataLen /= 2
	timeData := make([]Times, dataLen)

	isSync, _ := strconv.ParseBool(c.DefaultQuery("sync", "false"))

	var wg sync.WaitGroup
	wg.Add(dataLen)

	for i := 0; i < dataLen; i++ {
		if isSync {
			go fetchLineStopPairSync(&wg, i, pairData[i], &timeData[i])
		} else {
			go fetchLineStopPairAsync(&wg, i, pairData[i], &timeData[i])
		}
	}

	wg.Wait()

	ret := gin.H{"message": "OK"}
	code := 200

	if timeData != nil {
		ret["data"] = timeData
	}

	ret["elapsed"] = time.Since(_t).Milliseconds()

	c.JSON(code, ret)
}
