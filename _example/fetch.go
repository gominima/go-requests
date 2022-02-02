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
