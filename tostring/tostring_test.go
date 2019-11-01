package tostring

import (
	"testing"
)

func TestRequest(t *testing.T) {
	value := Request(&testRequest)
	if value == "" {
		t.Fatal("空文字が返されています")
	}
}

func TestRespons(t *testing.T) {
	value := Respons(&testResponse)
	if value == "" {
		t.Fatal("空文字が返されています")
	}
}

func TestBody(t *testing.T) {
	value := Body(testBody)
	if value != testBodystring {
		t.Fatal("正しい文字列が帰ってきていません")
	}
}

func TestHeader(t *testing.T) {
	value := Header(testHeaders)
	if value != testHeaderRequest {
		t.Fatal("空文字が返されています")
	}
}
