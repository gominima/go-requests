//requests is a go package to simplify net/http api requests
//you can declare a variable Request with headers data/body and url and
//make requests with it
package requests

//check
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	//"strings"
)

//Request Struct is to be declared before making a request
//Ex: ```
// req := Request{Url: "https://domain.api/abc",
//                Data: map[string]interface{}{"My request":"hello world"},
//                Headers: map[string]string{"Authorization":"Token 123jioaf"}}
// ```
type Request struct {
	Url     string
	Data    map[string]interface{}
	Headers map[string]string
}

//Response struct is to be returned after making request
type Response struct {
	Status     string
	Headers    interface{}
	Body       map[string]interface{}
	StatusCode int
}

//Post method is to be used after declaring req with a Data field
// ```req := Request{Url: "https://domain.api/abc",
//                Data: map[string]interface{}{"My request":"hello world"},
//                Headers: map[string]string{"Authorization":"Token 123jioaf"}}
// res, err := req.Post()
// body := res.Body //This is of type map[string]interface{}
// status := res.Status
// Headers := res.Headers
// statuscode := res.StatusCode```
func (r *Request) Post() (Response, error) {
	postBody, _ := json.Marshal(r.Data)
	responseBody := bytes.NewBuffer(postBody)
	//resp, err := http.Post(r.Url, r.Headers, responseBody)
	req, err := http.NewRequest("POST", r.Url, responseBody)
	if err != nil {
		return Response{}, err
	}
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")
	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	response := Response{}
	response.Status = resp.Status
	response.StatusCode = resp.StatusCode
	response.Headers = resp.Header
	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	response.Body = body
	return response, nil
}

func (r *Request) Get() (Response, error) {
	req, err := http.NewRequest("GET", r.Url, nil)

	if err != nil {
		return Response{}, err
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
		delete(r.Headers, key)
		break
	}
	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	response := Response{}
	response.Status = resp.Status
	response.StatusCode = resp.StatusCode
	response.Headers = resp.Header
	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	response.Body = body
	return response, nil
}

func (r *Request) PostFormUrlEncoded() (Response, error) {
	data := url.Values{}

	for key, value := range r.Data {
		data.Set(key, fmt.Sprintf("%v", value))
		delete(r.Data, key)
		break
	}
	for key, value := range r.Data {
		data.Add(key, fmt.Sprintf("%v", value))
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", r.Url, bytes.NewBufferString(data.Encode()))

	if err != nil {
		return Response{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)

	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	response := Response{}
	response.Status = resp.Status
	response.StatusCode = resp.StatusCode
	response.Headers = resp.Header
	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	response.Body = body
	return response, nil
}
