/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 21:22:22
 * @LastEditTime: 2022-10-20 21:44:09
 * @LastEditors: NyanCatda
 * @Description: 写入文件
 * @FilePath: \Cherino\Tools\File\File.go
 */
package File

import (
	"errors"

	"github.com/nyancatda/Cherino/Tools/Flag"
)

/**
 * @description: 代理写入文件
 * @param {string} ProxyType 代理类型
 * @param {string} URL 代理地址
 * @return {error} 错误信息
 */
func Write(ProxyType string, URL string) error {
	switch Flag.SaveType {
	case "json":
		return WriteJson(ProxyType, URL)
	case "txt":
		return WriteTXT(ProxyType, URL)
	default:
		return errors.New("SaveType Error")
	}
}
