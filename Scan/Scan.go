/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 19:42:08
 * @LastEditTime: 2022-10-20 23:17:44
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
	"github.com/nyancatda/Cherino/Tools/Flag"
	"github.com/nyancatda/Cherino/Tools/Pool"
)

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

	// 计算线程分配
	var IPMaxPool int
	if len(IPList) > Flag.Pool/2 {
		IPMaxPool = Flag.Pool / 2
	} else {
		IPMaxPool = len(IPList)
	}
	PortMaxPool := Flag.Pool - IPMaxPool

	// 开始扫描
	IPWaitGroup := Pool.NewPool(IPMaxPool)
	IPWaitGroup.Add(1)
	PortWaitGroup := Pool.NewPool(PortMaxPool)
	for _, IP := range IPList {
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
		}(IP)
	}
	IPWaitGroup.Wait()

	return nil
}
