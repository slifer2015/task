package errorx

import "errors"

type withCode interface {
	code() string
}

var (
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidMethod      = errors.New("invalid method")
	ErrInvalidURL         = errors.New("invalid url")
	ErrTooManyJobs        = errors.New("too many requests, please try later")
	ErrInvalidJobID       = errors.New("invalid job id")
	ErrJobNotFound        = errors.New("job not found")

	errorToCode = map[error]string{
		ErrInvalidRequestBody: "invalid_request_body",
		ErrInvalidMethod:      "invalid_method",
		ErrInvalidURL:         "invalid_url",
		ErrTooManyJobs:        "many_requests",
		ErrInvalidJobID:       "invalid_job_id",
		ErrJobNotFound:        "job_not_found",
	}
)

func Code(err error) string {
	wc, ok := err.(withCode)
	if ok {
		return wc.code()
	}

	code, ok := errorToCode[err]
	if !ok {
		return "not_implemented"
	}

	return code
}

type Error struct {
	Name  string         `json:"name"`
	Codes []ErrorDetails `json:"codes"`
}

type Result struct {
	Details string   `json:"details,omitempty"`
	Code    string   `json:"code,omitempty"`
	Errors  []*Error `json:"errors,omitempty"`
}

// CodeError is a generic error with code value
func CodeError(code string, err error) *Result {
	return &Result{
		Code:    code,
		Details: err.Error(),
	}
}

func ErrorDetailsFn(err error) ErrorDetails {
	return ErrorDetails{
		Message: err.Error(),
		Code:    Code(err),
	}
}

type ErrorDetails struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (r *Result) AddFieldError(field string, ed ErrorDetails) *Result {
	for _, e := range r.Errors {
		if e.Name == field {
			e.Codes = append(e.Codes, ed)
			return r
		}
	}
	r.Errors = append(r.Errors, &Error{
		Name:  field,
		Codes: []ErrorDetails{ed},
	})
	return r
}

func (r *Result) IsValid() bool {
	return len(r.Errors) == 0 && r.Details == ""
}
