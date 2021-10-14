package models

import "gorm.io/gorm"

var DB *gorm.DB

type Customer struct {
	gorm.Model
	AgentID int    `json:"agent_id"`
	RoomID  string `json:"room_id"`
	Status  string `json:"status"`
}

type CustomerData struct {
	AppID          string `json:"app_id"`
	CandidateAgent struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		IsAvailable bool   `json:"is_available"`
	} `json:"candidate_agent"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	RoomId       string `json:"room_id"`
	IsNewSession bool   `json:"is_new_session"`
	IsResolved   bool   `json:"is_resolved"`
}

type QiscusResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Queue struct {
	AgentID int `json:"agent_id"`
	RoomID  int `json:"room_id"`
}
