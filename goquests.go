package goquests

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	URL     string
	Data    map[string]interface{}
	Headers map[string]string
}

type Response struct {
	Status  string
	Headers interface{}
	Body    map[string]interface{}
}

func NewRequest(method string, r Request) (Response, error) {
	postBody, _ := json.Marshal(r.Data)
	Body := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(method, r.URL, Body)
	if err != nil {
		return Response{}, err
	}
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	rawresp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer rawresp.Body.Close()
	response := Response{}
	response.Status = rawresp.Status
	response.Headers = rawresp.Header
	var body map[string]interface{}
	json.NewDecoder(rawresp.Body).Decode(&body)
	response.Body = body
	return response, nil
}

func New() *Request {
	return &Request{}
}

func (r *Request) Post(option Request) (Response, error) {
	response, err := NewRequest("POST", option)
	return response, err
}

/***/
func (r *Request) Get(option Request) (Response, error) {
	response, err := NewRequest("GET", option)
	return response, err
}

func (r *Request) Put(option Request) (Response, error) {
	response, err := NewRequest("PUT", option)
	return response, err
}

func (r *Request) Patch(option Request) (Response, error) {
	response, err := NewRequest("PATCH", option)
	return response, err
}

func (r *Request) Delete(option Request) (Response, error) {
	response, err := NewRequest("DELETE", option)
	return response, err
}

func (r *Request) Trace(option Request) (Response, error) {
	response, err := NewRequest("TRACE", option)
	return response, err
}

func (r *Request) Options(option Request) (Response, error) {
	response, err := NewRequest("OPTIONS", option)
	return response, err
}

func (r *Request) Head(option Request) (Response, error) {
	response, err := NewRequest("Head", option)
	return response, err
}
