package service

import (
	"gethelpnow/cerror"
	"gethelpnow/domain"
	ctransport "gethelpnow/transport"
	"time"
)

var (
	ErrBadRequest = cerror.New("Bad Request", 400)
)

type addRequest struct {
	Title        string               `json:"title"`
	Start        time.Time            `json:"start"`
	End          time.Time            `json:"end"`
	Participants []domain.Participant `json:"participants"`
}

type addResponse struct {
	Err error `json:"err,omitempty"`
}

func makeAddEndpoint(svc Service) ctransport.Endpoint {
	return func(request interface{}) (response interface{}, err error) {
		req, ok := request.(addRequest)
		if !ok {
			return nil, ErrBadRequest
		}
		err = svc.Add(req.Title, req.Start, req.End, req.Participants)
		return addResponse{Err: err}, nil
	}
}

type listRequest struct {
	Start time.Time `schema:"start" url:"start"`
	End   time.Time `schema:"end" url:"end"`
}

type listResponse struct {
	Meetings []domain.Meeting `json:"meetings"`
	Err      error            `json:"err,omitempty"`
}

func makeListEndpoint(svc Service) ctransport.Endpoint {
	return func(request interface{}) (response interface{}, err error) {
		req, ok := request.(listRequest)
		if !ok {
			return nil, ErrBadRequest
		}
		meetings, err := svc.List(req.Start, req.End)
		return listResponse{Meetings: meetings, Err: err}, nil
	}
}

type listByParticipantRequest struct {
	Email string `schema:"email" url:"email"`
}
type listByParticipantResponse struct {
	Meetings []domain.Meeting `json:"meetings"`
	Err      error            `json:"err,omitempty"`
}

func makeListByParticipantEndpoint(svc Service) ctransport.Endpoint {
	return func(request interface{}) (response interface{}, err error) {
		req, ok := request.(listByParticipantRequest)
		if !ok {
			return nil, ErrBadRequest
		}
		meetings, err := svc.ListByParticipant(req.Email)
		return listByParticipantResponse{Meetings: meetings, Err: err}, nil
	}
}

type getRequest struct {
	Id string `schema:"id" url:"id"`
}

type getResponse struct {
	Meeting domain.Meeting `json:"meeting"`
	Err     error          `json:"err,omitempty"`
}

func makeGetEndpoint(svc Service) ctransport.Endpoint {
	return func(request interface{}) (response interface{}, err error) {
		req, ok := request.(getRequest)
		if !ok {
			return nil, ErrBadRequest
		}
		meeting, err := svc.Get(req.Id)
		return getResponse{Meeting: meeting, Err: err}, nil
	}
}
