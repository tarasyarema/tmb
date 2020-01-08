package main

type APIResponse struct {
	Status 	string `json:"status"`
    Data 	ResponseData `json:"data"`
}

type ResponseData struct {
	IBus []IBusData `json:"ibus"`
}

type IBusData struct {
	RouteId 	string `json:"routeId"`
	TimeS 		int `json:"t-in-s"`
	TimeM 		int `json"t-in-min"`
}

type Pair struct {
	Line int
	Stop int
}

type Creds struct {
	AppId  string
	AppKey string
}