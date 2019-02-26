package config

import (
	"github.com/Deansquirrel/goToolCommon"
	"strings"
)

type SysConfig struct {
	Total         Total         `toml:"total"`
	ServiceConfig serviceConfig `toml:"serviceConfig"`
}

type serviceConfig struct {
	Name        string `toml:"name"`
	DisplayName string `toml:"displayName"`
	Description string `toml:"description"`
}

//返回配置字符串
func (sc *SysConfig) GetConfigStr() (string, error) {
	return goToolCommon.GetJsonStr(sc)
}

//配置检查并格式化
func (sc *SysConfig) FormatConfig() {
	sc.Total.FormatConfig()
	sc.ServiceConfig.FormatConfig()
}

//格式化
func (sc *serviceConfig) FormatConfig() {
	sc.Name = strings.Trim(sc.Name, " ")
	sc.DisplayName = strings.Trim(sc.DisplayName, " ")
	sc.Description = strings.Trim(sc.Description, " ")
	if sc.Name == "" {
		sc.Name = "goWinServiceDemo"
	}
	if sc.DisplayName == "" {
		sc.DisplayName = "goWinServiceDemo"
	}
	if sc.Description == "" {
		sc.Description = sc.Name
	}
}
