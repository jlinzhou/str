package main

import (
	"fmt"
	//"log"
	"net/http"
	//"strategy_middle/models"
	"strategy_middle/route"
	"strategy_middle/setting"
	"strategy_middle/rabbitmq"
	_"strategy_middle/logging"
	//"github.com/astaxie/beego/logs"
	//"time"
	//"strategy_middle/models"
	//"encoding/json"
	//"strategy_middle/logging"

	//"time"
)



// //var Allws []*websocket.Conn

// func running() {
// 	var times int
// 	// 构建一个无限循环
// 	for {
// 		times++
// 		fmt.Println("tick", times)
// 		// 延时1秒
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	
	//go rabbitmq.Recv()

	//route.Init()

	//	for ws:=range route.Allws{

	//	}
	//
	// c := models.ConnecToDB("student2")

	// students := make([]models.Student, 20)
	// err := c.Find(nil).All(&students)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(students) 

	// //rabbitmq接收channel为test1的信息
	go rabbitmq.Recv(setting.DownChannel)

	// //每隔两秒发送一次到rabbitmq的心跳
	// hbdata:=`{"op":"Heartbeat","args":{"id":1}}`
	// go func(){
	// 	for{
	// 		rabbitmq.Send("test1", string(hbdata))
	// 		time.Sleep(10*time.Second)
	// 	}
	// }()


	router := route.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort), //监听的TCP地址
		Handler:        router,                               //http句柄
		ReadTimeout:    setting.ReadTimeout,                  //允许读取的最大时间
		WriteTimeout:   setting.WriteTimeout,                 //允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                              //请求头的最大字节数
	}
	s.ListenAndServe()
}
