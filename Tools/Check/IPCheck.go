/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 17:40:43
 * @LastEditTime: 2022-10-20 18:23:11
 * @LastEditors: NyanCatda
 * @Description: IP检查
 * @FilePath: \Cherino\Tools\IPCheck.go
 */
package Check

import (
	"errors"
)

/**
 * @description: 检查IP输入是否正确
 * @param {[]int} StartIP 开始端
 * @param {[]int} EndIP 结束端
 * @return {error} 错误信息
 */
func IPCheck(StartIP []int, EndIP []int) error {
	// 判断切片长度是否正确
	if len(StartIP) != 4 || len(EndIP) != 4 {
		return errors.New("IP length error")
	}

	// 判断是否IP合法
	for _, IPInt := range StartIP {
		if IPInt < 0 || IPInt > 255 {
			return errors.New("Start IP is invalid")
		}
	}
	for _, IPInt := range EndIP {
		if IPInt < 0 || IPInt > 255 {
			return errors.New("End IP is invalid")
		}
	}

	// 判断IP范围是否正确
	for StartIPIndex, StartIPValue := range StartIP {
		for _, EndIPValue := range EndIP {
			if StartIPIndex == 0 {
				if EndIPValue < StartIPValue {
					return errors.New("wrong IP range")
				}
				if EndIPValue >= StartIPValue {
					break
				}
			}
			if StartIPIndex == 1 {
				if EndIPValue < StartIPValue {
					return errors.New("wrong IP range")
				}
				if EndIPValue >= StartIPValue {
					break
				}
			}
			if StartIPIndex == 2 {
				if EndIPValue < StartIPValue {
					return errors.New("wrong IP range")
				}
				if EndIPValue > StartIPValue {
					break
				}
			}
			if StartIPIndex == 3 {
				if EndIPValue < StartIPValue {
					return errors.New("wrong IP range")
				}
				if EndIPValue >= StartIPValue {
					break
				}
			}
		}
	}

	return nil
}
