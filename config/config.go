package config

import (
	"log"
	"strconv"
	"strings"
)

// ProxyServer 代理服务器
var ProxyServer string

// ProxyPort 代理端口
var ProxyPort uint16

// TimeFormat 格式化时间
var TimeFormat string = "2006-01-02 15:04:05"

// DebugLevel 调试级别
var DebugLevel int

const (
	// LevelShort 简短格式
	LevelShort int = iota
	// LevelLong 长格式日志
	LevelLong
	// LevelDebug 长日志 + 更多日志
	LevelDebug
)

// SetProxyServer 设置代理服务器
func SetProxyServer(gProxyServerSpec string) {
	if gProxyServerSpec == "" {
		return
	}
	tmp := strings.Split(gProxyServerSpec, ":")
	if len(tmp) == 2 {
		portInt, err := strconv.Atoi(tmp[1])
		if err == nil {
			ProxyServer = tmp[0]
			ProxyPort = uint16(portInt)
			log.Printf("Proxy server is %s:%d\n", ProxyServer, ProxyPort)
		} else {
			log.Printf("Set proxy port err %s\n", err.Error())
		}
	}
}

// SetDebugLevel 调试级别
func SetDebugLevel(gDebug int) {
	DebugLevel = gDebug
}
