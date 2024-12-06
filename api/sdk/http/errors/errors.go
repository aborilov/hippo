package errors

import (
	"encoding/json"
	"net/http"

	"github.com/aborilov/hippo/foundation/logger"
)

// API error codes.
const (
	CodeInternalError  = "INTERNAL_ERROR"
	CodeInvalidRequest = "INVALID_REQUEST"
	CodeNotFound       = "NOT_FOUND"
)

var (
	// JSON - base error with "application/json" content type
	JSON = Error().SetContentType("application/json")

	// NotFoundError - base error with http status 404
	NotFoundError = JSON.SetCode(CodeNotFound).SetHTTPCode(http.StatusNotFound)

	// BadRequestError - base error with http status 400
	BadRequestError = JSON.SetCode(CodeInvalidRequest).SetHTTPCode(http.StatusBadRequest)

	// InternalError - base error with http status 500
	InternalError = JSON.SetCode(CodeInternalError).SetHTTPCode(http.StatusInternalServerError)
)

// Generic error response.
type ErrorResponse struct {
	Code       string `json:"code"`
	DetailCode string `json:"detail_code"`
	Message    string `json:"message"`
}

// NotFound - write NotFoundError error with message to response
func NotFound(w http.ResponseWriter, msg string) {
	NotFoundError.SetMessage(msg).Write(w)
}

// BadRequest - write BadRequestError error with message to response
func BadRequest(w http.ResponseWriter, msg string) {
	BadRequestError.SetMessage(msg).Write(w)
}

// Internal - write InternalError error with message to response and log err if it's not nil
func Internal(w http.ResponseWriter, msg string, err error) {
	if err != nil {
		// TODO: use default logger
		log, _ := logger.NewLogger()
		log.Error(err, msg)
	}
	InternalError.SetMessage(msg).Write(w)
}

type apiError struct {
	contentType string
	httpCode    int
	code        string
	message     string
	detailCode  string
}

// APIError interface
type APIError interface {
	SetContentType(string) APIError
	SetHTTPCode(int) APIError
	SetCode(string) APIError
	SetDetailCode(string) APIError
	SetMessage(string) APIError
	Write(w http.ResponseWriter)
}

// Error create new APIError
func Error() APIError {
	return apiError{}
}

func (e apiError) SetContentType(ct string) APIError {
	e.contentType = ct
	return e
}

func (e apiError) SetHTTPCode(code int) APIError {
	e.httpCode = code
	return e
}

func (e apiError) SetCode(code string) APIError {
	e.code = code
	return e
}

func (e apiError) SetDetailCode(code string) APIError {
	e.detailCode = code
	return e
}

func (e apiError) SetMessage(msg string) APIError {
	e.message = msg
	return e
}

func (e apiError) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", e.contentType)
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.WriteHeader(e.httpCode)
	err := json.NewEncoder(w).Encode(
		ErrorResponse{
			Message:    e.message,
			Code:       e.code,
			DetailCode: e.detailCode,
		})
	if err != nil {
		http.Error(w, `{"code": "internal_error", "message": "Unable to write error response"}`,
			http.StatusInternalServerError)
	}
}
