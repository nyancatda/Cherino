/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 17:51:46
 * @LastEditTime: 2022-10-20 19:27:18
 * @LastEditors: NyanCatda
 * @Description: 端口检查
 * @FilePath: \Cherino\Tools\Check\PortCheck.go
 */
package Check

import "errors"

/**
 * @description: 端口检查
 * @param {int} StartPort 端口开始
 * @param {int} EndPort 端口结束
 * @return {error} 错误信息
 */
func PortCheck(StartPort int, EndPort int) error {
	if StartPort <= 0 || StartPort > 65535 {
		return errors.New("Start Port is invalid")
	}
	if EndPort <= 0 || EndPort > 65535 {
		return errors.New("End Port is invalid")
	}

	if StartPort > EndPort {
		return errors.New("wrong Port range")
	}

	return nil
}
