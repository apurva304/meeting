package domain

type Participant struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Rsvp  RSVP   `json:"rsvp" bson:"rsvp"`
}
