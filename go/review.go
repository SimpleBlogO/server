package APIGO

import (
	"net/http"
	"github.com/boltdb/bolt"
	//"encoding/json"
	"fmt"
	"strings"
)

type Reviews struct {
	AllReviews []string `json:"AllReviews"`
}

func GetReviewByNameAndID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db, err := bolt.Open("smo.db",0600,nil)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}else{
		err = db.Update(func(tx *bolt.Tx) error {
			b,_ := tx.CreateBucketIfNotExists([]byte("Review"))
			if b != nil {
				path := strings.Split(r.URL.Path,"/")
				username := path[len(path) - 2]
				id := path[len(path) - 1]
				key := username + id
				data := b.Get([]byte(key))
				if data == nil{
					w.WriteHeader(http.StatusNotFound)
				}else{
					w.WriteHeader(http.StatusOK)
					fmt.Fprintf(w,string(data))
				}
			}else{
				w.WriteHeader(http.StatusInternalServerError)
			}
			return nil
		})
		defer db.Close()
	}
}

