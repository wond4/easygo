package easyhttp

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
)

// NewClient 使用的选项可使用 | 来设置多个
const (
	OptIdle      = 0         // 不使用选项
	OptCookieJar = 1 << iota // 使用 cookie
	OptNoSSL                 // 不使用 ssl
)

func NewClient(option uint) *http.Client {
	c := &http.Client{}
	if OptCookieJar&option != 0 {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return nil
		}
		c.Jar = jar
	}
	if OptNoSSL&option != 0 {
		tp := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		c.Transport = tp
	}
	return c
}
