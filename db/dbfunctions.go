package db

import (
	"strconv"
	"vocab/models"
)

func GetCard(word string) (models.Card, error) {

	card := models.Card{}
	query := "SELECT id, word, meaning FROM words where word = '" + word + "'"
	err := db.QueryRow(query).Scan(&card.ID, &card.Word, &card.Meaning)
	if err != nil {
		return card, err
	}
	return card, nil
}

func GetCards() ([]models.Card, error) {
	rows, err := db.Query("SELECT id, word, meaning FROM words")
	if err != nil {
		return nil, err
	}
	card := models.Card{}
	cards := []models.Card{}
	for rows.Next() {
		err = rows.Scan(&card.ID, &card.Word, &card.Meaning)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func InsertCard(card *models.Card) (err error) {

	query := "INSERT INTO words(word, meaning) VALUES('" + card.Word + "', '" + card.Meaning + "');"
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return
}

func EditCard(card *models.Card) (err error) {

	query := "UPDATE words SET word = '" + card.Word + "', meaning = '" + card.Meaning + "' WHERE id = " + strconv.Itoa(card.ID)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return
}

func DeleteCard(card *models.Card) (err error) {

	query := "DELETE FROM words WHERE id = " + strconv.Itoa(card.ID)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return
}
