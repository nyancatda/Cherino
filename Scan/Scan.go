/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 19:42:08
 * @LastEditTime: 2022-10-21 13:15:10
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
	"github.com/nyancatda/Cherino/Tools/Check"
	"github.com/nyancatda/Cherino/Tools/Flag"
	"github.com/nyancatda/Cherino/Tools/Pool"
)

/**
 * @description: 扫描可用代理
 * @param {string} ProxyType 代理类型，可选：socks4/socks5/http/https
 * @param {[]string} IPList IP列表
 * @param {int} StartPort 起始端口
 * @param {int} EndPort 结束端口
 * @param {func(ProxyType string, URL string)} StatusOK 代理可用时回调
 * @return {error} 错误信息
 */
func Proxy(ProxyType string, IPList []string, StartPort int, EndPort int, StatusOK func(ProxyType string, URL string)) error {
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
	err := Check.PortCheck(StartPort, EndPort)
	if err != nil {
		return err
	}

	// 开始扫描
	IPWaitGroup := Pool.NewPool(Flag.Pool / 2)
	IPWaitGroup.Add(1)
	PortWaitGroup := Pool.NewPool(Flag.Pool / 2)
	for true {
		if StartIP[3] > 255 {
			StartIP[3] = 0
			StartIP[2]++
		}
		if StartIP[2] > 255 {
			StartIP[2] = 0
			StartIP[1]++
		}
		if StartIP[1] > 225 {
			StartIP[1] = 0
			StartIP[0]++
		}

		go func(IP string) {
			defer IPWaitGroup.Done()
			for Port := StartPort; Port <= EndPort; Port++ {
				PortWaitGroup.Add(1)
				go func(IP string, Port int, WaitGroup *Pool.WaitGroup) {
					defer WaitGroup.Done()
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
				}(IP, Port, PortWaitGroup)
			}
			PortWaitGroup.Wait()
		}(fmt.Sprintf("%d.%d.%d.%d", StartIP[0], StartIP[1], StartIP[2], StartIP[3]))

		if StartIP[3] == EndIP[3] && StartIP[2] == EndIP[2] && StartIP[1] == EndIP[1] && StartIP[0] == EndIP[0] {
			break
		}

		StartIP[3]++
	}
	IPWaitGroup.Wait()

	return nil
}
