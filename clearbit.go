package clearbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	userAgent      = "github.com/pzurek/clearbit-" + libraryVersion
)

type Client struct {
	client    *http.Client
	key       string
	UserAgent string

	// Services used for talking to different parts of the GitHub API.
	Enrichements *EnrichmentService
}

func NewClient(key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient, UserAgent: userAgent}
	c.key = key

	c.Enrichements = &EnrichmentService{client: c}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, urlStr, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) error {

	req.SetBasicAuth(c.key, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)

		if err == io.EOF {
			err = nil // ignore EOF errors caused by empty response body
		}
	}

	return err
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

type ErrorResponse struct {
	Response      *http.Response // HTTP response that caused this error
	ResponseError Error          `json:"error"` // more detail on individual errors
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %v",
		r.Response.Request.Method,
		sanitizeURL(r.Response.Request.URL),
		r.Response.Status)
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %s\n%s.\n", e.Type, e.Message)
}

func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}
