/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 17:09:55
 * @LastEditTime: 2022-10-20 22:03:45
 * @LastEditors: NyanCatda
 * @Description: 测试Socks5代理是否可用
 * @FilePath: \Cherino\Scan\Socks5\Test.go
 */
package Socks5

import (
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

/**
 * @description: 测试代理是否可用
 * @param {string} URL 代理地址
 * @return {bool} 是否可用
 */
func Test(URL string) bool {
	// 设置代理参数
	Dialer, err := proxy.SOCKS5("tcp", URL, nil, proxy.Direct)
	if err != nil {
		return false
	}
	httpTransport := &http.Transport{}
	httpTransport.Dial = Dialer.Dial

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
