package models

type Card struct {
	ID      int    `json:"id"`
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
}
