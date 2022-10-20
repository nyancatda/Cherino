/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 19:55:23
 * @LastEditTime: 2022-10-20 22:03:55
 * @LastEditors: NyanCatda
 * @Description: 测试Socks4代理是否可用
 * @FilePath: \Cherino\Scan\Socks4\Test.go
 */
package Socks4

import (
	"net/http"
	"time"

	"h12.io/socks"
)

/**
 * @description: 测试代理是否可用
 * @param {string} URL 代理地址
 * @return {bool} 是否可用
 */
func Test(URL string) bool {
	// 设置代理参数
	Dialer := socks.Dial("socks4://" + URL)

	httpTransport := &http.Transport{}
	httpTransport.Dial = Dialer

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
