package db

import (
	"strconv"
	"vocab/models"
)

//GetCard returns searched word card from database
func GetCard(word string) ([]models.Card, error) {

	query := "SELECT id, word, meaning FROM words where word LIKE '%" + word + "%'"
	rows, err := db.Query(query)
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

//GetCards returns all the cards from database
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

//InsertCard insertes new card to database
func InsertCard(card *models.Card) (err error) {

	query := "INSERT INTO words(word, meaning) VALUES('" + card.Word + "', '" + card.Meaning + "');"
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return
}

//EditCard edit card in database from taken id
func EditCard(card *models.Card) (err error) {

	query := "UPDATE words SET word = '" + card.Word + "', meaning = '" + card.Meaning + "' WHERE id = " + strconv.Itoa(card.ID)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return
}

//DeleteCard deletes card from database from taken id
func DeleteCard(id int) (err error) {

	query := "DELETE FROM words WHERE id = " + strconv.Itoa(id)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return
}
