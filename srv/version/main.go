package main

import (
	"fmt"
	"github.com/micro/go-log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/etcdv3"

	"github.com/gopperin/sme-ray/srv/version/handler"
	"github.com/gopperin/sme-ray/srv/version/persist"
	myproto "github.com/gopperin/sme-ray/srv/version/proto" // import proto生成的类
)

func main() {

	err := persist.GMariadb.Init()
	if err != nil {
		fmt.Println("*** mariadb error : ", err.Error())
		return
	}
	fmt.Println("====== mariadb init ======")

	versionService()
}

func versionService() {

	service := micro.NewService(
		micro.Name("snc.gc.srv.version"), // 定义service的名称为version
		micro.Version("1.0.0"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init() // 初始化service

	// 注册服务
	myproto.RegisterVersionHandler(service.Server(), new(handler.Version))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
