package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/anggunpermata/custom-webhook/config"
	"github.com/anggunpermata/custom-webhook/models"
)

func AssignAgent(cus models.CustomerData, agentId string) (*models.QiscusResponse, error) {
	data := url.Values{}
	data.Set("room_id", cus.RoomId)
	data.Set("agent_id", agentId)
	data.Set("app_id", cus.AppID)
	// data.Set("user_id", strconv.Itoa(cus.LatestService.UserID))
	data.Set("name", cus.Name)

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

	uri := "https://multichannel.qiscus.com/api/v1/admin/service/assign_agent"
	method := "POST"

	// payloads := &bytes.Buffer{}
	// writer := multipart.NewWriter(payloads)
	// _ = writer.WriteField("app_id", "")
	// _ = writer.WriteField("user_id", "")
	// _ = writer.WriteField("name", "")
	// _ = writer.WriteField("notes", "")
	// _ = writer.WriteField("user_properties", "")
	// _ = writer.WriteField("extras", "")
	// _ = writer.WriteField("timezone_offset", "")
	// _ = writer.WriteField("allocate_agent", "")
	// _ = writer.WriteField("room_badge", "")
	// _ = writer.WriteField("origin", "")
	// _ = writer.WriteField("reset_extras", "")
	// err := writer.Close()
	// if err != nil {
	// 	// fmt.Println(err)
	// 	return nil, err
	// }

	// client := &http.Client{}
	// req, err := http.NewRequest(method, uri, payloads)
	// if err != nil {
	// 	return nil, err
	// }
	// AdminToken := config.GoDotEnvVariable("AdminToken")
	// AppCode := config.GoDotEnvVariable("AppId")
	// // SecretKey := config.GoDotEnvVariable("SecretKey")

	// req.Header.Add("Authorization", AdminToken)
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("Qiscus-App-Id", AppCode)

	// req.Header.Set("Content-Type", writer.FormDataContentType())

	// res, err := client.Do(req)
	// if err != nil {
	// 	return nil, fmt.Errorf("error from getting response")
	// }
	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("error from read response body")
	// }
	// // result := string(body)
	// return body, nil

	fmt.Println(payload)
	//---------------------
	client := &http.Client{}
	req, err := http.NewRequest(method, uri, payload)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error from structuring requests")
	}

	AdminToken := config.GoDotEnvVariable("AdminToken")
	AppCode := config.GoDotEnvVariable("AppId")
	//SecretKey := config.GoDotEnvVariable("SecretKey")

	req.Header.Add("Authorization", AdminToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Qiscus-App-Id", AppCode)
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
	//-------------------------------
}
