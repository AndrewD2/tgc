package session

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const baseURL = "https://www.thegamecrafter.com/api/session"

type TGCSession struct {
	Result struct {
		ID string `json:"id"`
		//WingObjectType string `json:"wing_object_type"`
		UserID    string `json:"user_id"`
		IPAddress string `json:"ip_address"`
	} `json:"result"`
}

func (s *TGCSession) Create(username string, password string, api_key_id string) {

	v := url.Values{"username": {username}, "password":{password}, "api_key_id":{api_key_id}}
	resp, err := http.PostForm(baseURL, v)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		log.Println(err)
	}
}

func (s *TGCSession) Delete() {
	r, err := http.NewRequest(http.MethodDelete, baseURL+s.Result.ID, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
	*s = s.Details()
}

func (s *TGCSession) Details() TGCSession {
	resp, err := http.Get(baseURL + s.Result.ID)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	var d TGCSession

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		log.Println(err)
	}

	return d

}
