package controllers

import (
	"log"
	"net/http"
	"vocabulary/db"
)

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
