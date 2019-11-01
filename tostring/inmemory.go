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
		Header:        testHeaders,
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
	testBodystring = "<!DOCTYPE html>\n<html>\n<header>\n</header>\n<body>\nTest Data\n</body>\n</html>\n"
	testBody       = ioutil.NopCloser(strings.NewReader(testBodystring))
	testHeaders    = http.Header{
		"accept":          []string{"text/html", "application/xhtml+xml", "application/xml;q=0.9", "image/webp", "image/apng", "*/*;q=0.8", "application/signed-exchange;v=b3"},
		"accept-encoding": []string{"gzip", "deflate", "br"},
		"accept-language": []string{"ja"},
		"host":            []string{"google.com"},
	}
	testStartLineRequest  = "GET https://google.com/search?q=otnet HTTP/1.1\r\n"
	testStartLineResponse = "HTTP/1.1 200 OK\r\n"
	testHeaderRequest     = "accept: text/html, application/xhtml+xml, application/xml;q=0.9, image/webp, image/apng, */*;q=0.8, application/signed-exchange;v=b3\r\naccept-encoding: gzip, deflate, br\r\naccept-language: ja\r\nhost: google.com"
	testHeaderResponse    = "Cache-Control: max-age=0, private, must-revalidate\r\nContent-Encoding: gzip"
	testRequestMessage    = testStartLineRequest + testHeaderRequest + "\r\n\r\n\r\n"
	testResponseMessage   = testStartLineResponse + testHeaderResponse + "\r\n\r\n" + testBodystring + "\r\n"
)
