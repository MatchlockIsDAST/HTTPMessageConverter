package stringto

import (
	"reflect"
	"strings"
	"testing"

	"github.com/MatchlockIsDAST/HTTPMessageConverter/inmemory"
)

func TestSplitHTTPRequestMessageForLine(t *testing.T) {
	s, h, b := SplitHTTPRequestMessageForLine(inmemory.TestRequestMessage)
	if s != "GET /search?q=otnet HTTP/1.1" {
		t.Fatal("GET : Start lineが正しくありません")
	}
	if !reflect.DeepEqual(h, []string{"accept: text/html, application/xhtml+xml, application/xml;q=0.9, image/webp, image/apng, */*;q=0.8, application/signed-exchange;v=b3", "accept-encoding: gzip, deflate, br", "accept-language: ja", "host: google.com"}) {
		t.Fatal("GET : Headerが異なります")
	}
	if b != "" {
		t.Fatal("GET : Bodyが異なります")
	}
	s, h, b = SplitHTTPRequestMessageForLine(inmemory.TestPOSTRequestMessage)
	if s != "POST /search HTTP/1.1" {
		t.Fatal("POST : Start lineが正しくありません")
	}
	if !reflect.DeepEqual(h, []string{"accept: text/html, application/xhtml+xml, application/xml;q=0.9, image/webp, image/apng, */*;q=0.8, application/signed-exchange;v=b3", "accept-encoding: gzip, deflate, br", "accept-language: ja", "host: google.com"}) {
		t.Fatal("POST : Headerが異なります")
	}
	if b != "q=hoge&fuga=fuga" {
		t.Error(b)

		t.Fatal("POST : Bodyが異なります")
	}
}
func TestIoReadCloser(t *testing.T) {

}
func TestSplitStartline(t *testing.T) {
	testStart := strings.Replace(inmemory.TestStartLineRequest, "\r\n", "", -1) //\r\nを削除
	method, uri, proto := SplitStartline(testStart)
	if method != "GET" {
		t.Fatal("Methodが一致しません")
	}
	if uri != "/search?q=otnet" {
		t.Fatal("URIが一致しません")
	}
	if proto != "HTTP/1.1" {
		t.Fatal("Protoが一致しません")
	}
}
func TestHTTPHeader(t *testing.T) {
	testHeader := []string{"accept: text/html, application/xhtml+xml, application/xml;q=0.9, image/webp, image/apng, */*;q=0.8, application/signed-exchange;v=b3", "accept-encoding: gzip, deflate, br", "accept-language: ja", "host: google.com"}
	result := HTTPHeader(testHeader)
	if !reflect.DeepEqual(result["accept"], []string{"text/html", "application/xhtml+xml", "application/xml;q=0.9", "image/webp", "image/apng", "*/*;q=0.8", "application/signed-exchange;v=b3"}) {
		t.Error("取得された値が異なります")
	}
}
