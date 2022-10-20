# Cherino
基于Golang编写的网络代理扫描工具

支持扫描IPv4任意网段下Socks4/Socks5/HTTP/HTTPS协议的网络代理
使用Golang编写，支持多线程扫描

⚠ 此工具仅用于学习，请不要滥用此工具，使用此工具造成的后果由您自行承担 ⚠

## ⚙️ 可用参数
```
-start_ip string
    起始IP (default "1.1.1.1")
-end_ip string
    结束IP (default "255.255.255.255")
-start_port int
    起始端口 (default 1)
-end_port int
    结束端口 (default 65535)
-socks4
    获取Socks4代理
-socks5
    获取Socks5代理
-http
    获取HTTP代理
-https
    获取HTTPS代理
-pool int
    最大线程数量 (default 500)
-timeout int
    测试超时时间，单位秒 (default 2)
-save_type string
    保存方式，可选：json/txt (default "json")
```

## 🎬 开始扫描
``` shell
Cherino.exe -socks4 -socks5 -http -https -pool 10000
```

## 🛠️ 构建
自行构建前需要拥有 Go >= 1.18
### 克隆仓库
```
git clone https://github.com/nyancatda/Cherino.git
```
### 编译项目
```
# 获取依赖包
go mod tidy

# 开始编译
go build .
```

### 📖 许可证
项目采用`Mozilla Public License Version 2.0`协议开源

二次修改源代码需要开源修改后的代码，对源代码修改之处需要提供说明文档