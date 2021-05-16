package httprequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//RequestWithBearerToken takes a url, token and rawQuery as arguments and returns
//the request data as []byte (bytes)
//
// e.g url =  http://localhost:8090
//
//e.g token =  rR324%23#4524+42_-52r2...
//
// e.g params = { {"topic_id": "45"}, {"class_id":"43"}}
func RequestWithBearerToken(method, url, token string, params map[string]string, body io.Reader) (*http.Response, error) {

	if size := len(params); size > 0 {
		url = url + "?"
		for key, value := range params {
			if size > 1 {
				size--
				url = fmt.Sprintf(url+"%s=%s&", key, value)
			} else {
				url = fmt.Sprintf(url+"%s=%s", key, value)
			}
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer " + token)
	req.Header.Set("Content-Type", "application/json")

	resp, er := http.DefaultClient.Do(req)
	if er != nil {
		return nil, er
	}

	return resp, nil
}

//NewReader return an io.Reader for objects
func NewReader(data interface{}) (io.Reader, error) {

	Dbyte, er := json.Marshal(data)
	if er != nil {
		return nil, er
	}

	buf := bytes.NewBuffer(Dbyte)

	return buf, nil
}
