package main

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

// IP 解析配置
type IP struct {
	IPAddress string `yaml:"harbor_ip"`
}

// UserPass 用户名密码解析
type UserPass struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Auth struct {
	Auth UserPass `yaml:"auth"`
}

// ReadFile 读取配置文件
func ReadFile(file string) (data map[string]any, err error) {
	// 打开文件
	dataBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.New("配置文件读取失败")
	}

	ip := IP{}
	// 解析yaml文件
	err = yaml.Unmarshal(dataBytes, &ip)
	if err != nil {
		return nil, errors.New("yml文件解析失败")
	}

	mp := make(map[string]any, 10)
	err = yaml.Unmarshal(dataBytes, mp)
	if err != nil {
		return nil, errors.New("yml文件用户名解析失败")
	}
	// Harborip 解析yaml文件字段
	//var Harborip = ip.IPAddress
	return mp, nil
	// 解析用户密码

}
