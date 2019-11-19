package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Card struct {
	ID            int    `json:"ID"`
	Word          string `json:"Word"`
	Transcription string `json:"Transcription"`
	Meaning       string `json:"Meaning"`
	Translation   string `json:"Translation"`
}

func main() {

	http.HandleFunc("/", getpage)

	db, err := sql.Open("mysql", "root:Lampochka95@tcp(127.0.0.1:3306)/wordlist")
	if err != nil {
		fmt.Println("Error connecting to DB:", err)
		return
	}

	var card Card
	err = db.QueryRow("SELECT ID, Word, Transcription, Meaning, Translation FROM elementary where ID = ?", 1).Scan(&card.ID, &card.Word, &card.Transcription, &card.Meaning, &card.Translation)
	if err != nil {
		fmt.Println("Error in query: ", err)
		return
	}
	fmt.Println(card.ID, card.Word, card.Transcription, card.Meaning, card.Translation)

	http.HandleFunc("/page", giveeReponse)
	http.ListenAndServe(":8080", nil)
}

func getpage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}

func giveeReponse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	word := r.FormValue("word")

	w.Write([]byte(word))
}
