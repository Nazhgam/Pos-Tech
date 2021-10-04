package helper

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	sqls "pos-tech/db"
	structures "pos-tech/models"
)

var Db sqls.SqlDb

func init() {
	d, err := sql.Open("sqlite3", "db/data.db")
	if err != nil {
		log.Println("NO DATABASE IN THIS LAPTOP")
		return
	}
	Db.DB = d
}
func GetPostById(w http.ResponseWriter, r *http.Request) {
	data := &structures.Task{}
	if r.Method == http.MethodPost {
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		data, err = Db.Get(data.Id)
		if err != nil {
			w.WriteHeader(500)
		}
	}
	b, _ := json.Marshal(data)
	w.WriteHeader(200)
	_, err := w.Write(b)
	if err != nil {
		return
	}

}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	data := &structures.Task{}
	if r.Method == http.MethodPost {
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = Db.WriteToSql(data)
		if err != nil {
			w.WriteHeader(500)
		}
	}
	w.WriteHeader(200)
}
func Update(w http.ResponseWriter, r *http.Request) {
	data := &structures.Task{}
	if r.Method == http.MethodPost {
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = json.Unmarshal(jsonData, &data)

		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = Db.UpdateSql(data)
		if err != nil {
			w.WriteHeader(500)
		}
	}
	w.WriteHeader(200)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	data := &structures.Task{}
	if r.Method == http.MethodPost {
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = json.Unmarshal(jsonData, &data)

		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = Db.DeleteFromSql(data.Id)
		if err != nil {
			w.WriteHeader(500)
		}
	}
	w.WriteHeader(200)
}
