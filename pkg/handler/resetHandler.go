package handler

import (
	"encoding/json"
	"kr2/pkg"
	"net/http"
)

func Reset(w http.ResponseWriter, r *http.Request) {
	var Respon Resp
	TURN = pkg.X
	Respon.Response = pkg.XTurn
	for i := range Field {
		Field[i] = make([]int, 3, 3)
	}
	Respon.Body = Field
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Respon)
}
