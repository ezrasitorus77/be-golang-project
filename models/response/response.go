package response

import (
	"encoding/json"
	"net/http"
)

type (
	Response struct {
		W http.ResponseWriter
	}

	responseBody struct {
		RC      string      `json:"RC"`
		Message string      `json:"Message"`
		Data    interface{} `json:"Data"`
		Error   error       `json:"Error"`
	}
)

func (resp *Response) SendResponse(status int, responseCode, message string, data interface{}, err error) {
	var (
		jsonEncoder *json.Encoder = json.NewEncoder(resp.W)
		rBody       responseBody  = responseBody{}
	)

	resp.W.Header().Set("Content-Type", "application/json")
	resp.W.WriteHeader(status)

	rBody = responseBody{
		RC:      responseCode,
		Message: message,
		Data:    data,
		Error:   err,
	}
	jsonEncoder.Encode(rBody)

	return
}
