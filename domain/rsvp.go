package domain

type RSVP string

const (
	Yes         RSVP = "Yes"
	No          RSVP = "No"
	MayBe       RSVP = "MayBe"
	NotAnswered RSVP = "Not Answered" // there was a space in the doc
)

func (r RSVP) IsValid() bool {
	switch r {
	case Yes, No, MayBe, NotAnswered:
		return true
	default:
		return false
	}
}
