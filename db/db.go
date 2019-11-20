package db

import (
	"vocabulary/config"
	"vocabulary/models"
	"database/sql"
	"fmt"
)

var (
	db *sql.DB
	err error
)

func Connect() error {
	dbConf := config.Peek().Database
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Addr, dbConf.User, dbConf.Pass, dbConf.DBName)
	db, err = sql.Open("postgres", psqlInfo)
	return err
}

func Close() error {
	return db.Close()
}

func GetCards() ([]models.GetCard, error)  {
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
