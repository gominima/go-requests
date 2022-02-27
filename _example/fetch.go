package main

import (
	"fmt"

	"github.com/gominima/go-requests"
)

func main() {
     client := goquests.New()
     resp, err := client.Get(goquests.Request{
	     URL: "https://random-data-api.com/api/users/random_user",
	     Data: make(map[string]interface{}),
	     Headers: make(map[string]string),
     })
     if err != nil {
	     panic(err)
     }
     fmt.Print(resp.Body)
}
