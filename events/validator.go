package events

import "time"

type EventValidator struct {
	Title       string    `json:"title" binding:"required,min=3"`
	Description string    `json:"description" binding:"required,min=3"`
	Image       string    `json:"image"`
	StartAt     time.Time `json:"start_at" binding:"required" time_format:"2006-01-02 15:04:05"`
	EndAt       time.Time `json:"end_at" binding:"required" time_format:"2006-01-02 15:04:05"`
	Status      string    `json:"status"`
}

func (validator *EventValidator) toModel() Event {
	eventModel := Event{}
	eventModel.Title = validator.Title
	eventModel.Description = validator.Description
	eventModel.Image = validator.Image
	eventModel.StartAt = validator.StartAt
	eventModel.EndAt = validator.EndAt
	eventModel.Status = validator.Status

	return eventModel
}
