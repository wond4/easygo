package easyhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
)

// NewReqBodyByStr 通过 string 生成 request body
func NewReqBodyByStr(str string) io.Reader {
	return strings.NewReader(str)
}

// NewReqBodyByJson 通过序列号 json 生成 request body
func NewReqBodyByJson(obj interface{}) io.Reader {
	content, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	return bytes.NewReader(content)
}

// NewReqBodyByFormUrl 通过 map 生成 application/x-www-form-urlencoded 格式 request body
func NewReqBodyByFormUrl(strMap map[string]string) io.Reader {
	data := &url.Values{}
	for k, v := range strMap {
		data.Add(k, v)
	}
	return strings.NewReader(data.Encode())
}

// NewReqBodyByFormData 通过 map 生成 multipart/form-data 格式 request body
func NewReqBodyByFormData(strMap map[string]string, fileMap map[string]string) (body io.Reader, contentType string) {
	buf := &bytes.Buffer{}
	form := multipart.NewWriter(buf)
	defer func() { _ = form.Close() }()
	for k, v := range strMap {
		err := form.WriteField(k, v)
		if err != nil {
			return nil, ""
		}
	}
	for k, v := range fileMap {
		_, err := form.CreateFormFile(k, v)
		if err != nil {
			return nil, ""
		}
	}
	return buf, form.FormDataContentType()
}
