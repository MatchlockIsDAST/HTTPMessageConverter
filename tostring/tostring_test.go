package tostring

import (
	"testing"
)

func TestRequest(t *testing.T) {
	value := Request(&testRequest)
	if value != testRequestMessage {
		t.Fatal("正しい文字列が帰ってきていません")
	}
	t.Error()
}

func TestRespons(t *testing.T) {
	value := Respons(&testResponse)
	if value != testResponseMessage {
		t.Fatal("正しい文字列が帰ってきていません")
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
		t.Fatal("正しい文字列が帰ってきていません")
	}
}
