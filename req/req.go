package req

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
}

func DoReq(url, option, path string, headers []string, args string, body string) error {
	var err error
	var req *http.Request

	uri := url + path
	if args != "" {
		uri = uri + "?" + args
	}

	reqBody := strings.NewReader(body)

	option = strings.ToUpper(option)
	switch option {
	case "POST":
		req, err = http.NewRequest("POST", uri, reqBody)
	case "OPTIONS":
		req, err = http.NewRequest("OPTIONS", uri, nil)
	default:
		req, err = http.NewRequest("GET", uri, nil)
	}
	if err != nil {
		return err
	}

	for _, header := range headers {
		h := strings.SplitN(header, ":", 2)
		if len(h) == 2 {
			req.Header.Add(h[0], h[1])
		}
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	formatRes(res)
	return nil
}

func formatRes(res *http.Response) {
	req := res.Request
	uri := req.URL.RequestURI()

	s1 := fmt.Sprintf("%s %s %s\nHost: %s\n%s",
		req.Method, uri, req.Proto, req.Host, joinHeaders(req.Header))
	fmt.Println(s1)

	s := fmt.Sprintf("%v %v\n", res.Proto, res.Status)
	s = fmt.Sprintf("%v%v", s, joinHeaders(res.Header))
	fmt.Println(s)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func joinHeaders(headers http.Header) string {
	var buffer bytes.Buffer
	for k, v := range headers {
		buffer.WriteString(k + ": " + strings.Join(v, ";") + "\n")
	}
	s := buffer.String()
	return s

}
