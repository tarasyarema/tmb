package main

// APIResponse ... Root object of the TMB iBus service response
type APIResponse struct {
	Status string       `json:"status"`
	Data   ResponseData `json:"data"`
}

// ResponseData ...
type ResponseData struct {
	IBus []IBusData `json:"ibus"`
}

// IBusData ...
type IBusData struct {
	RouteID string `json:"routeId"`
	TimeS   int    `json:"t-in-s"`
	TimeM   int    `json:"t-in-min"`
}

// Times ...
type Times struct {
	Time    int   `json:"time"`
	Meta    Pair  `json:"meta"`
	Elapsed int64 `json:"elapsed"`
}

// Pair ...
type Pair struct {
	Line string `json:"line"`
	Stop string `json:"stop"`
}

// Creds ...
type Creds struct {
	AppID  string
	AppKey string
}
