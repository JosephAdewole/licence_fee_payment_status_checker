package httpresp

import (
	"encoding/json"
	"net/http"
)

//Response is the data type of response
type Response = _response

type _response struct {
	Status  bool        `json:"status,omitempty"`
	Code    int         `json:"code,omitempty"`
	Name    string      `json:"name,omitempty"` //name of the error
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"` //for errors that occur even if request is successful
	Data    interface{} `json:"data,omitempty"`
}

//Default returns status: true, error : nil and data: data
func Default(data interface{}) _response {
	return _response{
		Status: true,
		Error:  nil,
		Data:   data}
}

//New returns  a new response object
func New(status bool, code int, message string, err, data interface{}) _response {
	return _response{
		Status:  status,
		Code:    code,
		Message: message,
		Error:   err,
		Data:    data,
	}
}

//SetError sets error field
func (r *_response) SetError(err interface{}) *_response {
	r.Error = err
	return r
}

func (r _response) reply(w http.ResponseWriter) {

	w.WriteHeader(r.Code)
	err := json.NewEncoder(w).Encode(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}

//SetMessage sets message field
func (r *_response) SetMessage(message string) *_response {
	r.Message = message
	return r
}

func (r _response) Reply(w http.ResponseWriter) {
	r.reply(w)
}

//ReplyOk replies with status code 200 -  (http - StatusOK)
//
//this is the standard response for successful requests
func (r _response) ReplyOK(w http.ResponseWriter) {
	r.Code = http.StatusOK
	r.reply(w)
}

//ReplyUnprocessableEntity replies with status code 422 -  (http - StatusOK)
//
//this is the standard response for when a request is unsuccessful but it's not
//an error
func (r _response) ReplyUnprocessableEntity(w http.ResponseWriter) {
	r.Code = http.StatusUnprocessableEntity
	r.reply(w)
}

//ReplyCreated replies with status code 201 -  (http - StatusCreated)
//
//the 201 status is sent to indicate that the new resource has been created.
//This is usually used in conjunction with the PUT request type.
func (r _response) ReplyCreated(w http.ResponseWriter) {
	r.Code = http.StatusCreated
	r.reply(w)
}

//ReplyNoContent replies with status code 204
// this is mostly used for delete request or when record of the requested resource is not found"
func (r _response) ReplyNoContent(w http.ResponseWriter) {
	r.Code = http.StatusNoContent
	r.reply(w)
}

//ReplyAccepted replies with status code 202
//
//mostly used when the request has been accepted,
//but not acted upon. The request may or may not be acted upon.
func (r _response) ReplyAccepted(w http.ResponseWriter) {
	r.Code = http.StatusAccepted
	r.reply(w)
}
