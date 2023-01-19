package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
	"k8s.io/klog/v2"
)

var dataplaneapi = "http://10.0.0.3:5555"

var (
	// longThrottleLatency defines threshold for logging requests. All requests being
	// throttled (via the provided rateLimiter) for more than longThrottleLatency will
	// be logged.
	longThrottleLatency = 50 * time.Millisecond

	// extraLongThrottleLatency defines the threshold for logging requests at log level 2.
	extraLongThrottleLatency = 1 * time.Second
)

var _ RequestInterface = &RESTClient{}

var noBackoff = &rest.NoBackoff{}

type RESTClient struct {
	Base             *url.URL
	Client           *http.Client
	createBackoffMgr func() rest.BackoffManager
	rateLimiter      flowcontrol.RateLimiter
}

func NewRESTClient(baseURL string) *RESTClient {
	var base *url.URL
	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		base = new(url.URL)
	} else {
		base = parsedUrl
	}

	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}
	base.RawQuery = ""
	base.Fragment = ""

	return &RESTClient{
		Base:             base,
		createBackoffMgr: readExpBackoffConfig,
		rateLimiter:      flowcontrol.NewTokenBucketRateLimiter(100.00, 10),
		Client:           &http.Client{},
	}
}

func NewRESTClientWithProxy(baseURL, proxyUrl string) *RESTClient {
	var (
		base   *url.URL
		client *http.Client
	)
	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		base = new(url.URL)
	} else {
		base = parsedUrl
	}

	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}
	base.RawQuery = ""
	base.Fragment = ""

	parsedProxyUrl, err := url.Parse(proxyUrl)
	if err != nil {
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		}
	} else {
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(parsedProxyUrl),
			},
		}
	}

	return &RESTClient{
		Base:             base,
		createBackoffMgr: readExpBackoffConfig,
		rateLimiter:      flowcontrol.NewTokenBucketRateLimiter(100.00, 10),
		Client:           client,
	}
}

func (c *RESTClient) Get() *Request {
	return c.Verb("GET")
}

func (c *RESTClient) Post() *Request {
	return c.Verb("POST")
}

func (c *RESTClient) Delete() *Request {
	return c.Verb("DELETE")
}

func (c *RESTClient) Put() *Request {
	return c.Verb("PUT")
}

func (c *RESTClient) Verb(verb string) *Request {
	return NewDefaultRequest(c).Verb(verb)
}

func get(url string) (body []byte) {
	client := &http.Client{}
	url = fmt.Sprintf("%s/%s", dataplaneapi, url)
	req, err := http.NewRequest("GET", url, nil)
	//req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicAuth("admin", "1fc917c7ad66487470e466c0ad40ddd45b9f7730a4b43e1b2542627f0596bbdc")))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		klog.V(3).Infof("Get to haproxy failed, uri => %s", url)
		return nil
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		klog.V(3).Infof("Convert to byte failed, uri => %s", url)
		return nil
	}
	return bodyText
}

func post(url string, payload []byte) (resp []byte, status int) {
	url = fmt.Sprintf("%s/%s", dataplaneapi, url)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
	}
	//req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicAuth("admin", "1fc917c7ad66487470e466c0ad40ddd45b9f7730a4b43e1b2542627f0596bbdc")))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	resp, err = ioutil.ReadAll(res.Body)
	status, _ = strconv.Atoi(res.Status)
	return
}

func delete(url string) int {
	url = fmt.Sprintf("%s/%s", dataplaneapi, url)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	//req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicAuth("admin", "1fc917c7ad66487470e466c0ad40ddd45b9f7730a4b43e1b2542627f0596bbdc")))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		klog.V(3).Infof("Delete to haproxy failed, uri => %s", url)
		return 0
	}
	code, _ := strconv.Atoi(resp.Status)
	return code
}

// checkWait returns true along with a number of seconds if the server instructed us to wait
// before retrying.
func checkWait(resp *http.Response) (int, bool) {
	switch r := resp.StatusCode; {
	// any 500 error code and 429 can trigger a wait
	case r == http.StatusTooManyRequests, r >= 500:
	default:
		return 0, false
	}
	i, ok := retryAfterSeconds(resp)
	return i, ok
}

// retryAfterSeconds returns the value of the Retry-After header and true, or 0 and false if
// the header was missing or not a valid number.
func retryAfterSeconds(resp *http.Response) (int, bool) {
	if h := resp.Header.Get("Retry-After"); len(h) > 0 {
		if i, err := strconv.Atoi(h); err == nil {
			return i, true
		}
	}
	return 0, false
}
