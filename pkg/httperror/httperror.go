package httperror

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

//Error is a type
type Error = _error

//Error ...
type _error struct {
	Code    int         `json:"code"`              //e.g 201, 200, 401
	Status  bool        `json:"status"`            //e.g false, true
	Name    string      `json:"name,omitempty"`    //name of the error
	Err     error       `json:"err,omitempty"`     //standard go error
	Message string      `json:"message,omitempty"` // any message for the error
	Error   interface{} `json:"error,omitempty"`   //custom error
}

//Default returns a new object
//
// status is false by default
func Default(err error) _error {

	//Check for empty(nil) error
	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = ""
	}

	return _error{Status: false, Err: err, Message: msg, Error: msg}
}

//New returns a new object
//
// status is false by default
func New(code int, status bool, name string, Err error, msg string, Error interface{}) _error {
	return _error{
		Code:    code,
		Status:  false,
		Name:    name,
		Err:     Err,
		Message: msg,
		Error:   Error}
}

/************************************SET FIELDS*********************************/

//SetStatus sets status field
func (e *_error) SetStatus(status bool) {
	e.Status = status
}

//SetCode sets code field
func (e *_error) SetCode(code int) {
	e.Code = code
}

//SetName sets name field
func (e *_error) SetName(name string) {
	e.Name = name
}

//SetErr sets err field
func (e *_error) SetErr(er error) {
	e.Err = er
}

//SetMessage sets message field
func (e *_error) SetMessage(message string) {
	e.Message = message
}

//SetError sets error field
func (e *_error) SetError(er interface{}) {
	e.Error = er
}

/**********************************GET FIELDS***************************************/

//GetStatus gets status field
func (e _error) GetStatus() bool {
	return e.Status
}

//GetCode gets code field
func (e _error) GetCode() int {
	return e.Code
}

//GetName gets name field
func (e _error) GetName() string {
	return e.Name
}

//GetErr gets err field
func (e _error) GetErr() error {
	return e.Err
}

//GetMessage sets message field
func (e _error) GetMessage() string {
	return e.Message
}

//GetError gets error field
func (e _error) GetError() interface{} {
	return e.Error
}

/******************************** HTTP REPLY***********************************/
func (e _error) ReplyInternalServerError(w http.ResponseWriter) {
	e.Code = http.StatusInternalServerError
	e.reply(w)
}

func (e _error) ReplyBadRequest(w http.ResponseWriter) {
	e.Code = http.StatusBadRequest
	e.reply(w)
}

//ReplyUnkwownResponse replies with status code 0
func (e _error) ReplyUnkwownResponse(w http.ResponseWriter) {
	e.Code = 0
	e.reply(w)
}

//ReplyUnprocessableEntity replies with status code 422 -  (http - StatusOK)
///
//this is the standard response for when a request is unsuccessful but it's not
//an error
func (e _error) ReplyUnprocessableEntity(w http.ResponseWriter) {
	e.Code = http.StatusUnprocessableEntity
	e.reply(w)
}

//Reply ...
func (e _error) Reply(w http.ResponseWriter) {
	e.reply(w)
}

func (e _error) reply(w http.ResponseWriter) {

	w.WriteHeader(e.Code) //Write http code of error

	if er := json.NewEncoder(w).Encode(e.web()); er != nil { //encode error message
		http.Error(w, er.Error(), http.StatusInternalServerError)
	}
}

//Removes standard go error field
func (e _error) web() interface{} {

	return struct {
		Code    int         `json:"code"`              //e.g 201, 200, 401
		Status  bool        `json:"status"`            //e.g false, true
		Name    string      `json:"name,omitempty"`    //name of the error
		Message string      `json:"message,omitempty"` // any message for the error
		Error   interface{} `json:"error,omitempty"`   //custom error
	}{
		Code:    e.Code,
		Status:  e.Status,
		Name:    e.Name,
		Message: e.Message,
		Error:   e.Error,
	}
}

/**************** OTHERS ************************/

//Returns the error in string format
func (e _error) String() string {
	var status string
	if e.Status {
		status = "true"
	} else {
		status = "false"
	}
	return "code: " + strconv.Itoa(e.Code) + "\nstatus: " + status + "\nerror: " + e.Err.Error()
}

//Read reads the next len(p) bytes from the buffer
//or until the buffer is drained.
//The return value n is the number of bytes read. If the buffer has no data to return,
// err is io.EOF (unless len(p) is zero); otherwise it is nil.
func (e _error) Read(p []byte) (n int, err error) {
	b, err := json.Marshal(e)

	if err == nil {
		buf := bytes.NewBuffer(b)
		n, err = buf.Read(p)
	}

	return
}

//WriteToWriter writes to io.writer interface
func (e _error) WriteToWriter(w io.Writer) (n int, err error) {
	b, err := json.Marshal(e)

	if err == nil {
		n, err = w.Write(b)
	}

	return
}
