package controller

// import (
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// )

// type RequestData struct {
// 	App_Id          string `json:"app_id"`
// 	Avatar_Url      string `json:"avatar_url"`
// 	Candidate_Agent string `json:"candidate_agent"`
// 	Email           string `json:"email"`
// 	Is_New_Session  bool   `json:"is_new_session"`
// 	Is_Resolved     bool   `json:"is_resolved"`
// 	Latest_Service  string `json:"latest_service"`
// 	Name            string `json:"name"`
// 	Room_Id         string `json:"room_id"`
// 	Source          string `json:"source"`
// }

// package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateRequest(w http.ResponseWriter, r *http.Request) {

	url := "https://multichannel.qiscus.com/api/v2/qiscus/initiate_chat"
	method := "POST"

	payload := strings.NewReader(`{
	"app_id": "oni-bgo2lummmhvzqxbt5",
	"user_id": "terlena_saya",
	"name": "terlena",
	"extras": "{\"a\":\"w\"}",
	"user_properties" : {"aji": "tra", "ma": "ta"},
	"reset_extras": true,
	"sdk_user_extras" : {"teststs": "dadadadss"}
}`)

	AdminToken := "QEX6fCm0cwIkPF9F2T22Q1ikXPWVPiqSd0ZJtlrpCOA"
	appId := "1093"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", AdminToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Qiscus-App-Id", appId)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
