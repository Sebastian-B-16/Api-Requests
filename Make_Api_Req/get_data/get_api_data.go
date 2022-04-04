package get_data

import (
	"context"
	"io"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type sendRequest struct{}

//Do send an HTTP request and returns an HTTP response
func (s sendRequest) Do(req *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(req)
}

type Extract interface {
	extractData(client HttpClient, ctx context.Context) ([]byte, error)
}
type Input struct {
	Link            string
	NumberOfRecords int
}

func (r *Input) extractData(client HttpClient, ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.Link, nil)
	if err != nil {
		return nil, err
	}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode > 299 {
		return nil, err
	}
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
