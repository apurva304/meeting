package service

import (
	"encoding/json"
	ctransport "gethelpnow/transport"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func MakeHandler(svc Service) http.Handler {
	addHandler := ctransport.NewHandler(
		makeAddEndpoint(svc),
		encodeAddResponse,
		decodeAddRequest,
	)

	listHandler := ctransport.NewHandler(
		makeListEndpoint(svc),
		encodeListResponse,
		decodeListRequest,
	)

	listByParticipantHandler := ctransport.NewHandler(
		makeListByParticipantEndpoint(svc),
		encodeListByParticipantResponse,
		decodeListByParticipantRequest,
	)

	getHandler := ctransport.NewHandler(
		makeGetEndpoint(svc),
		encodeGetResponse,
		decodeGetRequest,
	)

	r := mux.NewRouter()
	r.Handle("/meetings", addHandler).Methods(http.MethodPost)
	r.Handle("/meetings", listHandler).Methods(http.MethodGet)
	r.Handle("/meeting/{id}", getHandler).Methods(http.MethodGet)
	r.Handle("/meeting", listByParticipantHandler).Methods(http.MethodGet)

	return r
}
func decodeAddRequest(r *http.Request) (request interface{}, err error) {
	var req addRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, ErrBadRequest
	}
	return
}

func encodeAddResponse(w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(addResponse); ok && e.Err != nil {
		return e.Err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(&response)
}

func decodeListRequest(r *http.Request) (request interface{}, err error) {
	var req listRequest
	err = schema.NewDecoder().Decode(&req, r.URL.Query())
	if err != nil {
		return nil, ErrBadRequest
	}
	return
}

func encodeListResponse(w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(listResponse); ok && e.Err != nil {
		return e.Err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(&response)
}

func decodeListByParticipantRequest(r *http.Request) (request interface{}, err error) {
	var req listByParticipantRequest
	err = schema.NewDecoder().Decode(&req, r.URL.Query())
	if err != nil {
		return nil, ErrBadRequest
	}
	return
}

func encodeListByParticipantResponse(w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(listByParticipantResponse); ok && e.Err != nil {
		return e.Err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(&response)
}

func decodeGetRequest(r *http.Request) (request interface{}, err error) {
	var req getRequest
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || len(id) < 1 {
		return nil, ErrBadRequest
	}

	req.Id = id
	return req, nil
}

func encodeGetResponse(w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(getResponse); ok && e.Err != nil {
		return e.Err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(&response)
}
