<p align="center">
  <a href="https://gominima.studio">
  <img alt="go-requests" src="./assets/logo.png" />
</a>
</p>

<p align="center">
<a href="https://goreportcard.com/badge/github.com/gominima/go-requests"> <img src="https://goreportcard.com/badge/github.com/gominima/go-requests" /> </a>
<a href="https://img.shields.io/github/go-mod/go-version/gominima/go-requests"> <img src="https://img.shields.io/github/go-mod/go-version/gominima/go-requests" /></a>
<a href="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"> <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" /></a>
<a href="https://discord.gg/gRyCr5APmg"> <img src="https://img.shields.io/discord/916969864512548904" /></a>
<img src="https://img.shields.io/tokei/lines/github/gominima/go-requests" />
<img src="https://img.shields.io/github/languages/code-size/gominima/go-requests" />
</p>

## 🦄 Example

```go
package main

import (
	"fmt"
	"github.com/gominima/go-requests"
)

func main() {
	client := requests.NewRequest("https://random-data-api.com/api/users/random_user", make(map[string]interface{}), make(map[string]string))
	res, err := client.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Body)
	//response would be something like this
	/*
		{
	  "id": 5105,
	  "uid": "58718d5d-456f-45ac-ac0a-fbda86892afd",
	  "password": "RyhWmkV0TH",
	  "first_name": "Bart",
	  "last_name": "Brekke",
	  "username": "bart.brekke",
	  "email": "bart.brekke@email.com",
	  "avatar": "https://robohash.org/quaeratipsumiste.png?size=300x300&set=set1",
	  "gender": "Non-binary",
	  "phone_number": "+1-268 1-580-929-3581",
	  "social_insurance_number": "355857970",
	  "date_of_birth": "1982-09-08",
	  "employment": {
	    "title": "Community-Services Planner",
	    "key_skill": "Organisation"
	  },
	  "address": {
	    "city": "Hyattside",
	    "street_name": "Erdman Loop",
	    "street_address": "93982 Frederick Lake",
	    "zip_code": "86334",
	    "state": "West Virginia",
	    "country": "United States",
	    "coordinates": {
	      "lat": -82.88119017206014,
	      "lng": 150.2430514904285
	    }
	  },
	  "credit_card": {
	    "cc_number": "4198-3576-9309-1264"
	  },
	  "subscription": {
	    "plan": "Bronze",
	    "status": "Idle",
	    "payment_method": "Bitcoins",
	    "term": "Monthly"
	  }
	}
		**/
}


```

## ❓Why Go-requests

Go requests is made for making data fetch wat easier and fun to use, it uses all standard go packages without any other dependency, also is pretty fast and reliable .

<br/>

## ⭐ Contributing

**If you wanna help grow this project or say a thank you!**

1. Give minima a [GitHub star](https://github.com/gominima/go-requests/stargazers)
2. Fork requests and Contribute
4. Join our [Discord](https://discord.gg/gRyCr5APmg) community
