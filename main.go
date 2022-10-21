/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 17:06:29
 * @LastEditTime: 2022-10-21 12:55:00
 * @LastEditors: NyanCatda
 * @Description: 主文件
 * @FilePath: \Cherino\main.go
 */
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/Cherino/Scan"
	"github.com/nyancatda/Cherino/Tools"
	"github.com/nyancatda/Cherino/Tools/File"
	"github.com/nyancatda/Cherino/Tools/Flag"
)

func main() {
	PrintStartupScreen()
	// 设置日志参数
	AyaLog.LogLevel = AyaLog.INFO
	AyaLog.LogWriteFile = false

	// 解析命令行参数
	err := Flag.Get()
	if err != nil {
		AyaLog.Error("System", err)
		return
	}

	var WriteFile = func(ProxyType string, URL string) {
		err := File.Write(ProxyType, URL)
		if err != nil {
			AyaLog.Error("System", err)
			os.Exit(1)
		}

		AyaLog.Info("Scan", URL+" 是一个可用的"+ProxyType+"代理, 已写入文件")
	}

	AyaLog.Info("System", "起始IP: "+fmt.Sprintf("%d.%d.%d.%d", Flag.StartIP[0], Flag.StartIP[1], Flag.StartIP[2], Flag.StartIP[3]))
	AyaLog.Info("System", "结束IP: "+fmt.Sprintf("%d.%d.%d.%d", Flag.EndIP[0], Flag.EndIP[1], Flag.EndIP[2], Flag.EndIP[3]))
	AyaLog.Info("System", "起始端口: "+strconv.Itoa(Flag.StartPort))
	AyaLog.Info("System", "结束端口: "+strconv.Itoa(Flag.EndPort))
	AyaLog.Info("System", "最大线程数: "+strconv.Itoa(Flag.Pool))
	AyaLog.Info("System", "保存类型: "+Flag.SaveType)

	// 生成IP列表
	AyaLog.Info("System", "正在生成IP列表......")
	IPList := Tools.IPList(Flag.StartIP, Flag.EndIP)
	AyaLog.Info("System", "生成IP列表完成, 共计 "+strconv.Itoa(len(IPList))+" 个IP")

	// 扫描可用代理
	if Flag.Socks5 {
		AyaLog.Info("System", "Socks5代理扫描开始")
		err := Scan.Proxy("socks5", IPList, Flag.StartPort, Flag.EndPort, WriteFile)
		if err != nil {
			AyaLog.Error("Scan", err)
			return
		}
	}
	if Flag.Socks4 {
		AyaLog.Info("System", "Socks4代理扫描开始")
		err := Scan.Proxy("socks4", IPList, Flag.StartPort, Flag.EndPort, WriteFile)
		if err != nil {
			AyaLog.Error("Scan", err)
			return
		}
	}
	if Flag.HTTP {
		AyaLog.Info("System", "HTTP代理扫描开始")
		err := Scan.Proxy("http", IPList, Flag.StartPort, Flag.EndPort, WriteFile)
		if err != nil {
			AyaLog.Error("Scan", err)
			return
		}
	}
	if Flag.HTTPS {
		AyaLog.Info("System", "HTTPS代理扫描开始")
		err := Scan.Proxy("https", IPList, Flag.StartPort, Flag.EndPort, WriteFile)
		if err != nil {
			AyaLog.Error("Scan", err)
			return
		}
	}

	os.Exit(0)
}

func PrintStartupScreen() {
	fmt.Println(`
 _____ _               _             
/ ____| |             (_)            
| |    | |__   ___ _ __ _ _ __   ___  
| |    | '_ \ / _ \ '__| | '_ \ / _ \ 
| |____| | | |  __/ |  | | | | | (_) |
 \_____|_| |_|\___|_|  |_|_| |_|\___/ 											
	`)
}
