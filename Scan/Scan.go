/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 19:42:08
 * @LastEditTime: 2022-10-20 19:50:33
 * @LastEditors: NyanCatda
 * @Description: 扫描可用代理
 * @FilePath: \Cherino\Scan\Scan.go
 */
package Scan

import (
	"errors"
	"fmt"

	Socks5Proxy "github.com/nyancatda/Cherino/Scan/Socks5"
	"github.com/nyancatda/Cherino/Tools"
	"github.com/nyancatda/Cherino/Tools/Check"
	"github.com/nyancatda/Cherino/Tools/Pool"
)

// 最大线程数
var MaxPool = 500

/**
 * @description: 扫描可用代理
 * @param {string} ProxyType 代理类型，可选：socks5
 * @param {[]int} StartIP 起始IP
 * @param {[]int} EndIP 结束IP
 * @param {int} StartPort 起始端口
 * @param {int} EndPort 结束端口
 * @return {[]string} 可用代理列表
 * @return {error} 错误信息
 */
func Proxy(ProxyType string, StartIP []int, EndIP []int, StartPort int, EndPort int) ([]string, error) {
	// 检查输入
	switch ProxyType {
	case "socks5":
		break
	default:
		return nil, errors.New("ProxyType is invalid")
	}

	err := Check.IPCheck(StartIP, EndIP)
	if err != nil {
		return nil, err
	}
	err = Check.PortCheck(StartPort, EndPort)
	if err != nil {
		return nil, err
	}

	// 生成IP列表
	IPList := Tools.IPList(StartIP, EndIP)

	// 循环扫描端口范围内的代理
	var OKProxyList []string

	WaitGroup := Pool.NewPool(MaxPool)

	Request := func(IP string, Port int) {
		fmt.Println(1)
		URL := IP + ":" + fmt.Sprintf("%d", Port)

		switch ProxyType {
		case "socks5":
			if Socks5Proxy.Test(URL) {
				OKProxyList = append(OKProxyList, URL)
			}
		}

		WaitGroup.Done()
	}

	for _, IP := range IPList {
		for Port := StartPort; Port <= EndPort; Port++ {
			WaitGroup.Add(1)
			go Request(IP, Port)
		}
	}

	return OKProxyList, nil
}
