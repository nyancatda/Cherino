/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 19:42:08
 * @LastEditTime: 2022-10-20 22:12:57
 * @LastEditors: NyanCatda
 * @Description: 扫描可用代理
 * @FilePath: \Cherino\Scan\Scan.go
 */
package Scan

import (
	"errors"
	"fmt"

	HTTPProxy "github.com/nyancatda/Cherino/Scan/HTTP"
	HTTPSProxy "github.com/nyancatda/Cherino/Scan/HTTPS"
	Socks4Proxy "github.com/nyancatda/Cherino/Scan/Socks4"
	Socks5Proxy "github.com/nyancatda/Cherino/Scan/Socks5"
	"github.com/nyancatda/Cherino/Tools"
	"github.com/nyancatda/Cherino/Tools/Check"
	"github.com/nyancatda/Cherino/Tools/Pool"
)

// 最大线程数
var MaxPool = 500

/**
 * @description: 扫描可用代理
 * @param {string} ProxyType 代理类型，可选：socks4/socks5/http/https
 * @param {[]int} StartIP 起始IP
 * @param {[]int} EndIP 结束IP
 * @param {int} StartPort 起始端口
 * @param {int} EndPort 结束端口
 * @param {func(ProxyType string, URL string)} StatusOK 代理可用时回调
 * @return {error} 错误信息
 */
func Proxy(ProxyType string, StartIP []int, EndIP []int, StartPort int, EndPort int, StatusOK func(ProxyType string, URL string)) error {
	// 检查输入
	switch ProxyType {
	case "socks5":
		break
	case "socks4":
		break
	case "http":
		break
	case "https":
		break
	default:
		return errors.New("ProxyType is invalid")
	}

	err := Check.IPCheck(StartIP, EndIP)
	if err != nil {
		return err
	}
	err = Check.PortCheck(StartPort, EndPort)
	if err != nil {
		return err
	}

	// 生成IP列表
	IPList := Tools.IPList(StartIP, EndIP)

	// 循环扫描端口范围内的代理
	WaitGroup := Pool.NewPool(MaxPool)

	Request := func(IP string, Port int) {
		URL := IP + ":" + fmt.Sprintf("%d", Port)

		var ProxStatus bool
		switch ProxyType {
		case "socks5":
			ProxStatus = Socks5Proxy.Test(URL)
		case "socks4":
			ProxStatus = Socks4Proxy.Test(URL)
		case "http":
			ProxStatus = HTTPProxy.Test(URL)
		case "https":
			ProxStatus = HTTPSProxy.Test(URL)
		}

		if ProxStatus {
			StatusOK(ProxyType, URL)
		}

		WaitGroup.Done()
	}

	for _, IP := range IPList {
		for Port := StartPort; Port <= EndPort; Port++ {
			WaitGroup.Add(1)
			go Request(IP, Port)
		}
	}

	WaitGroup.Wait()

	return nil
}
