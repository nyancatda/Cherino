/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 20:57:58
 * @LastEditTime: 2022-10-20 22:10:17
 * @LastEditors: NyanCatda
 * @Description: 获取参数
 * @FilePath: \Cherino\Tools\Flag\Flag.go
 */
package Flag

import (
	"errors"
	"flag"
)

var StartIP string
var EndIP string
var StartPort int
var EndPort int
var Socks5 bool
var Socks4 bool
var HTTP bool
var HTTPS bool
var Pool int
var TimeOut int64
var SaveType string

/**
 * @description: 获取参数
 * @return {error} 错误信息
 */
func Get() error {
	// 获取参数
	FlagStartIP := flag.String("start_ip", "1.1.1.1", "起始IP")
	FlagEndIP := flag.String("end_ip", "255.255.255.255", "结束IP")

	FlagStartPort := flag.Int("start_port", 1, "起始端口")
	FlagEndPort := flag.Int("end_port", 65535, "结束端口")

	FlagSocks5 := flag.Bool("socks5", false, "获取Socks5代理")
	FlagSocks4 := flag.Bool("socks4", false, "获取Socks4代理")
	FlagHTTP := flag.Bool("http", false, "获取HTTP代理")
	FlagHTTPS := flag.Bool("https", false, "获取HTTPS代理")

	FlagPool := flag.Int("pool", 500, "最大线程数量")

	FlagTimeOut := flag.Int64("timeout", 2, "测试超时时间，单位秒")

	FlagSaveType := flag.String("save_type", "json", "保存方式，可选：json/txt")
	flag.Parse()

	// 检查参数
	if !*FlagSocks5 && !*FlagSocks4 && !*FlagHTTP && !*FlagHTTPS {
		return errors.New("必须指定至少一个代理类型")
	}
	if *FlagPool < 1 {
		return errors.New("最大线程数量不能小于1")
	}
	switch *FlagSaveType {
	case "json":
		break
	case "txt":
		break
	default:
		return errors.New("保存方式不可用")
	}

	// 参数写入变量
	StartIP = *FlagStartIP
	EndIP = *FlagEndIP
	StartPort = *FlagStartPort
	EndPort = *FlagEndPort
	Socks5 = *FlagSocks5
	Socks4 = *FlagSocks4
	HTTP = *FlagHTTP
	HTTPS = *FlagHTTPS
	Pool = *FlagPool
	TimeOut = *FlagTimeOut
	SaveType = *FlagSaveType

	return nil
}
