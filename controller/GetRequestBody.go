package controller

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RequestData struct {
	App_Id          string `json:"app_id"`
	Avatar_Url      string `json:"avatar_url"`
	Candidate_Agent string `json:"candidate_agent"`
	Email           string `json:"email"`
	Is_New_Session  bool   `json:"is_new_session"`
	Is_Resolved     bool   `json:"is_resolved"`
	Latest_Service  string `json:"latest_service"`
	Name            string `json:"name"`
	Room_Id         string `json:"room_id"`
	Source          string `json:"source"`
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e RequestData
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Success", http.StatusOK)
	return
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}