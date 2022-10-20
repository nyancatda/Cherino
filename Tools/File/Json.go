/*
 * @Author: NyanCatda
 * @Date: 2022-10-20 21:23:42
 * @LastEditTime: 2022-10-20 22:16:37
 * @LastEditors: NyanCatda
 * @Description: 写入Json文件
 * @FilePath: \Cherino\Tools\File\Json.go
 */
package File

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var JsonFileName = "%s_proxy.json"

/**
 * @description: 追加写入Json文件
 * @param {string} ProxyType 代理类型
 * @param {string} URL 代理地址
 * @return {error} 错误信息
 */
func WriteJson(ProxyType string, URL string) error {
	FileName := fmt.Sprintf(JsonFileName, ProxyType)

	var ProxyList []string

	Path := filepath.Clean("./" + FileName)

	// 读取文件内容
	ReadFile, err := os.OpenFile(Path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer ReadFile.Close()
	Content, err := ioutil.ReadAll(ReadFile)
	if err != nil {
		return err
	}
	if len(Content) == 0 {
		Content = []byte("[]")
	}
	if err := json.Unmarshal(Content, &ProxyList); err != nil {
		return err
	}

	// 追加内容
	ProxyList = append(ProxyList, URL)

	// 写入文件
	WriteFile, err := os.OpenFile(Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer WriteFile.Close()
	// 转换为Json
	ProxyListContent, err := json.Marshal(ProxyList)
	if err != nil {
		return err
	}
	n, _ := WriteFile.Seek(0, os.SEEK_END)
	_, err = WriteFile.WriteAt(ProxyListContent, n)
	if err != nil {
		return err
	}

	return nil
}
