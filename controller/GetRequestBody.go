package controller

import (
	"encoding/json"
	"fmt"
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
