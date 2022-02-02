package requests

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

func NewRequest(url string, data map[string]interface{}, headers map[string]string) *Request {
	return &Request{
		URL:     url,
		Data:    data,
		Headers: headers,
	}
}
func (r *Request) Post() (Response, error) {
	postBody, _ := json.Marshal(r.Data)
	Body := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", r.URL, Body)
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

/***/
func (r *Request) Get() (Response, error) {
	postBody, _ := json.Marshal(r.Data)
	Body := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("GET", r.URL, Body)
	if err != nil {
		return Response{}, err
	}
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	response := Response{}
	response.Status = resp.Status
	response.Headers = resp.Header
	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	response.Body = body
	return response, nil
}

func (r *Request) Put() (Response, error) {
	postBody, _ := json.Marshal(r.Data)
	Body := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", r.URL, Body)
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

func (r *Request) Patch() (Response, error) {
	postBody, _ := json.Marshal(r.Data)
	Body := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PATCH", r.URL, Body)
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
