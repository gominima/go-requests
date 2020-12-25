package requests

//check
import (
	"bytes"
	"encoding/json"
	"net/http"
	//"net/url"
)

type Request struct {
	Url     string
	Data    map[string]interface{}
	Headers map[string]string
}

type Response struct {
	Status  string
	Headers interface{}
	Body    map[string]interface{}
}

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
