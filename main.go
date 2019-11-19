package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Card struct {
	ID      int    `json:"ID"`
	Word    string `json:"Word"`
	Meaning string `json:"Meaning"`
}

func main() {

	db, err := sql.Open("mysql", "root:Lampochka95@tcp(127.0.0.1:3306)/vocabulary")
	if err != nil {
		fmt.Println("Error connecting to DB:", err)
		return
	}

	var card Card
	var word string
	fmt.Scanf("%s", &word)
	err = db.QueryRow("SELECT ID, Word, Meaning FROM words where Word = ?", word).Scan(&card.ID, &card.Word, &card.Meaning)
	if err != nil {
		fmt.Println("Error in db.QueryRow(: ", err)
		return
	}
	fmt.Println(card.ID, card.Word, card.Meaning)

	// http.HandleFunc("/", index)
	// http.HandleFunc("/page", giveReponse)
	// http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}

func giveReponse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	word := r.FormValue("word")

	w.Write([]byte(word))
}
