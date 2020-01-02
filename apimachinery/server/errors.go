package server

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloustone/pandas/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

func serverError(err error) middleware.Responder {
	if err != nil {
		logrus.Errorf(err.Error())
		switch status.Code(err) {
		case codes.InvalidArgument:
			return badRequestError(err)
		case codes.NotFound:
			return notFoundError(err)
		default:
			return internalServerError(err)
		}
	}
	return nil
}

func badRequestError(err error) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		w.WriteHeader(http.StatusBadRequest)
		payload, _ := json.Marshal(&models.Error{
			Description: err.Error(),
		})
		w.Write(payload)
	})
}

func notFoundError(err error) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		w.WriteHeader(http.StatusNotFound)
		payload, _ := json.Marshal(&models.Error{
			Description: err.Error(),
		})
		w.Write(payload)
	})
}

func internalServerError(err error) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		w.WriteHeader(http.StatusInternalServerError)
		payload, _ := json.Marshal(&models.Error{
			Description: err.Error(),
		})
		w.Write(payload)
	})
}
