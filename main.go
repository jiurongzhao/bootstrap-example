package main

import (
	"fmt"

	_ "github.com/jiurongzhao/bootstrap-config-yaml/yaml"
	"github.com/jiurongzhao/bootstrap-global/config"
	"github.com/jiurongzhao/bootstrap-global/log"
	_ "github.com/jiurongzhao/bootstrap-log-logrus/logrus"
)

type Author struct {
	Name string `name:"name"`
	Role string `name:"role"`
}

type Bootstrap struct {
	Name   string `name:"name"`
	Author Author `name:"author"`
}

func main() {
	// 初始化配置
	if err := config.InitGlobalInstance("yaml", "resource/app.yaml"); err != nil {
		fmt.Errorf(err.Error())
		return
	}
	// 初始化日志
	if err := log.InitGlobalLogger("logrus"); err != nil {
		fmt.Errorf(err.Error())
		return
	}

	// 通过键获取配置
	bootstrapName, ok := config.Get("bootstrap.name")
	if ok {
		log.Info("bootstrap.name:", bootstrapName)
	} else {
		log.Warn("not found bootstrap.name")
	}
	// 解析配置到结构体
	var bootstrap Bootstrap
	if err := config.Resolve("bootstrap", &bootstrap); err != nil {
		log.Error(err.Error())
	} else {
		log.Info("bootstrap:", bootstrap)
	}

	var author Author
	if err := config.Resolve("bootstrap.author", &author); err != nil {
		log.Error(err.Error())
	} else {
		log.Info("bootstrap.author:", author)
	}
}
