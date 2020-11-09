package ctransport

import (
	"encoding/json"
	"gethelpnow/cerror"
	"net/http"
)

type Endpoint func(request interface{}) (response interface{}, err error)
type ResponseEncoder func(w http.ResponseWriter, response interface{}) error
type RequestDecoder func(r *http.Request) (request interface{}, err error)

// encode errors from business-logic
func encodeError(err error, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var status int
	switch err.(type) {
	case *cerror.Error:
		status = err.(*cerror.Error).StatusCode
	default:
		status = 500
	}

	errWarp := cerror.ErrorWrapper{
		Error: err.Error(),
	}

	json.NewEncoder(w).Encode(&errWarp)
	w.WriteHeader(status)
	return nil
}

func NewHandler(endpoint Endpoint, endcode ResponseEncoder, decode RequestDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := decode(r)
		if err != nil {
			encodeError(err, w)
			return
		}

		response, err := endpoint(request)
		if err != nil {
			encodeError(err, w)
			return
		}

		err = endcode(w, response)
		if err != nil {
			encodeError(err, w)
			return
		}
	}
}
