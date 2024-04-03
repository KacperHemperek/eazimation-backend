package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

func HttpHandler(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var apiErr *Error
		err := h(w, r)
		defer func(now time.Time) {
			t := time.Since(now)
			if apiErr != nil {
				var apiErrMess string
				if apiErr.Cause != nil {
					apiErrMess = apiErr.Cause.Error()
				} else {
					apiErrMess = "none"
				}
				slog.Error(
					"Error in handler",
					"method", r.Method,
					"url", r.URL.String(),
					"cause", apiErrMess,
					"code", apiErr.Code,
					"time", t,
				)
				return
			}
			if err != nil {
				slog.Error(
					"Error in handler",
					"method", r.Method,
					"url", r.URL.String(),
					"cause", err,
					"code", 500,
					"time", t,
				)
				return
			}
			slog.Info("Request", "method", r.Method, "url", r.URL.String(), "time", t)
		}(time.Now())

		if err != nil {
			if errors.As(err, &apiErr) {
				_ = WriteJSON(w, apiErr.Code, apiErr)
			} else {
				_ = WriteJSON(w, http.StatusInternalServerError, &Error{
					Message: "Internal Server Error",
					Code:    http.StatusInternalServerError,
				})
			}
		}
	}
}

func WriteJSON(w http.ResponseWriter, s int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	if v != nil {
		return json.NewEncoder(w).Encode(v)
	}
	return nil
}

type HandlerFunc = func(w http.ResponseWriter, r *http.Request) error

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Cause   error  `json:"-"`
}

func (e *Error) Error() string {
	return e.Message
}
