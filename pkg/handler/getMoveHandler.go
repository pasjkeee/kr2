package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kr2/pkg"
	"net/http"
	"strconv"
)

func GetMove(w http.ResponseWriter, r *http.Request) {

	var Respon Resp

	x, err := strconv.Atoi(mux.Vars(r)["x"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(mux.Vars(r)["y"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if Field[x][y] != pkg.EMPTY {
		Respon.Response = pkg.NoEmpty
		json.NewEncoder(w).Encode(&Respon)
		return
	}

	Field[x][y] = TURN
	Respon.Body = Field

	if CheckRow(Respon, w) || CheckColumn(Respon, w) {
		return
	}

	//диагональ главаная
	cur := Field[0][0]
	badRes := CheckDiagonal(cur)
	if !badRes && Won(cur, Respon, w) {
		return
	}

	//диагональ не главаная
	cur = Field[0][len(Field)-1]
	badRes = CheckDiagonalNotMain(cur)

	if !badRes && Won(cur, Respon, w) {
		return
	}

	isEmpty := false
	for _, row := range Field {

		if isEmpty == true {
			break
		}
		for _, rowItem := range row {
			if rowItem == pkg.EMPTY {
				isEmpty = true
				break
			}
		}
	}

	if isEmpty == false {
		Respon.Response = pkg.NoWon
		for i := range Field {
			Field[i] = make([]int, 3, 3)
		}
		Respon.Body = Field
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Respon)
		return
	}

	if TURN == pkg.X {
		TURN = pkg.O
		Respon.Response = pkg.OTurn
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Respon)
	} else {
		TURN = pkg.X
		Respon.Response = pkg.XTurn
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Respon)
	}
}
