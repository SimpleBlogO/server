package APIGO

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"net/http"
	"strings"
)

//json tag struct []byte capital
type User struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus string `json:"userStatus"`
}

//curl localhost:8080/v1/user -X POST  -d 'username=12312312&email=123&password=123&phone=23&userStatus=123'
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	db, err := bolt.Open("smo.db", 0600, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		err = db.Update(func(tx *bolt.Tx) error {
			b, _ := tx.CreateBucketIfNotExists([]byte("User"))
			if b != nil {
				username := r.Form["username"][0]
				data := b.Get([]byte(username))
				if data == nil {
					user := User{
						r.Form["username"][0],
						r.Form["email"][0],
						r.Form["password"][0],
						r.Form["phone"][0],
						r.Form["userStatus"][0]}
					tb, _ := json.Marshal(user)
					b.Put([]byte(username), tb)
					w.WriteHeader(http.StatusOK)
					//return fmt.Errorf("create bucket: %s", err)
				} else {
					//var user User
					//if err := json.Unmarshal(data,&user); err == nil{
					//	fmt.Print(user)
					w.WriteHeader(http.StatusBadRequest)
					//}else{
					//	fmt.Print(err)
					//}
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return nil
		})
		defer db.Close()
	}
}

// curl localhost:8080/v1/user/123123123 -i
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db, err := bolt.Open("smo.db", 0600, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		err = db.Update(func(tx *bolt.Tx) error {
			b, _ := tx.CreateBucketIfNotExists([]byte("User"))
			if b != nil {
				path := strings.Split(r.URL.Path, "/")
				username := path[len(path)-1]
				data := b.Get([]byte(username))
				if data == nil {
					w.WriteHeader(http.StatusNotFound)
				} else {
					b.Delete([]byte(username))
					w.WriteHeader(http.StatusOK)
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return nil
		})
		defer db.Close()
	}
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	db, err := bolt.Open("smo.db", 0600, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		err = db.Update(func(tx *bolt.Tx) error {
			b, _ := tx.CreateBucketIfNotExists([]byte("User"))
			if b != nil {
				path := strings.Split(r.URL.Path, "/")
				username := path[len(path)-1]
				data := b.Get([]byte(username))
				if data == nil {
					w.WriteHeader(http.StatusNotFound)
				} else {
					w.WriteHeader(http.StatusOK)
					fmt.Fprintf(w, string(data))
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return nil
		})
		defer db.Close()
	}
}

//curl localhost:8080/v1/user/login\?username=12312312\&password=123 -i
func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	db, err := bolt.Open("smo.db", 0600, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		err = db.Update(func(tx *bolt.Tx) error {
			b, _ := tx.CreateBucketIfNotExists([]byte("User"))
			if b != nil {
				username := r.Form["username"][0]
				password := r.Form["password"][0]
				fmt.Println(username, password)
				data := b.Get([]byte(username))
				if data == nil {
					w.WriteHeader(http.StatusBadRequest)
				} else {
					var user User
					if err := json.Unmarshal(data, &user); err == nil {
						if password == user.Password {
							w.WriteHeader(http.StatusOK)
						} else {
							w.WriteHeader(http.StatusBadRequest)
						}
					} else {
						w.WriteHeader(http.StatusInternalServerError)
					}
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return nil
		})
		defer db.Close()
	}
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// curl -i localhost:8080/v1/user/12312312 -X PUT -d 'username=12312312&email=123&password=1234&phone=23&userStatus=123'
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ParseForm()
	db, err := bolt.Open("smo.db", 0600, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		err = db.Update(func(tx *bolt.Tx) error {
			b, _ := tx.CreateBucketIfNotExists([]byte("User"))
			if b != nil {
				path := strings.Split(r.URL.Path, "/")
				username := path[len(path)-1]
				data := b.Get([]byte(username))
				if data == nil {
					w.WriteHeader(http.StatusNotFound)
				} else {
					w.WriteHeader(http.StatusOK)
					user := User{
						r.Form["username"][0],
						r.Form["email"][0],
						r.Form["password"][0],
						r.Form["phone"][0],
						r.Form["userStatus"][0]}
					tb, _ := json.Marshal(user)
					b.Put([]byte(username), tb)
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return nil
		})
		defer db.Close()
	}
}
