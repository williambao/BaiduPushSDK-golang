// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//封装了http调用相关的一些方法
package push

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	//"strconv"
	"strings"
)

//执行http请求
func httpExecute(
	method string, urlStr string, contentType string, body string, oauthParams *OrderedParams) (*http.Response, error) {
	// Create base request.
	v := url.Values{}
	for _, key := range oauthParams.Keys() {
		v.Add(key, oauthParams.Get(key))
	}
	req, err := http.NewRequest(method, urlStr, strings.NewReader(v.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Add("User-Agent", "BCCS_SDK/3.0 Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	if err != nil {
		return nil, errors.New("NewRequest failed: " + err.Error())
	}
	HttpClient := &http.Client{}
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, errors.New("Do: " + err.Error())
	}

	debugHeader := ""
	for k, vals := range req.Header {
		for _, val := range vals {
			debugHeader += "[key: " + k + ", val: " + val + "]"
		}
	}

	// StatusMultipleChoices is 300, any 2xx response should be treated as success
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)
		return resp, errors.New(string(bytes))
	}
	return resp, err
}

//获得HTTP请求的body部分内容
func getBody(method, url string, oauthParams *OrderedParams) (*string, error) {
	resp, err := httpExecute(method, url, "", "", oauthParams)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	bodyStr := string(bodyBytes)
	/*
		if c.debug {
			fmt.Printf("STATUS: %d %s\n", resp.StatusCode, resp.Status)
			fmt.Println("BODY RESPONSE: " + bodyStr)
		}
	*/
	return &bodyStr, nil
}

//计算签名的字符串
func requestString(method string, urlPath string, secretkey string, params *OrderedParams) string {
	result := method + urlPath
	for _, key := range params.Keys() {
		result += fmt.Sprintf("%s=%s", key, params.Get(key))
	}
	return result + secretkey
}

//调用API
func CallApiServer(httpMethod string, server string, class string, method string, params *OrderedParams, secretkey string, i interface{}) error {
	reqString := requestString(httpMethod, server+class+method, secretkey, params)
	h := md5.New()
	h.Write([]byte(url.QueryEscape(reqString)))
	signature := hex.EncodeToString(h.Sum(nil))
	params.Add("sign", signature)
	result, err := getBody("POST", server+class+method, params)
	if err == nil {
		json.Unmarshal([]byte(*result), &i)
		return nil
	} else {
		return err
	}
}

//排序后的参数列表
type OrderedParams struct {
	allParams   map[string]string
	keyOrdering []string
}

func NewOrderedParams() *OrderedParams {
	return &OrderedParams{
		allParams:   make(map[string]string),
		keyOrdering: make([]string, 0),
	}
}

func (o *OrderedParams) Get(key string) string {
	return o.allParams[key]
}

func (o *OrderedParams) Keys() []string {
	sort.Sort(o)
	return o.keyOrdering
}

func (o *OrderedParams) Add(key, value string) {
	o.AddUnescaped(key, url.QueryEscape(value))
}

func (o *OrderedParams) AddUnescaped(key, value string) {
	o.allParams[key] = value
	o.keyOrdering = append(o.keyOrdering, key)
}

func (o *OrderedParams) Len() int {
	return len(o.keyOrdering)
}

func (o *OrderedParams) Less(i int, j int) bool {
	return o.keyOrdering[i] < o.keyOrdering[j]
}

func (o *OrderedParams) Swap(i int, j int) {
	o.keyOrdering[i], o.keyOrdering[j] = o.keyOrdering[j], o.keyOrdering[i]
}

func (o *OrderedParams) Clone() *OrderedParams {
	clone := NewOrderedParams()
	for _, key := range o.Keys() {
		clone.AddUnescaped(key, o.Get(key))
	}
	return clone
}
