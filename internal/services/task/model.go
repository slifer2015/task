package task

import "encoding/json"

type CreateTaskInput struct {
	ID          string
	Method      string
	URL         string
	Headers     map[string]string
	Body        json.RawMessage
	QueryParams map[string]string
}

type CreateTaskResponse struct {
	ID string
}

type status string

const (
	statusNew        = "new"
	statusInProgress = "in_progress"
	statusError      = "error"
	statusDone       = "done"
)

type JobResult struct {
	Status         status              `json:"status"`
	HTTPStatusCode int                 `json:"httpStatusCode"`
	Headers        map[string][]string `json:"headers"`
	Length         int                 `json:"length"`
	Response       string              `json:"response"`
}
