package handler

import (
	"encoding/json"
	"kr2/pkg"
	"net/http"
)

func Won(cur int, Respon Resp, w http.ResponseWriter) bool {
	for i := range Field {
		Field[i] = make([]int, 3, 3)
	}
	if cur == pkg.O {
		Respon.Response = pkg.OWon
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Respon)
		return true
	} else {
		Respon.Response = pkg.XWon
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Respon)
		return true
	}
}

func CheckRow(Respon Resp, w http.ResponseWriter) bool {
MainLoop1:
	for i := 0; i < len(Field); i++ {
		cur := Field[i][0]

		if cur == pkg.EMPTY {
			continue
		}

		for j := 1; j < len(Field[i]); j++ {
			if cur != Field[i][j] {
				continue MainLoop1
			}
		}

		if Won(cur, Respon, w) {
			return true
		}
	}

	return false
}

func CheckColumn(Respon Resp, w http.ResponseWriter) bool {

MainLoop2:
	for i := 0; i < len(Field); i++ {
		cur := Field[0][i]

		if cur == pkg.EMPTY {
			continue
		}

		for j := 1; j < len(Field[i]); j++ {
			if cur != Field[j][i] {
				continue MainLoop2
			}
		}

		if Won(cur, Respon, w) {
			return true
		}
	}

	return false
}

func CheckDiagonal(cur int) bool {
	for i := 1; i < len(Field); i++ {
		if cur == pkg.EMPTY {
			return true
		}

		if cur != Field[i][i] {
			return true
		}
	}

	return false
}

func CheckDiagonalNotMain(cur int) bool {
	for i := 1; i < len(Field); i++ {
		if cur == pkg.EMPTY {
			return true
		}

		if cur != Field[i][len(Field)-1-i] {
			return true
		}
	}

	return false
}
