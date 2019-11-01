package tostring

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var testRequest = http.Request{}
var testResponse = http.Response{}
var testBody = ioutil.NopCloser(strings.NewReader(""))
var testHeader = http.Header{}
