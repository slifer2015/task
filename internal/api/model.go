package api

import (
	"encoding/json"
)

type createTaskRequest struct {
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	QueryParams map[string]string `json:"query_params"`
	Body        json.RawMessage   `json:"body"`
}

type createTaskResponse struct {
	ID string `json:"id"`
}
