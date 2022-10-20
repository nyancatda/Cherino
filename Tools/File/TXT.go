/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 21:37:55
 * @LastEditTime: 2022-10-20 21:41:23
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \Cherino\Tools\File\TXT.go
 */
package File

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var TXTFileName = "%s_proxy.txt"

/**
 * @description: 追加写入TXT文件
 * @param {string} ProxyType 代理类型
 * @param {string} URL 代理地址
 * @return {error} 错误信息
 */
func WriteTXT(ProxyType string, URL string) error {
	FileName := fmt.Sprintf(JsonFileName, ProxyType)
	Path := filepath.Clean("./" + FileName)

	Content := URL + "\n"

	// 追加写入文件
	WriteFile, err := os.OpenFile(Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer WriteFile.Close()

	write := bufio.NewWriter(WriteFile)
	_, err = write.WriteString(Content)
	if err != nil {
		return err
	}
	if err := write.Flush(); err != nil {
		return err
	}

	return nil
}
