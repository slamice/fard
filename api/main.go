package main

import (
	"encoding/json"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/slamice/fard/operations"
	"log"
	"net/http"
	"os"
)

var sessionArray []*r.Session

func initDb() {
	session, err := r.Connect(os.Getenv("RETHINKDB_URL"), "frd")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = r.DbCreate("frd").Run(session).Exec()
	if err != nil {
		log.Println(err)
	}

	err = r.TableCreate("articles").Run(session).Exec()
	if err != nil {
		log.Println(err)
	}

	sessionArray = append(sessionArray, session)
}

func main() {

	initDb()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/new", insertArticle)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Error: %v", err)
	}
}

func insertArticle(res http.ResponseWriter, req *http.Request) {
	session := sessionArray[0]

	b := new(Article)
	json.NewDecoder(req.Body).Decode(b)

	var response r.WriteResponse

	err := r.Table("articles").Insert(b).Run(session).One(&response)
	if err != nil {
		log.Fatal(err)
		return
	}
	data, _ := json.Marshal("{'article':'saved'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	session := sessionArray[0]
	var response []Article

	err := r.Table("articles").Run(session).All(&response)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(response)

	res.Header().Set("Content-Type", "application/json")
	res.Write(data)
}
