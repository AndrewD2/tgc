package tgc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const baseURL = "https://www.thegamecrafter.com/api/session"

type Session struct {
	Result struct {
		ID string `json:"id"`
		//WingObjectType string `json:"wing_object_type"`
		UserID    string `json:"user_id"`
		IPAddress string `json:"ip_address"`
	} `json:"result"`
}

func (s *Session) Create(username string, password string, api_key_id string) error {

	v := url.Values{"username": {username}, "password": {password}, "api_key_id": {api_key_id}}
	resp, err := http.PostForm(baseURL, v)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var e Error
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			log.Println(err)
		}

		err := fmt.Errorf("TGC Session: %d - %s", e.Error.Code, e.Error.Message)
		return err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		log.Println(err)
	}

	return nil
}

func (s *Session) Delete() {
	_, err := http.NewRequest(http.MethodDelete, baseURL+s.Result.ID, nil)
	if err != nil {
		log.Println(err)
	}

	*s = s.Details()
}

func (s *Session) Details() (d Session) {
	resp, err := http.Get(baseURL + s.Result.ID)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		log.Println(err)
	}

	return

}
