package service

import (
	"gethelpnow/domain"
	"gethelpnow/pagination"
	"gethelpnow/repositories/mock"
	"testing"
	"time"
)

var (
	mockRepo           = mock.NewMockRepo()
	svc                = NewService(mockRepo)
	lastAddedMeetingId = ""
)

func TestAdd(t *testing.T) {
	part := []domain.Participant{domain.Participant{
		Name:  "Generic Name",
		Email: "firstname@provider.com",
		Rsvp:  domain.Yes,
	}}
	meeting, err := svc.Add("title", time.Now().Add(5*time.Minute), time.Now().Add(30*time.Minute), part)
	if err != nil {
		t.Fail()
		t.Log(err)
	}
	lastAddedMeetingId = meeting.Id
}

func TestList(t *testing.T) {
	meeting, err := svc.List(time.Now().Add(-1*30*time.Minute), time.Now().Add(30*time.Minute), pagination.Pagination{})
	if err != nil {
		t.Fail()
	}
	if len(meeting) < 1 {
		t.Fail() // bcoz adding one in add test
	}
}

func TestListByEmail(t *testing.T) {
	meeting, err := svc.ListByParticipant("firstname@provider.com", pagination.Pagination{})
	if err != nil {
		t.Fail()
	}
	if len(meeting) < 1 {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	meeting, err := svc.Get(lastAddedMeetingId)
	if err != nil {
		t.Fail()
	}

	if len(meeting.Title) < 1 {
		t.Fail()
	}
}
