package inmemory

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	TestRequest = http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme:   "https",
			Host:     "google.com",
			Path:     "/search",
			RawPath:  "/search",
			RawQuery: "q=otnet",
		},
		Header:        TestHeaders,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          nil,
		ContentLength: 0,
		Host:          "google.com",
		Form:          nil,
	}
	TestResponse = http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header: http.Header{
			"Cache-Control":    []string{"max-age=0", "private", "must-revalidate"},
			"Content-Encoding": []string{"gzip"},
		},
		Body: TestBody,
	}
	TestBodystring = "<!DOCTYPE html>\n<html>\n<header>\n</header>\n<body>\nTest Data\n</body>\n</html>"
	TestBody       = ioutil.NopCloser(strings.NewReader(TestBodystring))
	TestHeaders    = http.Header{
		"accept":          []string{"text/html", "application/xhtml+xml", "application/xml;q=0.9", "image/webp", "image/apng", "*/*;q=0.8", "application/signed-exchange;v=b3"},
		"accept-encoding": []string{"gzip", "deflate", "br"},
		"accept-language": []string{"ja"},
		"host":            []string{"google.com"},
	}
	TestStartLineRequest     = "GET /search?q=otnet HTTP/1.1\r\n"
	TestStartLinePOSTRequest = "POST /search HTTP/1.1\r\n"
	TestStartLineResponse    = "HTTP/1.1 200 OK\r\n"
	TestHeaderRequest        = "accept: text/html, application/xhtml+xml, application/xml;q=0.9, image/webp, image/apng, */*;q=0.8, application/signed-exchange;v=b3\r\naccept-encoding: gzip, deflate, br\r\naccept-language: ja\r\nhost: google.com"
	TestHeaderResponse       = "Cache-Control: max-age=0, private, must-revalidate\r\nContent-Encoding: gzip"
	TestRequestMessage       = TestStartLineRequest + TestHeaderRequest + "\r\n\r\n\r\n"
	TestResponseMessage      = TestStartLineResponse + TestHeaderResponse + "\r\n\r\n" + TestBodystring + "\r\n"
	TestPOSTRequestMessage   = TestStartLinePOSTRequest + TestHeaderRequest + "\r\n\r\n" + "q=hoge&fuga=fuga" + "\r\n"
)
