package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"vocab/db"
	"vocab/models"

	"github.com/julienschmidt/httprouter"
)

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

func GetCard(to http.ResponseWriter, from *http.Request, params httprouter.Params) {

	var response *models.Response
	key := params.ByName("key")
	card, err := db.GetCard(key)
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
		Payload: card,
	}
	response.Send(to)
	return
}

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

func DeleteCard(to http.ResponseWriter, from *http.Request, params httprouter.Params) {

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
	err = db.DeleteCard(&card)
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
