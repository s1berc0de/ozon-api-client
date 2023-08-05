package request

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func Send[T any](c *http.Client, req *http.Request, contentType ContentType) (*T, *http.Response, error) {
	out := new(T)

	req.Header.Set("Content-Type", string(contentType))

	resp, err := c.Do(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "can't do request")
	}

	defer resp.Body.Close() //nolint

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, errors.Wrap(err, "can't read all response data")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, resp, errors.Errorf("invalid status code: %d, %s", resp.StatusCode, resp.Status)
	}

	if err = json.Unmarshal(body, out); err != nil {
		return nil, resp, errors.Wrap(err, "can't decode response")
	}

	return out, resp, nil
}
