package APIGO

import (
	"net/http"
	"github.com/boltdb/bolt"
	//"encoding/json"
	"fmt"
	"strings"
)

type SubColumn struct{
    Username string `json:"username"`
    Column string `json:"column"`
}

type DBColumn struct{
	Columns []SubColumn `json:"Columns"`
}

type Columns struct {
	AllColumns []string `json:"AllColumns"`
}

func GetColumnByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db, err := bolt.Open("smo.db",0600,nil)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}else{
		err = db.Update(func(tx *bolt.Tx) error {
			b,_ := tx.CreateBucketIfNotExists([]byte("Column"))
			if b != nil {
				path := strings.Split(r.URL.Path,"/")
				username := path[len(path) - 1]
				data := b.Get([]byte(username))
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

