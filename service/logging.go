package service

import (
	"gethelpnow/domain"
	"log"
	"time"
)

type loggingService struct {
	Service
	logger *log.Logger
}

func NewLoggingService(svc Service, logger *log.Logger) *loggingService {
	return &loggingService{
		Service: svc,
		logger:  logger,
	}
}
func (svc *loggingService) Add(title string, start time.Time, end time.Time, participants []domain.Participant) (err error) {
	defer func(begin time.Time) {
		svc.logger.Println(
			"method", "Add",
			"title", title,
			"start", start,
			"end", end,
			"participants", participants,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return svc.Service.Add(title, start, end, participants)
}
func (svc *loggingService) List(start time.Time, end time.Time) (meetings []domain.Meeting, err error) {
	defer func(begin time.Time) {
		svc.logger.Println(
			"method", "List",
			"start", start,
			"end", end,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return svc.Service.List(start, end)
}
func (svc *loggingService) ListByParticipant(email string) (meetings []domain.Meeting, err error) {
	defer func(begin time.Time) {
		svc.logger.Println(
			"method", "ListByParticipant",
			"email", email,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return svc.Service.ListByParticipant(email)
}
func (svc *loggingService) Get(id string) (meeting domain.Meeting, err error) {
	defer func(begin time.Time) {
		svc.logger.Println(
			"method", "Get",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return svc.Service.Get(id)
}
