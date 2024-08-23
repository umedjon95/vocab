package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"vocab/db"
	"vocab/models"

	"github.com/julienschmidt/httprouter"
)

// Ping check state of the project
func Ping(to http.ResponseWriter, from *http.Request, params httprouter.Params) {

	var response *models.Response

	response = &models.Response{
		Code:    200,
		Message: "OK",
		Payload: "Working",
	}
	response.Send(to)
}

// GetCards gets all cards from DB
func GetCards(to http.ResponseWriter, from *http.Request, params httprouter.Params) {

	var response *models.Response

	cards, err := db.GetCards()
	if err != nil {
		response = &models.Response{
			Code:    500,
			Message: "DATABASE_ERROR: " + err.Error(),
		}
		response.Send(to)
		log.Println("db.GetCards err:", err)
		return
	}
	response = &models.Response{
		Code:    200,
		Message: "OK",
		Payload: cards,
	}
	response.Send(to)
}

// RandCard gets 4 random cards from DB
func RandCard(to http.ResponseWriter, from *http.Request, params httprouter.Params) {

	var response *models.Response

	cards, err := db.RandCard()
	if err != nil {
		response = &models.Response{
			Code:    500,
			Message: "DATABASE_ERROR: " + err.Error(),
		}
		response.Send(to)
		log.Println("db.RandCard err:", err)
		return
	}
	response = &models.Response{
		Code:    200,
		Message: "OK",
		Payload: cards,
	}
	response.Send(to)
}

// CheckCard weather card exists in DB
func CheckCard(to http.ResponseWriter, from *http.Request, params httprouter.Params) {

	var response *models.Response

	// 1. get word object;
	data, err := ioutil.ReadAll(from.Body)
	if err != nil {
		response = &models.Response{
			Code:    400,
			Message: "BAD_REQUEST: " + err.Error(),
		}
		response.Send(to)
		return
	}
	card := models.Card{}
	err = json.Unmarshal(data, &card)
	if err != nil {
		response = &models.Response{
			Code:    401,
			Message: "BAD_REQUEST: " + err.Error(),
		}
		response.Send(to)
		return
	}

	// 2. check database;
	exists, err := db.CheckCard(&card)
	if err != nil {
		log.Println(err.Error())
		response = &models.Response{
			Code:    500,
			Message: "DATABASE_ERROR: " + err.Error(),
		}
		response.Send(to)
		return
	}
	// 3. say ok;
	response = &models.Response{
		Code:    200,
		Message: "OK",
		Payload: exists,
	}
	response.Send(to)
}

// GetCard gets cards from DB by given word
func GetCard(to http.ResponseWriter, from *http.Request, params httprouter.Params) {

	var response *models.Response
	key := params.ByName("key")
	cards, err := db.GetCard(key)
	if err != nil {
		response = &models.Response{
			Code:    500,
			Message: "DATABASE_ERROR: " + err.Error(),
		}
		response.Send(to)
		return
	}

	response = &models.Response{
		Code:    200,
		Message: "OK",
		Payload: cards,
	}
	response.Send(to)
	return
}

// InsertCard insertes a new card to DB
func InsertCard(to http.ResponseWriter, from *http.Request, params httprouter.Params) {

	var response *models.Response

	// 1. get word object;
	data, err := ioutil.ReadAll(from.Body)
	if err != nil {
		response = &models.Response{
			Code:    400,
			Message: "BAD_REQUEST: " + err.Error(),
		}
		response.Send(to)
		return
	}
	card := models.Card{}
	err = json.Unmarshal(data, &card)
	if err != nil {
		response = &models.Response{
			Code:    401,
			Message: "BAD_REQUEST: " + err.Error(),
		}
		response.Send(to)
		return
	}

	// 2. validate;

	// 3. insert into database;
	err = db.InsertCard(&card)
	if err != nil {
		response = &models.Response{
			Code:    500,
			Message: "DATABASE_ERROR: " + err.Error(),
		}
		response.Send(to)
		return
	}
	// 4. say ok;
	response = &models.Response{
		Code:    200,
		Message: "OK",
	}
	response.Send(to)
}

// EditCard edites a card by given id
func EditCard(to http.ResponseWriter, from *http.Request, params httprouter.Params) {

	var response *models.Response

	// 1. get word object;
	data, err := ioutil.ReadAll(from.Body)
	if err != nil {
		response = &models.Response{
			Code:    400,
			Message: "BAD_REQUEST: " + err.Error(),
		}
		response.Send(to)
		return
	}
	card := models.Card{}
	err = json.Unmarshal(data, &card)
	if err != nil {
		response = &models.Response{
			Code:    401,
			Message: "BAD_REQUEST: " + err.Error(),
		}
		response.Send(to)
		return
	}

	// 2. validate;

	// 3. edit the word;
	err = db.EditCard(&card)
	if err != nil {
		response = &models.Response{
			Code:    500,
			Message: "DATABASE_ERROR: " + err.Error(),
		}
		response.Send(to)
		return
	}
	// 4. say ok;
	response = &models.Response{
		Code:    200,
		Message: "OK",
	}
	response.Send(to)
}

// DeleteCard deletes a card by given id
func DeleteCard(to http.ResponseWriter, from *http.Request, params httprouter.Params) {

	var response *models.Response
	id, err := strconv.Atoi(params.ByName("key"))
	if err != nil {
		response = &models.Response{
			Code:    400,
			Message: "BAD_REQUEST: " + err.Error(),
		}
		response.Send(to)
		return
	}
	// 1. get word object;
	data, err := ioutil.ReadAll(from.Body)
	if err != nil {
		response = &models.Response{
			Code:    400,
			Message: "BAD_REQUEST: " + err.Error(),
		}
		response.Send(to)
		return
	}
	card := models.Card{}
	err = json.Unmarshal(data, &card)
	if err != nil {
		response = &models.Response{
			Code:    401,
			Message: "BAD_REQUEST: " + err.Error(),
		}
		response.Send(to)
		return
	}

	// 2. validate;

	// 3. edit the word;
	err = db.DeleteCard(id)
	if err != nil {
		response = &models.Response{
			Code:    500,
			Message: "DATABASE_ERROR: " + err.Error(),
		}
		response.Send(to)
		return
	}
	// 4. say ok;
	response = &models.Response{
		Code:    200,
		Message: "OK",
	}
	response.Send(to)
}
