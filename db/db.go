package db

import (
	"database/sql"
	"fmt"
	"vocab/config"
	"vocab/models"
)

var (
	db  *sql.DB
	err error
)

func Connect() error {
	dbConf := config.Peek().Database
	psqlInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s",
		dbConf.User, dbConf.Pass, dbConf.Addr, dbConf.DBName)
	db, err = sql.Open("mysql", psqlInfo)

	return err
}

func Close() error {
	return db.Close()
}

func GetCard(word string) (models.GetCard, error) {

	card := models.GetCard{}
	err := db.QueryRow("SELECT id, word, meaning FROM words where word = ?", word).Scan(&card.ID, &card.Word, &card.Meaning)
	if err != nil {
		return card, err
	}
	return card, nil
}

func GetCards() ([]models.GetCard, error) {
	rows, err := db.Query("SELECT id, word, meaning FROM words")
	if err != nil {
		return nil, err
	}
	card := models.GetCard{}
	cards := []models.GetCard{}
	for rows.Next() {
		err = rows.Scan(&card.ID, &card.Word, &card.Meaning)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}
