package controllers

import (
	"log"
	"net/http"
	"text/template"
	"vocab/db"
)

func GetPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}

func GetCard(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println("r.ParseForm err:", err)
		return
	}

	word := r.FormValue("word")

	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	card, err := db.GetCard(word)
	if err != nil {
		w.Write([]byte("<p>Ошибка при обработке запроса</p>"))
		log.Println("db.GetCard err:", err)
		return
	}

	page := "<p>" + card.Word + " - " + card.Meaning + "</p>"

	w.Write([]byte(page))
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	cards, err := db.GetCards()
	if err != nil {
		w.Write([]byte("<p>Ошибка при обработке запроса</p>"))
		log.Println("db.GetCards err:", err)
		return
	}
	page := ""
	for _, card := range cards {
		page += "<p>" + card.Word + " - " + card.Meaning + "</p>"
	}

	w.Write([]byte(page))
}
