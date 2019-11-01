package tostring

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	testRequest = http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme:   "https",
			Host:     "google.com",
			Path:     "/search",
			RawPath:  "/search",
			RawQuery: "q=otnet",
		},
		Header:        testHeader,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          nil,
		ContentLength: 0,
		Host:          "google.com",
		Form:          nil,
	}
	testResponse = http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header: http.Header{
			"Cache-Control":    []string{"max-age=0", "private", "must-revalidate"},
			"Content-Encoding": []string{"gzip"},
		},
	}
	testBodystring = "<!DOCTYPE html><html><header></header><body>Test Data</body></html>"
	testBody       = ioutil.NopCloser(strings.NewReader(testBodystring))
	testHeader     = http.Header{
		"Host":            []string{"google.com"},
		"accept":          []string{"text/html", "application/xhtml+xml", "application/xml;q=0.9", "image/webp", "image/apng", "*/*;q=0.8", "application/signed-exchange;v=b3"},
		"accept-encoding": []string{"gzip", "deflate", "br"},
		"accept-language": []string{"ja"},
	}
)
