package domain

import (
	"gethelpnow/cerror"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrMeetingNotFound = cerror.New("Meeting Not Found", 404)
)

type MeetingRepository interface {
	Add(m Meeting) (err error)
	Get(id string) (meeting Meeting, err error)
	List(start time.Time, end time.Time) (meetings []Meeting, err error)
	ListByParticipant(participantEmail string) (meetings []Meeting, err error)
}

type Meeting struct {
	Id                string        `json:"id" bson:"_id"`
	Title             string        `json:"title" bson:"title"`
	Participants      []Participant `json:"participants" bson:"participants"`
	StartTime         time.Time     `json:"start_time" bson:"start_time"`
	EndTime           time.Time     `json:"end_time" bson:"end_time"`
	CreationTimestamp time.Time     `json:"creation_timestamp" bson:"creation_timestamp"`
}

func NewMeeting(title string, start time.Time, end time.Time, part []Participant) Meeting {
	return Meeting{
		Id:                primitive.NewObjectID().Hex(),
		Title:             title,
		Participants:      part,
		StartTime:         start,
		EndTime:           end,
		CreationTimestamp: time.Now(),
	}
}
