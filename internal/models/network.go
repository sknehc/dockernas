package models

import "log"

type NetworkInfo struct {
	IP                 string            `json:"ip"`
	Domain             string            `json:"domain"`
	DNSServerEnable    bool              `json:"DNSServerEnable"`
	DNSServerLoading   bool              `json:"DNSServerLoading"`
	HttpGatewayEnable  bool              `json:"httpGatewayEnable"`
	HttpGatewayLoading bool              `json:"httpGatewayLoading"`
	HttpProxyConfigs   []HttpProxyConfig `json:"httpProxyConfigs"`
}

type HttpProxyConfig struct {
	HostName     string `json:"hostName"  gorm:"unique;not null"`
	InstanceName string `json:"instanceName"`
	Port         string `json:"port"`
	CreateTime   int64  `json:"createTime"`
}

type InstanceHttpPorts struct {
	InstanceName string   `json:"instanceName"`
	Ports        []string `json:"ports"`
}

func AddHttpProxyConfig(proxyConfig *HttpProxyConfig) {
	err := GetDb().Create(proxyConfig).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func DelHttpProxyConfig(proxyConfig *HttpProxyConfig) {
	err := GetDb().Where("host_name = ?", proxyConfig.HostName).Delete(proxyConfig).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func GetHttpProxyConfig() []HttpProxyConfig {
	var configs []HttpProxyConfig
	err := GetDb().Find(&configs).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return configs
}

func GetHttpProxyConfigByHostName(hostName string) HttpProxyConfig {
	var config HttpProxyConfig
	err := GetDb().First(&config, "host_name=?", hostName).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return config
}
