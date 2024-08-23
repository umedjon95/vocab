package db

import (
	"math/rand"
	"strconv"

	"github.com/umedjon95/vocab/models"
)

// GetCard returns searched word card from DB
func GetCard(word string) ([]models.Card, error) {

	query := "SELECT id, word, meaning FROM words where word LIKE '%" + word + "%';"
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

// GetCards returns all the cards from DB
func GetCards() ([]models.Card, error) {

	rows, err := db.Query("SELECT id, word, meaning FROM words;")
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

// RandCard returns 4 random cards from DB
func RandCard() ([]models.Card, error) {

	rows, err := db.Query(`	SELECT 
								id,
								word,
								meaning
							FROM
								words
							WHERE
								id IN
								(
									SELECT
										trunc (
											random () * (select max(id) FROM words) + 1 
										)
									FROM
										generate_series (1, 4)
								);`)
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

	for i, _ := range cards {
		if i == 0 {
			cards[rand.Intn(3)+1].Meaning = cards[i].Meaning
			cards[i].Meaning = ""
		} else {
			cards[i].ID = 0
			cards[i].Word = ""
		}
	}

	return cards, nil
}

// CheckCard check weather card exist in DB
func CheckCard(card *models.Card) (exists bool, err error) {

	query := `SELECT EXISTS(
		SELECT
			(1)
		FROM
			words
		WHERE
			id = ` + strconv.Itoa(card.ID) + ` and
			word = '` + card.Word + `' and
			meaning = '` + card.Meaning + `');`

	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		err = rows.Scan(&exists)
		if err != nil {
			return false, err
		}
	}
	if err != nil {
		return false, err
	}

	return
}

// InsertCard insertes new card to DB
func InsertCard(card *models.Card) (err error) {

	query := "INSERT INTO words(word, meaning) VALUES('" + card.Word + "', '" + card.Meaning + "');"
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return
}

// EditCard edit card in DB from taken id
func EditCard(card *models.Card) (err error) {

	query := "UPDATE words SET word = '" + card.Word + "', meaning = '" + card.Meaning + "' WHERE id = " + strconv.Itoa(card.ID)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return
}

// DeleteCard deletes card from DB from taken id
func DeleteCard(id int) (err error) {

	query := "DELETE FROM words WHERE id = " + strconv.Itoa(id)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return
}
