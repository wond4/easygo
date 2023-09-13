package easyhttp

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// Body
type Body struct {
	body    io.ReadCloser
	content []byte
}

// NewBody 构造一个新的 Body 实例
func NewBody(body io.ReadCloser) *Body {
	return &Body{body: body}
}

// 丢弃 Body 中的内容并关闭
func (b *Body) Discard() {
	if b.body == nil {
		return
	}
	if b.content != nil {
		b.content = nil
		return
	}
	defer func() {
		_ = b.body.Close()
	}()
	_, _ = io.Copy(ioutil.Discard, b.body)
}

// 以 []byte 返回 Body 中的内容并关闭
func (b *Body) Bytes() (bytes []byte, err error) {
	if b.body == nil {
		return nil, fmt.Errorf("body is null")
	}
	if b.content != nil {
		return b.content, nil
	}
	defer func() {
		err = b.body.Close()
	}()
	b.content, err = ioutil.ReadAll(b.body)
	if err != nil {
		b.content = nil
		return
	}
	return b.content, nil
}

// 以 string 返回 Body 中的内容并关闭
func (b *Body) String() (str string, err error) {
	content, err := b.Bytes()
	if err != nil {
		return
	}
	str = string(content)
	return
}

// 以 string 返回 Body 中的内容并关闭
func (b *Body) Map() (mapObj map[string]interface{}, err error) {
	content, err := b.Bytes()
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &mapObj)
	return
}

// 以 Json 对象序列化 Body 中的内容并关闭
func (b *Body) Json(objPointer interface{}) (err error) {
	content, err := b.Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(content, objPointer)
}
