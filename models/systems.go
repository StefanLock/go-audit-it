package models

import "time"

type System struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Host        string    `json:"host"`
	Status      bool      `json:"status"`
	LastCheckIn time.Time `json:"Last_Check_In"`
}
