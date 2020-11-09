package mock

import (
	"gethelpnow/domain"
	"sync"
	"time"
)

type mockRepo struct {
	meetings []domain.Meeting
	sync.Mutex
}

func (repo *mockRepo) Add(m domain.Meeting) (err error) {
	repo.Mutex.Lock()
	repo.meetings = append(repo.meetings, m)
	repo.Mutex.Unlock()
	return
}
func (repo *mockRepo) Get(id string) (meeting domain.Meeting, err error) {
	repo.Mutex.Lock()
	for _, m := range repo.meetings {
		if id == m.Id {
			return m, nil
		}
	}
	err = domain.ErrMeetingNotFound
	repo.Mutex.Unlock()
	return
}
func (repo *mockRepo) List(start time.Time, end time.Time, skip int64, limit int64) (meetings []domain.Meeting, err error) {
	repo.Mutex.Lock()
	for _, m := range repo.meetings {
		if m.StartTime.After(start) && m.StartTime.Before(end) {
			meetings = append(meetings, m)
		}
	}
	repo.Mutex.Unlock()
	return
}
func (repo *mockRepo) ListByParticipant(participantEmail string, skip int64, limit int64) (meetings []domain.Meeting, err error) {
	repo.Mutex.Lock()
	for _, m := range repo.meetings {
		for _, p := range m.Participants {
			if p.Email == participantEmail {
				meetings = append(meetings, m)
			}
		}
	}
	repo.Mutex.Unlock()
	return
}
func (repo *mockRepo) Count(start time.Time, end time.Time, emails []string) (count int64, err error) {
	repo.Mutex.Lock()
	meetings, err := repo.List(start, end, 0, 0)
	if err != nil {
		return
	}

	partEmailMap := make(map[string]struct{})
	for _, m := range meetings {
		for _, p := range m.Participants {
			partEmailMap[p.Email] = struct{}{}
		}
	}

	for k, _ := range partEmailMap {
		if contains(emails, k) {
			count += 1
		}
	}
	repo.Mutex.Unlock()
	return
}

func contains(s []string, str string) bool {
	for _, a := range s {
		if a == str {
			return true
		}
	}
	return false
}
