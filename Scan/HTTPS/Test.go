/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 20:16:40
 * @LastEditTime: 2022-10-20 22:03:58
 * @LastEditors: NyanCatda
 * @Description: 测试HTTPS代理是否可用
 * @FilePath: \Cherino\Scan\HTTPS\Test.go
 */
package HTTPS

import (
	"net/http"
	"net/url"
	"time"
)

/**
 * @description: 测试代理是否可用
 * @param {string} URL 代理地址
 * @return {bool} 是否可用
 */
func Test(URL string) bool {
	// 设置代理参数
	URI, err := url.Parse("https://" + URL)
	if err != nil {
		return false
	}
	HTTPProxy := http.ProxyURL(URI)

	httpTransport := &http.Transport{
		Proxy: HTTPProxy,
	}

	// 设置超时时间
	TimeOut := time.Second * 2

	// 设置请求参数
	httpClient := &http.Client{
		Transport: httpTransport,
		Timeout:   TimeOut,
	}

	// 测试是否可以正常链接
	Resp, err := httpClient.Get("http://www.gstatic.com/generate_204")
	if err != nil {
		return false
	}
	defer Resp.Body.Close()

	if Resp.StatusCode != 204 {
		return false
	}

	return true
}
