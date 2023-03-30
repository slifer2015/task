package task

import (
	"bytes"
	"net/http"
)

func (s *Service) runSingleWorker() {
	go func() {
		for {
			data := <-s.jobBuffer
			s.process(data)
		}
	}()
}

func (s *Service) process(data CreateTaskInput) {
	s.storage.Set(data.ID, JobResult{
		Status: statusInProgress,
	})

	req, err := http.NewRequest(data.Method, data.URL, bytes.NewReader(data.Body))
	if err != nil {
		s.storage.Set(data.ID, JobResult{
			Status: statusError,
		})
		return
	}

	if len(data.QueryParams) > 0 {
		q := req.URL.Query()
		for k, v := range data.QueryParams {
			q.Add(k, v)
		}
	}
	for k, v := range data.Headers {
		req.Header.Set(k, v)
	}

	resp, statusCode, headers, err := s.requester.Do(req)
	if err != nil {
		s.storage.Set(data.ID, JobResult{
			Status:         statusError,
			HTTPStatusCode: 0,
		})
		return
	}
	if statusCode < 200 || statusCode >= 300 {
		s.storage.Set(data.ID, JobResult{
			Status:         statusError,
			HTTPStatusCode: statusCode,
			Headers:        headers,
			Length:         len(resp),
			Response:       string(resp),
		})
		return
	}
	s.storage.Set(data.ID, JobResult{
		Status:         statusDone,
		HTTPStatusCode: statusCode,
		Headers:        headers,
		Length:         len(resp),
		Response:       string(resp),
	})
}
