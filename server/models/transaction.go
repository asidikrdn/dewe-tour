package models

import "time"

type Transaction struct {
	ID          string    `json:"id" gorm:"type: varchar(255);PRIMARY_KEY"`
	CounterQty  int       `json:"counterQty" gorm:"type: int"`
	Total       int       `json:"total" gorm:"type: int"`
	Status      string    `json:"status" gorm:"type: varchar(255)"`
	BookingDate time.Time `json:"bookingDate"`
	Token       string    `json:"token" gorm:"type: varchar(255)"`
	TripID      int       `json:"tripId" gorm:"type: int"`
	Trip        TripResponse
	UserID      int `json:"id_user" gorm:"type: int"`
	User        UserResponse
}

type TransactionResponse struct {
	ID         string `json:"id"`
	CounterQty int    `json:"counterQty" gorm:"type: int"`
	Total      int    `json:"total"`
	Status     string `json:"status" gorm:"type: varchar(255)"`
	TripID     int    `json:"tripId" gorm:"type: int"`
	UserID     int    `json:"id_user" gorm:"type: int"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
