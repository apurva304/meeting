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
	client *mongo.Client
	dbName string
}

func NewMeetingRepository(client *mongo.Client, dbName string) *meetingRepository {
	return &meetingRepository{
		client: client,
		dbName: dbName,
	}
}

func (repo *meetingRepository) Add(m domain.Meeting) (err error) {
	coll := repo.client.Database(repo.dbName).Collection(meetingCollectionName)
	_, err = coll.InsertOne(context.TODO(), m)
	if err != nil {
		return
	}

	return
}
func (repo *meetingRepository) Get(id string) (meeting domain.Meeting, err error) {
	coll := repo.client.Database(repo.dbName).Collection(meetingCollectionName)
	err = coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&meeting)
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

	coll := repo.client.Database(repo.dbName).Collection(meetingCollectionName)
	curr, err := coll.Find(ctx, bson.M{"start_time": bson.M{"$gte": start, "$lte": end}}, opts...)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = nil
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

	coll := repo.client.Database(repo.dbName).Collection(meetingCollectionName)
	curr, err := coll.Find(ctx, bson.M{"participants.email": participantEmail}, opts...)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = nil
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

	coll := repo.client.Database(repo.dbName).Collection(meetingCollectionName)

	count, err = coll.CountDocuments(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = nil
			return
		}
		return
	}

	return
}
