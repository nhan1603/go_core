package httpserv

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Success is the response format when http handler succeeds
type Success struct {
	Message string `json:"message,omitempty"`
}

type Response struct {
	Success bool `json:"success"`
}

type FriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type UpdateReceiveResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}

// RespondJSON handles conversion of the requested result to JSON format
func RespondJSON(ctx context.Context, w http.ResponseWriter, obj interface{}) {
	RespondJSONWithHeaders(ctx, w, obj, nil)
}

// RespondJSONWithHeaders handles conversion of the requested result to JSON format
func RespondJSONWithHeaders(ctx context.Context, w http.ResponseWriter, obj interface{}, headers map[string]string) {
	// Set HTTP headers
	w.Header().Set("Content-Type", "application/json")
	for key, value := range headers {
		w.Header().Set(key, value)
	}

	status := http.StatusOK
	var respBytes []byte
	var err error

	switch parsed := obj.(type) {
	case *Error:
		if parsed.Status >= http.StatusInternalServerError && parsed.Status != http.StatusServiceUnavailable {
			parsed.Desc = DefaultErrorDesc
		}
		status = parsed.Status
		if status == 0 {
			status = http.StatusInternalServerError
		}
		respBytes, err = json.Marshal(parsed)
	case error:
		fmt.Println(obj)
		status = http.StatusInternalServerError
		respBytes, err = json.Marshal(&Error{
			Status: http.StatusInternalServerError,
			Code:   DefaultErrorCode,
			Desc:   obj.(error).Error(),
		})
	default:
		respBytes, err = json.Marshal(obj)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(status)
	w.Write(respBytes)
}
