package handler

type Resp struct {
	Response string  `json:"response"`
	Body     [][]int `json:"body"`
}
