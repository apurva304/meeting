package service

import (
	"gethelpnow/cerror"
	"gethelpnow/domain"
	"time"
)

var (
	ErrInvalidArgument = cerror.New("Invalid Argument", 400)
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
	if start.IsZero() || end.IsZero() || start.After(end) || start.Before(time.Now()) {
		err = ErrInvalidArgument
		return
	}

	meetings, err = svc.meetingRepo.List(start, end)
	if err != nil {
		return
	}
	return
}
func (svc *service) ListByParticipant(email string) (meetings []domain.Meeting, err error) {
	if len(email) < 1 {
		err = ErrInvalidArgument
		return
	}

	meetings, err = svc.meetingRepo.ListByParticipant(email)
	if err != nil {
		return
	}
	return
}
func (svc *service) Get(id string) (meeting domain.Meeting, err error) {
	if len(id) < 1 {
		err = ErrInvalidArgument
		return
	}
	meeting, err = svc.meetingRepo.Get(id)
	if err != nil {
		return
	}
	return
}
