package requester

import (
	"io"
	"net/http"
)

type Rest struct {
}

func (r *Rest) Do(req *http.Request) ([]byte, int, map[string][]string, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, nil, err
	}
	return data, resp.StatusCode, resp.Header, nil
}

func NewRequester() *Rest {
	return &Rest{}
}
