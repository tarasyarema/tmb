package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"sync"

	"github.com/jochasinga/requests"
)

func addAuthQuery(r *requests.Request) {
	r.Params.Add("app_id", apiCreds.AppID)
	r.Params.Add("app_key", apiCreds.AppKey)
}

func fetchLineStopPairs(ps []Pair) ([]Times, error) {
	pairReg := regexp.MustCompile(`\/(\d+)`)

	var data []Times

	pairLen := len(ps)
	urls := make([]string, pairLen)

	for i, p := range ps {
		urls[i] = fmt.Sprintf("%v/ibus/lines/%v/stops/%v", baseURL, p.Line, p.Stop)
	}

	pool := requests.NewPool(pairLen)
	results, err := pool.Get(urls, addAuthQuery)

	if err != nil {
		log.Println(err)
		return data, errors.New("could not create request pool")
	}

	errCount := 0

	for res := range results {
		if res.Error != nil {
			log.Println(res.Error)
			errCount++
			continue
		}

		var tmp APIResponse
		json.Unmarshal(res.Bytes(), &tmp)

		tmpPair := pairReg.FindAllString(res.Response.Request.URL.Path, -1)

		tmpLine, _ := strconv.Atoi(tmpPair[0][1:])
		tmpStop, _ := strconv.Atoi(tmpPair[1][1:])

		tmpTimes := Times{0, Pair{tmpLine, tmpStop}}

		if len(tmp.Data.IBus) > 0 {
			tmpTimes.Time = tmp.Data.IBus[0].TimeS
		}

		data = append(data, tmpTimes)
	}

	if errCount != 0 {
		log.Printf("Error count: %v\n", errCount)
		return data, errors.New("there were errors during pool request resolve")
	}

	return data, nil
}

func fetchLineStopPairAsync(wg *sync.WaitGroup, id int, p Pair, d *Times) {
	defer wg.Done()

	var empty Times
	url := fmt.Sprintf("%v/ibus/lines/%v/stops/%v", baseURL, p.Line, p.Stop)

	rc, err := requests.GetAsync(url)

	if err != nil {
		log.Println(err)
		d = &empty

		return
	}

	res := <-rc

	if res.Error != nil {
		log.Println(err)
		d = &empty

		return
	}

	var tmp APIResponse
	d.Meta = p

	json.Unmarshal(res.Bytes(), &tmp)

	if len(tmp.Data.IBus) > 0 {
		d.Time = tmp.Data.IBus[0].TimeS
	}
}
