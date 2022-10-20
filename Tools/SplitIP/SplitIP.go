/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 21:13:29
 * @LastEditTime: 2022-10-20 21:18:38
 * @LastEditors: NyanCatda
 * @Description: 分割IP
 * @FilePath: \Cherino\Tools\SplitIP\SplitIP.go
 */
package SplitIP

import (
	"errors"
	"strconv"
	"strings"
)

/**
 * @description: 字符串转换为Int数组IP
 * @param {string} IP IP地址
 * @return {[]int} IP地址数组
 * @return {error} 错误信息
 */
func StrToArrayInt(IP string) ([]int, error) {
	var IPArray []int

	IPArrayStr := strings.Split(IP, ".")

	if len(IPArrayStr) != 4 {
		return nil, errors.New("IP is invalid")
	}

	for _, Value := range IPArrayStr {
		// 转化为int
		IPInt, err := strconv.Atoi(Value)
		if err != nil {
			return nil, errors.New("IP is invalid")
		}

		IPArray = append(IPArray, IPInt)
	}

	return IPArray, nil
}
