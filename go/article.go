package APIGO

import (
	"net/http"
	"github.com/boltdb/bolt"
	"encoding/json"
	"fmt"
	"strings"
)

type Article struct {
	ID string `json:"ID"`
	Title string `json:"title"`
	Author string `json:"author"`
	Content string `json:"content"`
}

type Articles struct{
	AllArticles []Article `json:"AllArticles"`
}

func GetUserArticleByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db, err := bolt.Open("smo.db",0600,nil)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}else{
		err = db.Update(func(tx *bolt.Tx) error {
			b,_ := tx.CreateBucketIfNotExists([]byte("Article"))
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

func GetUserArticleByNameAndID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db, err := bolt.Open("smo.db",0600,nil)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}else{
		err = db.Update(func(tx *bolt.Tx) error {
			b,_ := tx.CreateBucketIfNotExists([]byte("Article"))
			if b != nil {
				path := strings.Split(r.URL.Path,"/")
				username := path[len(path) - 2]
				id := path[len(path) - 1]
				data := b.Get([]byte(username))
				var all Articles
				if err := json.Unmarshal(data,&all); err == nil{
					for i := 0 ; i < len(all.AllArticles) ; i ++{
						if all.AllArticles[i].ID == id{
							w.WriteHeader(http.StatusOK)
							dd,_ := json.Marshal(all.AllArticles[i])
							fmt.Fprintf(w,string(dd))
							break;
						}
					}
					w.WriteHeader(http.StatusNotFound)
				}else{
					w.WriteHeader(http.StatusInternalServerError)
				}
			}else{
				w.WriteHeader(http.StatusInternalServerError)
			}
			return nil
		})
		defer db.Close()
	}
}

