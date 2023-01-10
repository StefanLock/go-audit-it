package models

import (
	"fmt"
	"time"
)

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.Now().Format(time.RFC822)))
	return []byte(stamp), nil
}

type System struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Host        string   `json:"host"`
	Status      string   `json:"status"`
	LastCheckIn JSONTime `json:"Last_Check_In"`
}
