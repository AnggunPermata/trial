package controller

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestData struct {
	AppID          string      `json:"app_id"`
	AvatarURL      string      `json:"avatar_url"`
	CandidateAgent interface{} `json:"candidate_agent"`
	Email          string      `json:"email"`
	Extras         string      `json:"extras"`
	IsNewSession   bool        `json:"is_new_session"`
	IsResolved     bool        `json:"is_resolved"`
	LatestService  interface{} `json:"latest_service"`
	Name           string      `json:"name"`
	RoomID         string      `json:"room_id"`
	Source         string      `json:"source"`
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var req RequestData

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct...
	fmt.Fprintf(w, "%+v", req)
}
*/

import (
	"bytes"
	"fmt"
	"github.com/anggunpermata/custom-webhook/config"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func InitiateChat(w http.ResponseWriter, r *http.Request) {

	url := "https://multichannel.qiscus.com/api/v1/qiscus/initiate_chat"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("app_id", "")
	_ = writer.WriteField("user_id", "")
	_ = writer.WriteField("name", "")
	_ = writer.WriteField("notes", "")
	_ = writer.WriteField("user_properties", "")
	_ = writer.WriteField("extras", "")
	_ = writer.WriteField("timezone_offset", "")
	_ = writer.WriteField("allocate_agent", "")
	_ = writer.WriteField("room_badge", "")
	_ = writer.WriteField("origin", "")
	_ = writer.WriteField("reset_extras", "")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	AdminToken := config.GoDotEnvVariable("AdminToken")
	AppCode := config.GoDotEnvVariable("AppId")

	req.Header.Add("Authorization", AdminToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Qiscus-App-Id", AppCode)

	req.Header.Set("Content-Type", writer.FormDataContentType())
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
