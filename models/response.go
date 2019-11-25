package models

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func (response *Response) Send(to http.ResponseWriter) {
	data, _ := json.Marshal(response)
	to.Header().Set("Content-Type", "application/json")
	to.WriteHeader(response.Code)
	to.Write(data)
}
