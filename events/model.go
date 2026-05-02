package events

import (
	"time"

	"github.com/fahrurben/geteventgo/users"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	OwnerID     uint
	Owner       users.UserModel
	Title       string
	Description string
	Image       string `gorm:"size:1024"`
	StartAt     time.Time
	EndAt       time.Time
	Status      string
}

type Pricing struct {
	gorm.Model
	EventID uint
	Event   Event
	Name    string `gorm:"size:255"`
	Price   float64
	Quota   uint
	OrderNo uint
}
