package gobitlaunch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	version      = "1.1.0"
	userAgent    = "gobitlaunch/" + version
	maxRateLimit = 900 * time.Millisecond
	retryLimit   = 3
	baseURI      = "https://app.bitlaunch.io/api"
)

// Client manages interaction with the API
type Client struct {
	token   string
	hclient *retryablehttp.Client

	Account       *AccountService
	Server        *ServerService
	Transaction   *TransactionService
	CreateOptions *CreateOptionsService
	SSHKey        *SSHKeyService
}

// NewClient takes an API token and returns a new BitLaunch API client
func NewClient(token string) *Client {
	c := Client{
		token:   token,
		hclient: retryablehttp.NewClient(),
	}

	// set up http client
	c.hclient.Logger = nil
	c.hclient.ErrorHandler = c.errorHandler
	c.hclient.RetryMax = retryLimit
	c.hclient.RetryWaitMin = maxRateLimit / 3
	c.hclient.RetryWaitMax = maxRateLimit

	// add services
	c.Account = &AccountService{&c}
	c.Server = &ServerService{&c}
	c.Transaction = &TransactionService{&c}
	c.CreateOptions = &CreateOptionsService{&c}
	c.SSHKey = &SSHKeyService{&c}

	return &c
}

// NewRequest creates an API Request
func (c *Client) NewRequest(method, path string, body []byte) (*http.Request, error) {

	fullURI := baseURI + path

	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, fullURI, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer: "+c.token)
	req.Header.Add("User-Agent", userAgent)

	return req, nil
}

// DoRequest performs a http request
func (c *Client) DoRequest(r *http.Request, data interface{}) error {
	rreq, err := retryablehttp.FromRequest(r)
	if err != nil {
		return err
	}

	res, err := c.hclient.Do(rreq)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK {
		if data != nil {
			if err := json.Unmarshal(body, data); err != nil {
				return err
			}
		}
		return nil
	}

	return fmt.Errorf("error %d %s", res.StatusCode, string(body))
}

func (c *Client) errorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	if resp == nil {
		return nil, fmt.Errorf("gave up after %d attempts, last error unavailable (resp == nil)", retryLimit+1)
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gave up after %d attempts, last error unavailable (error reading response body: %v)", retryLimit+1, err)
	}

	return nil, fmt.Errorf("gave up after %d attempts, last error: %#v", retryLimit+1, strings.TrimSpace(string(buf)))
}
