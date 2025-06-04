package notification

import (
	"fmt"
	"graphql-hasura-demo/internal/dto/hasura"
)

type Service struct {
	repository *repository
}

func (s *Service) NotifyTaskUpdated(payload hasura.UpdatedEventRequest[Grade]) (*hasura.UpdatedEventResponse, error) {
	fmt.Printf("Payload: %v\n", payload)
	newScore := payload.Event.Data.New.Score
	oldScore := payload.Event.Data.Old.Score
	studentId := payload.Event.Data.Old.StudentId

	if newScore == oldScore {
		return &hasura.UpdatedEventResponse{
			Status:  "success",
			Message: "Score is not updated",
		}, nil
	}

	message := fmt.Sprintf("Score's studentId: %d --- %.2f changed to %.2f", studentId, newScore, oldScore)
	notification := NewNotification(message, payload.Event.Data.New.StudentId)
	_, err := s.repository.save(notification)
	if err != nil {
		return &hasura.UpdatedEventResponse{
			Status:  "error",
			Message: "Failed to create notification",
		}, nil
	}

	return &hasura.UpdatedEventResponse{
		Status:  "success",
		Message: "Notification created",
	}, nil
}

func NewService() *Service {
	repository := NewRepository()
	service := Service{repository}

	return &service
}
