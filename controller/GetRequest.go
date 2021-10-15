package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/anggunpermata/custom-webhook/config"
	"github.com/anggunpermata/custom-webhook/models"
)

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
)
*/

/*
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

func AssignAgent(cus models.CustomerData, agentId string) (*models.QiscusResponse, error) {
	data := url.Values{}
	data.Set("room_id", cus.RoomID)
	data.Set("agent_id", agentId)
	data.Set("app_id", cus.AppID)
	data.Set("user_id", strconv.Itoa(cus.LatestService.UserID))
	// data.Set("name", cus.Name)

	var w http.ResponseWriter
	var r *http.Request
	resp, err := InitiateChat(w, r, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	newResp := new(models.QiscusResponse)
	if err2 := json.Unmarshal(resp, &newResp.Data); err2 != nil {
		return nil, err
	} else {
		newResp.Message = "Assign Agent Success"
	}
	return newResp, nil
}

func InitiateChat(w http.ResponseWriter, r *http.Request, payload *strings.Reader) ([]byte, error) {

	uri := "https://multichannel.qiscus.com/api/v2/qiscus/initiate_chat"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, uri, payload)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error from structuring requests")
	}

	AdminToken := config.GoDotEnvVariable("AdminToken")
	// AppCode := config.GoDotEnvVariable("AppId")
	// SecretKey := config.GoDotEnvVariable("SecretKey")

	req.Header.Add("Authorization", AdminToken)
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Qiscus-App-Id", AppCode)
	// req.Header.Add("Qiscus-Secret-Key", SecretKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error from getting response")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error from read response body")
	}
	// result := string(body)
	return body, nil
}
