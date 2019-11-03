package tostring

import (
	"testing"

	"github.com/MatchlockIsDAST/HTTPMessageConverter/inmemory"
)

func TestRequest(t *testing.T) {
	value := Request(&inmemory.TestRequest)
	if value != inmemory.TestRequestMessage { //Header Keyソートが不安定
		t.Fatal("正しい文字列が帰ってきていません")
	}
}

func TestRespons(t *testing.T) {
	value := Respons(&inmemory.TestResponse)
	if value != inmemory.TestResponseMessage {
		t.Fatal("正しい文字列が帰ってきていません")
	}
}

func TestBody(t *testing.T) {
	value := Body(inmemory.TestBody)
	if value != inmemory.TestBodystring {
		t.Fatal("正しい文字列が帰ってきていません")
	}
}

func TestHeader(t *testing.T) {
	value := Header(inmemory.TestHeaders)
	if value != inmemory.TestHeaderRequest { //Header Keyソートが不安定
		t.Fatal("正しい文字列が帰ってきていません")
	}
}
