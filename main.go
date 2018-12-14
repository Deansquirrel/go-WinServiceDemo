package main

import (
	"github.com/BurntSushi/toml"
	"github.com/Deansquirrel/go-tool"
	"github.com/kardianos/service"
	"log"
	"os"
	"winServiceDemo/global"
	"winServiceDemo/worker"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	//服务所执行的代码
	w := worker.Worker{}
	w.Do()
}

func (p *program) Stop(s service.Service) error {
	return nil
}

var sysConfig global.SysConfig

func main() {

	go_tool.Log("程序启动")
	defer go_tool.Log("程序退出")

	dir, err := go_tool.GetCurrPath()
	if err != nil {
		dir = ""
	} else {
		dir = dir + "\\"
	}
	configPath := dir + "config"

	_, err = toml.DecodeFile(configPath, &sysConfig)
	if err != nil {
		msg := "加载配置文件失败"
		log.Println(msg)
		go_tool.Log(msg)
		return
	}

	svcConfig := &service.Config{
		Name:        sysConfig.ServiceConfig.Name,
		DisplayName: sysConfig.ServiceConfig.DisplayName,
		Description: sysConfig.ServiceConfig.Description,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err = s.Install()
			msg := ""
			if err != nil {
				msg = err.Error()
			} else {
				msg = "服务安装成功"
			}
			log.Println(msg)
			go_tool.Log(msg)
			return
		case "uninstall":
			err = s.Uninstall()
			msg := ""
			if err != nil {
				msg = err.Error()
			} else {
				msg = "服务卸载成功"
			}
			log.Println(msg)
			go_tool.Log(msg)
			return
		default:
			log.Println("未识别的参数名称\n安装服务:install\n卸载服务:uninstall")
			return
		}
	} else {
		err = s.Run()
		if err != nil {
			go_tool.Log(err.Error())
			log.Fatal(err)
		}
	}

}
