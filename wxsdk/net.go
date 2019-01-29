package wxsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// net 请求结构体
type net struct {
	client    *http.Client // 可重复使用的client
	baseURL   *url.URL     // 请求根地址
	isRelease bool         // 是否是主产环境
}

// superAgent 请求参数
type superAgent struct {
	net         *net              // 当前请求包实例
	url         string            // 请求地址
	method      string            // 请求方式
	contentType string            // 请求类型
	body        interface{}       // 发送请求的body
	header      map[string]string // 头文件
}

const (
	contentJSON = "application/json;charset=utf-8"
	contentXML  = "application/xml;charset=utf-8"
	contentText = "text/plain;charset=utf-8"
)

// newNet 初始化一个请求包对象
func newNet() *net {
	var release bool
	netMode := os.Getenv("NET_MODE")
	if len(netMode) == 0 {
		release = false
	} else if netMode == "release" {
		release = true
	}

	return &net{
		client:    http.DefaultClient,
		isRelease: release,
	}
}

// getClient 获取http client
func (n *net) getClient() *http.Client {
	return n.client
}

// newWithClient 初始化一个请求包对象，自己传入Client
func newWithClient(client *http.Client) *net {
	return &net{
		client: client,
	}
}

// get 发送 get 请求
func (n *net) get(url string) *superAgent {
	return &superAgent{net: n, url: url, method: "GET"}
}

// post 发送 post 请求
func (n *net) post(url string) *superAgent {
	return &superAgent{net: n, url: url, method: "POST"}
}

// put 发送 put 请求
func (n *net) put(url string) *superAgent {
	return &superAgent{net: n, url: url, method: "PUT"}
}

// delete 发送 delete 请求
func (n *net) delete(url string) *superAgent {
	return &superAgent{net: n, url: url, method: "DELETE"}
}

// json 设置请求数据内容，默认用 Content-Type=application/json; 方式发送json数据
func (s *superAgent) json(body interface{}) *superAgent {
	s.body = body
	s.contentType = contentJSON
	return s
}

// xml 设置请求数据内容，默认用 Content-Type=application/xml; 方式发送xml数据
func (s *superAgent) xml(body interface{}) *superAgent {
	s.body = body
	s.contentType = contentXML
	return s
}

// text 设置请求数据内容，默认用 Content-Type=text/plain; 方式发送string数据
func (s *superAgent) text(body string) *superAgent {
	s.body = body
	s.contentType = contentText
	return s
}

// setHeader 设置请求头内容
func (s *superAgent) setHeader(header map[string]string) *superAgent {
	s.header = header
	return s
}

// end 开始http请求
func (s *superAgent) end(ctx context.Context, v interface{}) (*http.Response, error) {
	if len(s.contentType) > 0 && s.body == nil {
		s.body = ""
	}
	var req *http.Request
	var err error
	buf := new(bytes.Buffer)
	switch s.contentType {
	case contentJSON:
		err = json.NewEncoder(buf).Encode(s.body)
	case contentXML:
		err = xml.NewEncoder(buf).Encode(s.body)
	case contentText:
		_, err = buf.WriteString(s.body.(string))
	}
	if err != nil {
		return nil, err
	}

	// 转换 url
	rel, err := url.Parse(s.url)
	if err != nil {
		return nil, err
	}
	u := s.net.baseURL.ResolveReference(rel)

	req, err = http.NewRequest(s.method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if len(s.contentType) > 0 {
		req.Header.Set("Content-Type", s.contentType)
	}
	for key, value := range s.header {
		req.Header.Set(key, value)
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// 执行网络请求
	resp, err := s.net.client.Do(req)
	if err != nil {

		if ctx != nil {
			// 如果发生错误，并且上下文已被取消，则返回上下文的错误
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}
		}

		// 如果错误类型为url.Error，在返回之前解析URL
		if e, ok := err.(*url.Error); ok {
			if uri, err := url.Parse(e.URL); err == nil {
				e.URL = uri.String()
				return nil, e
			}
		}
		return nil, err
	}

	defer func() {
		// 耗尽512个字节则关闭body以使Transport重用连接
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	if w, ok := v.(io.Writer); ok {
		io.Copy(w, resp.Body)
	} else {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if !s.net.isRelease {
			log.Printf("url: %s , response body: %s", s.url, string(body))
		}

		if v != nil {
			// 默认为 contentType 不为xml的情况下，所有返回都用json解析
			if strings.EqualFold(s.contentType, contentXML) {
				err = xml.Unmarshal(body, v)
			} else {
				err = json.Unmarshal(body, v)
			}
		}

		if err == io.EOF {
			err = nil
		}
		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf(string(body))
		}
	}

	return resp, err
}
