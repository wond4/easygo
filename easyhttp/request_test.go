package easyhttp

import (
	"io"
	"io/ioutil"
	"testing"
)

func readString(r io.Reader) string {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return ""
	}
	return string(bs)
}

func assertEqual(t *testing.T, a, b interface{}) {
	t.Helper()
	assert(t, a == b, "not equal. %T: %v != %T: %v", a, a, b, b)
}

func assert(t *testing.T, a bool, hintFormat string, args ...interface{}) {
	t.Helper()
	if !a {
		if hintFormat == "" {
			t.Error("assert failed")
		} else {
			t.Errorf(hintFormat, args...)
		}
	}
}

func TestNewReqBodyByFormUrl(t *testing.T) {
	data := map[string]string{
		"abc": "123",
		"ddd": "中国人",
	}
	assertEqual(t, readString(NewReqBodyByFormUrl(data)), "abc=123&ddd=%E4%B8%AD%E5%9B%BD%E4%BA%BA")
	assertEqual(t, readString(NewReqBodyByFormUrl(nil)), "")
}
