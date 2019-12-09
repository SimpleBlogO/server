package APIGO

import (
	"net/http"
)

type Review struct {

}

func GetReviewByNameAndID(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

