package stringto

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//SplitHTTPRequestMessageForLine はrawrequest (string型)を入力することによってstartlineとheaderlineとbodyに分割する
func SplitHTTPRequestMessageForLine(rawrequest string) (start string, header string, body string) {
	return start, header, body
}

//IoReadCloser string型をio.ReadCloser型に変換します
func IoReadCloser(rawbody string) (ioReadCloser io.ReadCloser) {
	ioReadCloser = ioutil.NopCloser(strings.NewReader(rawbody))
	return ioReadCloser
}

//HTTPHeader string型をhttp.Header型に変換します
func HTTPHeader(rawheaders []string) (httpheader http.Header) {
	httpheader = map[string][]string{}
	for _, v := range rawheaders {
		splitHeader := strings.Split(v, ":")
		httpheader[splitHeader[0]] = strings.Split(splitHeader[1], ",")
	}
	return httpheader
}
