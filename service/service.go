package service

import (
	"gethelpnow/domain"
	"time"
)

type Service interface {
	Add(title string, start time.Time, end time.Time, participants []domain.Participant) (err error)
	List(start time.Time, end time.Time) (meetings []domain.Meeting, err error)
	ListByParticipant(email string) (meetings []domain.Meeting, err error)
	Get(id string) (meeting domain.Meeting, err error)
}

type service struct {
	meetingRepo domain.MeetingRepository
}

func NewService(meetingRepo domain.MeetingRepository) *service {
	return &service{
		meetingRepo: meetingRepo,
	}
}

func (svc *service) Add(title string, start time.Time, end time.Time, participants []domain.Participant) (err error) {
	return
}
func (svc *service) List(start time.Time, end time.Time) (meetings []domain.Meeting, err error) {
	return
}
func (svc *service) ListByParticipant(email string) (meetings []domain.Meeting, err error) {
	return
}
func (svc *service) Get(id string) (meeting domain.Meeting, err error) {
	return
}
