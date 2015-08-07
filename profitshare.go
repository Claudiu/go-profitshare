package profitshare

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const apiDomain = "http://api.profitshare.ro/"

// ProfitShare API instance
type ProfitShare struct {
	user      string
	key       string
	SleepTime time.Duration
}

type Paginator struct {
	ItemsPerPage int `json:"itemsPerPage"`
	CurrentPage  int `json:"currentPage"`
	TotalPages   int `json:"totalPages"`
}

// APIError holds API error information
type APIError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Creates a new Instance of ProfitShare API
func NewProfitShare(user, key string) (ps *ProfitShare) {
	return &ProfitShare{user, key, time.Millisecond * 200}
}

func (ps *ProfitShare) createAuth(verb, method, query string) (string, string) {
	date := time.Now().Format("Mon, 02 Jan 2006 15:04:05 -0700")

	signature := verb + method + "?" + query + "/" + ps.user + date
	//fmt.Println(signature)
	mac := hmac.New(sha1.New, []byte(ps.key))
	mac.Write([]byte(signature))
	key := hex.EncodeToString(mac.Sum(nil))

	return key, date
}

func (ps *ProfitShare) request(method, uri string, postValues *url.Values) []byte {
	client := &http.Client{}

	url, err := url.Parse(uri)

	if err != nil {
		panic(fmt.Errorf("Error setting-up request: %s", err.Error()))
	}

	req, err := http.NewRequest(method, apiDomain+url.String(), nil)

	if method == "POST" {
		pvreader := strings.NewReader(postValues.Encode())
		req.Body = ioutil.NopCloser(pvreader)
	}

	key, date := ps.createAuth(method, url.Path, url.RawQuery)

	if err != nil {
		panic(fmt.Errorf("Error setting-up request: %s", err.Error()))
	}

	req.Header.Add("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	req.Header.Add("X-PS-Client", ps.user)
	req.Header.Add("X-PS-Accept", "json")
	req.Header.Add("X-PS-Auth", key)
	req.Header.Add("Date", date)

	resp, err := client.Do(req)

	if err != nil {
		panic(fmt.Errorf("Error setting-up request: %s", err.Error()))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(fmt.Errorf("Error setting-up request: %s", err.Error()))
	}

	if strings.Contains(string(body), "error") {
		rez := APIError{}
		_ = json.Unmarshal(body, &rez)

		panic(fmt.Errorf("API Error: %s (%s)", rez.Error.Message, rez.Error.Code))
	}

	return body
}

func (ps *ProfitShare) Get(uri string) []byte {
	return ps.request("GET", uri, nil)
}
