package requester

import (
	"net/http"
)

type Requester interface {
	Do(req *http.Request) ([]byte, int, map[string][]string, error)
}
