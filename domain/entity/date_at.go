package entity

import (
	"encoding/json"
	"time"
)

// DateAt 时间
type DateAt struct {
	Date time.Time
}

func NewDateAt() *DateAt {
	return &DateAt{
		Date: time.Now(),
	}
}

func (a DateAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Date   time.Time `json:"date"`
		Format string    `json:"format"`
		Unix   int64     `json:"unix"`
	}{
		Date:   a.Date,
		Format: a.DefaultFormat(),
		Unix:   a.Date.Unix(),
	})
}

func (a *DateAt) DefaultFormat() string {
	return a.Date.Format("2006-01-02 15:04:05")
}
