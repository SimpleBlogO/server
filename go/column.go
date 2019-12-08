package 

import (
	"net/http"
)

type Column struct {

}

func GetColumnByName(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

