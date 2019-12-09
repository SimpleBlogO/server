package main

import (
	"encoding/json"
	sw "github.com/SimpleBlogO/server/go"
	"github.com/boltdb/bolt"
	"os"
)

func decode(filename string, v1 interface{}) interface{} {
	filePtr, _ := os.Open(filename)
	defer filePtr.Close()
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	switch v1.(type) {
	case sw.Articles:
		v := v1.(sw.Articles)
		decoder.Decode(&v)
		return v
	case sw.DBColumn:
		v := v1.(sw.DBColumn)
		decoder.Decode(&v)
		return v
	case sw.DBReview:
		v := v1.(sw.DBReview)
		decoder.Decode(&v)
		return v
	}
	return v1
}

func main() {
	var b sw.Articles
	v := decode("data/article.json", b).(sw.Articles)
	db, _ := bolt.Open("smo.db", 0600, nil)
	defer db.Close()

	//article
	username := v.AllArticles[0].Author
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("Article"))
		tb, _ := json.Marshal(v)
		b.Put([]byte(username), tb)
		return nil
	})

	//Column
	var b1 sw.DBColumn
	v1 := decode("data/column.json", b1).(sw.DBColumn)
	for j := 0; j < len(v1.Columns); {
		username1 := v1.Columns[j].Username
		initJ := j
		for j < len(v1.Columns) && v1.Columns[j].Username == username1 {
			j++
		}
		var comments []string
		for i := initJ; i < j; i++ {
			comments = append(comments, v1.Columns[i].Column)
		}
		var c = sw.Columns{comments}
		db.Update(func(tx *bolt.Tx) error {
			b, _ := tx.CreateBucketIfNotExists([]byte("Column"))
			tb, _ := json.Marshal(c)
			b.Put([]byte(username1), tb)
			return nil
		})
	}

	//review:
	var b2 sw.DBReview
	v2 := decode("data/review.json", b2).(sw.DBReview)
	for j := 0; j < len(v2.Reviews); {
		username2 := v2.Reviews[j].Username
		id := v2.Reviews[j].ID
		initJ := j
		for j < len(v2.Reviews) && v2.Reviews[j].Username == username2 && v2.Reviews[j].ID == id {
			j++
		}
		var reviews []string
		for i := initJ; i < j; i++ {
			reviews = append(reviews, v2.Reviews[i].Comment)
		}
		d := sw.Reviews{reviews}
		db.Update(func(tx *bolt.Tx) error {
			b, _ := tx.CreateBucketIfNotExists([]byte("Review"))
			tb, _ := json.Marshal(d)
			key := username2 + id
			b.Put([]byte(key), tb)
			return nil
		})
	}
}
