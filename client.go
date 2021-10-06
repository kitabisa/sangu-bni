package bni

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
	"moul.io/http2curl"
)

// Client BNI Client data
type Client struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
	LogLevel     int
	Logger       Logger
	HTTPOption   HTTPOption
}

// HTTPOption for heimdall properties
type HTTPOption struct {
	Timeout           time.Duration
	BackoffInterval   time.Duration
	MaxJitterInterval time.Duration
	RetryCount        int
}

// NewClient : this function will always be called when the library is in use
func NewClient() Client {
	logOption := LogOption{
		Format:          "text",
		Level:           "info",
		TimestampFormat: "2006-01-02T15:04:05-0700",
		CallerToggle:    false,
	}

	// default HTTP Option
	httpOption := HTTPOption{
		Timeout:           10 * time.Second,
		BackoffInterval:   2 * time.Millisecond,
		MaxJitterInterval: 5 * time.Millisecond,
		RetryCount:        3,
	}

	logger := *NewLogger(logOption)

	return Client{
		// LogLevel is the logging level used by the BNI library
		// 0: No logging
		// 1: Errors only
		// 2: Errors + informational (default)
		// 3: Errors + informational + debug
		LogLevel:   2,
		Logger:     logger,
		HTTPOption: httpOption,
	}
}

// getHTTPClient will get heimdall http client
func getHTTPClient(opt HTTPOption) *httpclient.Client {
	backoff := heimdall.NewConstantBackoff(opt.BackoffInterval, opt.MaxJitterInterval)
	retrier := heimdall.NewRetrier(backoff)

	return httpclient.NewClient(
		httpclient.WithHTTPTimeout(opt.Timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(opt.RetryCount),
	)
}

// NewRequest : send new request
func (c *Client) NewRequest(method string, fullPath string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		c.Logger.Error("Request creation failed: %v", err)
		return nil, err
	}

	if headers != nil {
		for k, vv := range headers {
			req.Header.Set(k, vv)
		}
	}

	return req, nil
}

// ExecuteRequest : execute request
func (c *Client) ExecuteRequest(req *http.Request, v *map[string]interface{}) error {

	command, _ := http2curl.GetCurlCommand(req)
	start := time.Now()
	c.Logger.Info("Start requesting: %v ", req.URL)
	res, err := getHTTPClient(c.HTTPOption).Do(req)
	if err != nil {
		c.Logger.Error("Request failed. Error : %v , Curl Request : %v", err, command)
		return err
	}
	defer res.Body.Close()

	c.Logger.Info("Completed in %v", time.Since(start))
	c.Logger.Info("Curl Request: %v ", command)

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error("Cannot read response body: %v ", err)
		return err
	}

	c.Logger.Info("BNI HTTP status response : %d", res.StatusCode)
	c.Logger.Info("BNI response body : %s", string(respBody))

	if res.StatusCode != 200 {
		err = fmt.Errorf("%d: %s", res.StatusCode, http.StatusText(res.StatusCode))
		return err
	}

	err = json.Unmarshal(respBody, v)
	if err != nil {
		return err
	}

	return nil
}

// Call the BNI API at specific `path` using the specified HTTP `method`. The result will be
// given to `v` if there is no error. If any error occurred, the return of this function is the error
// itself, otherwise nil.
func (c *Client) Call(method, path string, header map[string]string, body io.Reader, v *map[string]interface{}) error {
	req, err := c.NewRequest(method, path, header, body)

	if err != nil {
		return err
	}

	return c.ExecuteRequest(req, v)
}

// ===================== END HTTP CLIENT ================================================
