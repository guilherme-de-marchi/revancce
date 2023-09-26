package pkg

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func NewHTTPRequest(method, endpoint string, body any) (int, string, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return 0, "", err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		ServerDomain+endpoint,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return 0, "", err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, "", err
	}

	respBody, err := io.ReadAll(resp.Body)
	return resp.StatusCode, string(respBody), err
}
