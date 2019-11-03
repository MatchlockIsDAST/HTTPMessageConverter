package tostring

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

//Request net/http Request Structをstring型に変換する
/*
e.g
method url proto
header name: header value
:
:

body...

*/
func Request(rawrequest *http.Request) (stringrequest string) {
	q := rawrequest.URL.Query()
	u := rawrequest.URL.Path + "?" + q.Encode()
	stringrequest += fmt.Sprintf("%v %v %v\r\n", rawrequest.Method, u, rawrequest.Proto)
	stringrequest += Header(rawrequest.Header) + "\r\n\r\n"
	stringrequest += Body(rawrequest.Body) + "\r\n"
	return stringrequest
}

//Respons net/http Response Structをstring型に変換する
func Respons(rawresponse *http.Response) (stringresponse string) {
	stringresponse += fmt.Sprintf("%v %v\r\n", rawresponse.Proto, rawresponse.Status)
	stringresponse += Header(rawresponse.Header) + "\r\n\r\n"
	stringresponse += Body(rawresponse.Body) + "\r\n"
	return stringresponse
}

//Body io.ReadCloser (http.Body)をstring型に変換する
func Body(rawbody io.ReadCloser) (stringbody string) {
	if rawbody == nil {
		return ""
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(rawbody)
	stringbody = buf.String()
	return stringbody
}

//Header http.Headerをstring型に変換します
/*
e,g

hoge: hogehoge
fuga: fugafuga
*/
func Header(rawheader http.Header) (stringhead string) {
	headers := []string{}
	for name := range rawheader {
		headers = append(headers, name)
	}
	for i, name := range headers {
		headers[i] = fmt.Sprintf("%v: %v", name, strings.Join(rawheader[name], ", "))
	}
	stringhead = strings.Join(headers, "\r\n")
	return stringhead
}

//
