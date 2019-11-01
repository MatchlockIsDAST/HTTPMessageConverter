package tostring

import (
	"bytes"
	"io"
	"net/http"
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
	return stringrequest
}

//Respons net/http Response Structをstring型に変換する
func Respons(rawresponse *http.Response) (stringresponse string) {
	return stringresponse
}

//Body io.ReadCloser (http.Body)をstring型に変換する
func Body(rawbody io.ReadCloser) (stringbody string) {
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

	return stringhead
}

//
