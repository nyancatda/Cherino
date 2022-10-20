/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 17:54:52
 * @LastEditTime: 2022-10-20 19:10:24
 * @LastEditors: NyanCatda
 * @Description: 列出IP范围内的所有IP
 * @FilePath: \Cherino\Tools\IPList.go
 */
package Tools

import "fmt"

/**
 * @description: 根据输入的IP范围生成IP列表，没有验证输入的IP是否合法
 * @param {[]int} StartIP 起始IP
 * @param {[]int} EndIP 结束IP
 * @return {[]string} IP列表
 */
func IPList(StartIP []int, EndIP []int) []string {
	var IPList []string

	for true {
		if StartIP[3] == EndIP[3] && StartIP[2] == EndIP[2] && StartIP[1] == EndIP[1] && StartIP[0] == EndIP[0] {
			break
		}

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

		IPList = append(IPList, fmt.Sprintf("%d.%d.%d.%d", StartIP[0], StartIP[1], StartIP[2], StartIP[3]))
		StartIP[3]++
	}

	// 添加最后一个IP
	IPList = append(IPList, fmt.Sprintf("%d.%d.%d.%d", EndIP[0], EndIP[1], EndIP[2], EndIP[3]))

	return IPList
}
