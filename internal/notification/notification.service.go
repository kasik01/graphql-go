package notification

import (
	"fmt"
)

type Service struct {
	repository *repository
}

func (s *Service) NotifyTaskUpdated(payload TaskUpdatedEvent) (*TaskUpdatedEventResponse, error) {
	if payload.Event.Data.New.IsCompleted == payload.Event.Data.Old.IsCompleted {
		return &TaskUpdatedEventResponse{
			Status:  "success",
			Message: "Status is not updated",
		}, nil
	}

	status := "incomplete"
	if payload.Event.Data.New.IsCompleted {
		status = "complete"
	}
	message := fmt.Sprintf("Task '%s' status changed to: %s", payload.Event.Data.New.Description, status)
	notification := NewNotification(message, payload.Event.Data.New.UserId)
	_, err := s.repository.save(notification)
	if err != nil {
		return &TaskUpdatedEventResponse{
			Status:  "error",
			Message: "Failed to create notification",
		}, nil
	}

	return &TaskUpdatedEventResponse{
		Status:  "success",
		Message: "Notification created",
	}, nil
}

func NewService() *Service {
	repository := NewRepository()
	service := Service{repository}

	return &service
}
