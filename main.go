package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"bubble/setting"
	"fmt"
	"os"
)
//编译: go build 执行: bubble.exe conf/config.ini
func main() {
	//command := exec.Command("/bubble.exe conf/config.ini")
	//err := command.Run()
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Usage：./bubble conf/config.ini")
		return
	}

	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接

	// 模型绑定-将 model和数据库的表建立关联
	dao.DB.AutoMigrate(&models.Todo{})

	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}

}
