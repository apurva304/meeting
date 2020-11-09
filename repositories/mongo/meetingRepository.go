package repositories

import (
	"context"
	"gethelpnow/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	meetingCollectionName = "Meeting"
)

type meetingRepository struct {
	collection *mongo.Collection
}

func NewMeetingRepository(db *mongo.Database) *meetingRepository {
	return &meetingRepository{
		collection: db.Collection(meetingCollectionName, nil),
	}
}

func (repo *meetingRepository) Add(m domain.Meeting) (err error) {
	_, err = repo.collection.InsertOne(context.TODO(), m, nil)
	if err != nil {
		return
	}

	return
}
func (repo *meetingRepository) Get(id string) (meeting domain.Meeting, err error) {
	err = repo.collection.FindOne(context.TODO(), bson.M{"_id": id}, nil).Decode(&meeting)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = domain.ErrMeetingNotFound
			return
		}
		return
	}
	return
}
func (repo *meetingRepository) List(start time.Time, end time.Time, skip int64, limit int64) (meetings []domain.Meeting, err error) {
	opts := []*options.FindOptions{
		options.Find().SetLimit(skip),
		options.Find().SetSkip(limit),
	}
	ctx := context.TODO()

	curr, err := repo.collection.Find(ctx, bson.M{"start_time": bson.M{"$gte": start, "$lte": end}}, opts...)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = domain.ErrMeetingNotFound
			return
		}
		return
	}

	err = curr.All(ctx, &meetings)
	if err != nil {
		return
	}
	return
}
func (repo *meetingRepository) ListByParticipant(participantEmail string, skip int64, limit int64) (meetings []domain.Meeting, err error) {
	opts := []*options.FindOptions{
		options.Find().SetLimit(skip),
		options.Find().SetSkip(limit),
	}
	ctx := context.TODO()

	curr, err := repo.collection.Find(ctx, bson.M{"participants.email": participantEmail}, opts...)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = domain.ErrMeetingNotFound
			return
		}
		return
	}

	err = curr.All(ctx, &meetings)
	if err != nil {
		return
	}
	return
}
func (repo *meetingRepository) Count(start time.Time, end time.Time, emails []string) (count int64, err error) {
	ctx := context.TODO()

	startM := bson.M{"start_time": bson.M{"$gte": start, "$lte": end}}
	endM := bson.M{"end_time": bson.M{"$gte": start, "$lte": end}}

	filter := bson.M{"participants.email": bson.M{"$in": emails}, "$or": []bson.M{startM, endM}}

	count, err = repo.collection.CountDocuments(ctx, filter, nil)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = domain.ErrMeetingNotFound
			return
		}
		return
	}

	return
}
