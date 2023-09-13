package easyhttp

import "testing"

func TestNewClient(t *testing.T) {
	c1 := NewClient(OptIdle)
	assert(t, c1.Transport == nil, "")
	assert(t, c1.Jar == nil, "")
	c2 := NewClient(OptCookieJar)
	assert(t, c2.Transport == nil, "")
	assert(t, c2.Jar != nil, "")
	c3 := NewClient(OptNoSSL)
	assert(t, c3.Transport != nil, "")
	assert(t, c3.Jar == nil, "")
	c4 := NewClient(OptCookieJar | OptNoSSL)
	assert(t, c4.Transport != nil, "")
	assert(t, c4.Jar == nil, "")
}
