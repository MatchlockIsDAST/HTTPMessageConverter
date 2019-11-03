package stringto

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//SplitHTTPRequestMessageForLine はrawrequest (string型)を入力することによってstartlineとheaderlineとbodyに分割する
//現状Chankに対応していない
func SplitHTTPRequestMessageForLine(rawrequest string) (start string, header []string, body string) {
	rawrequest = strings.Replace(rawrequest, "\r", "", -1)
	splited := strings.Split(rawrequest, "\n\n")
	if len(splited) > 1 {
		body = strings.Replace(splited[1], "\n", "", -1)
	}
	splited = strings.Split(splited[0], "\n")
	start = splited[0]
	header = splited[1:]
	return start, header, body
}

//IoReadCloser string型をio.ReadCloser型に変換します
func IoReadCloser(rawbody string) (ioReadCloser io.ReadCloser) {
	ioReadCloser = ioutil.NopCloser(strings.NewReader(rawbody))
	return ioReadCloser
}

//SplitStartline startlineをmethod, uriAndQuery, protoに分割します
func SplitStartline(rawstartline string) (method, uriAndQuery, proto string) {
	splited := strings.Split(rawstartline, " ")
	method, uriAndQuery, proto = splited[0], splited[1], splited[2]
	return method, uriAndQuery, proto
}

//HTTPHeader string型をhttp.Header型に変換します
func HTTPHeader(rawheaders []string) (httpheader http.Header) {
	httpheader = http.Header{}
	for _, v := range rawheaders {
		splitHeader := strings.Split(v, ": ")
		val := strings.Replace(splitHeader[1], ", ", ",", -1)
		httpheader[splitHeader[0]] = strings.Split(val, ",")
	}
	return httpheader
}
