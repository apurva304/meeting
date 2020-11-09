package service

import (
	"gethelpnow/cerror"
	"gethelpnow/domain"
	"gethelpnow/pagination"
	"time"
)

var (
	ErrInvalidArgument = cerror.New("Invalid Argument", 400)
	ErrInvalidMeeting  = cerror.New("Invalid Meeting or One of the participants not free", 500)
)

type Service interface {
	Add(title string, start time.Time, end time.Time, participants []domain.Participant) (meeting domain.Meeting, err error)
	List(start time.Time, end time.Time, page pagination.Pagination) (meetings []domain.Meeting, err error)
	ListByParticipant(email string, page pagination.Pagination) (meetings []domain.Meeting, err error)
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

func (svc *service) Add(title string, start time.Time, end time.Time, participants []domain.Participant) (meeting domain.Meeting, err error) {
	if len(title) < 1 || start.IsZero() || end.IsZero() || start.After(end) || start.Before(time.Now()) || len(participants) < 1 {
		err = ErrInvalidArgument
		return
	}

	var emails []string
	for _, part := range participants {
		emails = append(emails, part.Email)
	}

	count, err := svc.meetingRepo.Count(start, end, emails)
	if err != nil {
		return
	}

	if count > 0 {
		err = ErrInvalidMeeting
		return
	}
	meeting = domain.NewMeeting(title, start, end, participants)

	err = svc.meetingRepo.Add(meeting)
	if err != nil {
		return
	}

	return
}

func (svc *service) List(start time.Time, end time.Time, page pagination.Pagination) (meetings []domain.Meeting, err error) {
	if start.IsZero() || end.IsZero() || start.After(end) {
		err = ErrInvalidArgument
		return
	}

	meetings, err = svc.meetingRepo.List(start, end, page.GetSkip(), page.GetLimit())
	if err != nil {
		return
	}
	return
}

func (svc *service) ListByParticipant(email string, page pagination.Pagination) (meetings []domain.Meeting, err error) {
	if len(email) < 1 {
		err = ErrInvalidArgument
		return
	}

	meetings, err = svc.meetingRepo.ListByParticipant(email, page.GetSkip(), page.GetLimit())
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
