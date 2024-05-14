package main

import (
	"docker-pull/config"
	"fmt"
	"github.com/jinzhu/configor"
)

//var HarborIP = struct {
//	HarborIP string `yaml:"harborIP"`
//
//	Auth struct {
//		Username string `yaml:"username"`
//		Password string `yaml:"password"`
//	}
//}{}

// HarborIP 解析配置文件
var HarborIP = config.HarborIP{}

func main() {
	configor.Load(&HarborIP, "config.yml")
	ip := HarborIP.HarborIP
	username := HarborIP.Auth.Username
	passwd := HarborIP.Auth.Password
	project := HarborIP.Project
	fmt.Println("IP:", ip)
	fmt.Println("username:", username)
	fmt.Println("passwd:", passwd)
	fmt.Println("project:", project)
	for _, i := range project {
		fmt.Println(i)
	}
}
