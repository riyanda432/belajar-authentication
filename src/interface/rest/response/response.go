package response

import (
	"encoding/json"
	"math"
	"net/http"
	error_helper "github.com/riyanda432/belajar-authentication/src/helper"

)

// HealthCheckMessage displays information about the service
type HealthCheckMessage struct {
	ServiceName string `json:"serviceName"`
	Version     string `json:"version"`
	CommitId    string `json:"commitId"`
	UpdatedAt   string `json:"updatedAt"`
	Status      string `json:"status"`
}

// Meta consist of pagination details
type Meta struct {
	Page      int     `json:"page,omitempty"`
	PerPage   int     `json:"perPage,omitempty"`
	TotalPage float64 `json:"totalPage,omitempty"`
}

// ResponseMessage consist of payload details
// Data -> Payload
// Meta -> Pagination etc
type ResponseMessage struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ErrorMessage struct {
	Message      string `json:"message"`
	ErrorMessage string `json:"errorMessage"`
	Type         string `json:"type"`
	Code         int    `json:"code"`
}

type ValidationErrorMessage struct {
	Message      string      `json:"message"`
	ErrorMessage interface{} `json:"errorMessage"`
	Type         string      `json:"type"`
	Code         int         `json:"code"`
}

type IResponseClient interface {
	JSON(w http.ResponseWriter, message string, data interface{}, meta *Meta) error
	BuildMeta(page int, perPage int, count int64) *Meta
	HttpError(w http.ResponseWriter, err error) error
}

type responseClient struct{}

func NewResponseClient() IResponseClient {
	return &responseClient{}
}

func (r *responseClient) JSON(
	w http.ResponseWriter,
	message string,
	data interface{},
	meta *Meta,
) error {
	response := ResponseMessage{
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	return nil
}

func (r *responseClient) HttpError(
	w http.ResponseWriter,
	err error,
) error {
	var respError error_helper.HttpError

	if cerr, ok := err.(*error_helper.CommonError); ok {
		respError = cerr.ToHttpError()
	} else {
		respError = error_helper.NewCommonError(error_helper.UNKNOWN_ERROR, err).ToHttpError()
	}

	resp, _ := json.Marshal(respError)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respError.GetHttpStatus())
	w.Write(resp)
	return nil
}

func (r *responseClient) BuildMeta(page int, perPage int, count int64) *Meta {
	return &Meta{
		Page:      page,
		PerPage:   perPage,
		TotalPage: math.Ceil(float64(count) / float64(perPage)),
	}
}
