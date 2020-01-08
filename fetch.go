package main

import (
	"fmt"
	"encoding/json"
	"github.com/jochasinga/requests"
)

func addAuthQuery(r *requests.Request) {
	r.Params.Add("app_id", api_creds.AppId)
	r.Params.Add("app_key", api_creds.AppKey)
}

func fetchLineStopPairs(ps []Pair) []IBusData {
	urls := make([]string, len(ps))
	
	for i, p := range ps {
		urls[i] = fmt.Sprintf("%v/ibus/lines/%v/stops/%v", base_url, p.Line, p.Stop)
	}

	pool := requests.NewPool(len(urls))

	results, err := pool.Get(urls, addAuthQuery)
	if err != nil {
		panic(err)
	}
	
	var data []IBusData

	for res := range results {
		if res.Error != nil {
			panic(res.Error)
		}

		var tmp APIResponse
		json.Unmarshal(res.Bytes(), &tmp)

		if len(tmp.Data.IBus) > 0 {
			data = append(data, tmp.Data.IBus[0])
		}
	}

	return data
}