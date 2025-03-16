package common

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ParseReqBody(r *http.Request, req any) error {
	if r.Body != nil {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return errors.New("failed to read request body")
		}
		defer r.Body.Close()

		if len(body) > 0 {
			if err := json.Unmarshal(body, req); err != nil {
				return errors.New("invalid JSON")
			}
		}
	}

	return nil
}
