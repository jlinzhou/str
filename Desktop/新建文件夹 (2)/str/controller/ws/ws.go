package ws

import (
	"net/http"
	"runtime"
	"github.com/astaxie/beego/logs"
	"strategy_middle/controller/interflow"
	//"gopkg.in/mgo.v2/bson"
	//"time"
	//"Goroutine"
	//"os"
	"bytes"
	"encoding/json"
	"strconv"
	//"reflect"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strategy_middle/controller/ws/api"
	"strategy_middle/constant"
)




func parserWs(message []uint8,ws *websocket.Conn,mt int)(err error){

	msg :=&struct{
		Id int `json:"id"`
		Op string `json:"op"`
		Args interface{} `json:"args"`
	}{}
	
	err =json.Unmarshal(message, msg)

	if err!=nil{
		return
	}

	// if msg.Op =="auth"{
	// 	api.Send_init(message,ws,mt)
	// }
	caseApi := api.ApiOp{Op:msg.Op,Message:message,Ws:ws,Mt:mt}
	switch msg.Op {
	case constant.Allapi["auth"].Name:
		err=caseApi.Send_init()
	case constant.Allapi["createStrategy"].Name:
		err=caseApi.CreateStrategy()
	case constant.Allapi["startStrategy"].Name:
		err=caseApi.StartStrategy()
	case constant.Allapi["pauseStrategy"].Name:
		err=caseApi.PauseStrategy()
	case constant.Allapi["deleteStrategy"].Name:
		err=caseApi.DeleteStrategy()
	case constant.Allapi["modifyStrategy"].Name:
		err=caseApi.ModifyStrategy()
	case constant.Allapi["modifyMarginOverride"].Name:
		err=caseApi.ModifyMarginOverride()
	case constant.Allapi["modifyFeeOverride"].Name:
		err=caseApi.ModifyFeeOverride()
	default:
		//err=caseApi.Send_callback_fail(3001,"fail",constant.Errmap[3001],constant.Errmap[3001])
	}
	return
	// else if(msg.Op =="createStrategy" || msg.Op =="startStrategy" || 
	// 	msg.Op =="pauseStrategy" || msg.Op =="deleteStrategy" || msg.Op =="modifyStrategy"){
	// 	strategyop := api.StrategyOp{Op:msg.Op,Message:message,Ws:ws,Mt:mt}
	// 	err=strategyop.CreateStrategy()
	// 	//var strategyfunc  api.StrategyFunc
	// 	//strategyfunc =strategyop
	// 	//strategyfunc.ChooseMethod()
	// }
	// }else if(msg.Op =="modifyMarginOverride" || msg.Op =="modifyFeeOverride"){
	// 	modifyoverrideop := api.ModifyOverrideOp{Op:msg.Op,Message:message,Ws:ws,Mt:mt}
	// 	var modifyoverridefunc  api.ModifyOverrideFunc
	// 	modifyoverridefunc =modifyoverrideop
	// 	modifyoverridefunc.ChooseMethod()
	// }
	
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// Ping webSocket请求Ping 返回pong
func Ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}


	// t:=reflect.TypeOf(ws.RemoteAddr())
	// logs.Info(t)
	// for i:=0;i<t.NumMethod();i++{
		
	// 	m:=t.Method(i)
	// 	logs.Info(m.Name,m.Type)
	// }
	//ws.RemoteAddr().String()
	//interflow.Allws[ws] = 12345678
	//ws关闭的话就从rabbitmq.Allws中删除
	//defer  delete(interflow.Allws,ws)

	//logs.Info("users:         ",interflow.ClientMgr.GetAllUsers())

	defer ws.Close()
	

	for {
		// 读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			// 客户端关闭连接时也会进入
			logs.Warning(err,mt)
			//delete(interflow.Allws,ws)
			interflow.ClientMgr.DelClient(ws)
			break
		}
		// fmt.Printf("mt : %T", mt)
		// fmt.Println("mt=",mt)
		//发送到rabbitmq
		//
		//
		//
		//logs.Info("mt : %T", reflect.TypeOf(message))
		logs.Debug("ws_recv--",ws.RemoteAddr().String(),"--",string(message))

		err=parserWs(message,ws,mt)
		if err!=nil{
			errdata:=api.ApiOp{Op:"unknownPattern",Message:message,Ws:ws,Mt:mt}
			errdata.Send_callback_fail(3002,"fail",constant.Errmap[3002],constant.Errmap[3002])
		}

		// 

		
		// //分消息类型发送到rabbitmq
		// if msg.Op !="HeartBeat"{
		// 	//rabbitmq.Send("test1", string(message))
		// }


		// t:=reflect.TypeOf(msg)
		// v:=reflect.ValueOf(msg)
		// for i:=0;i<t.NumField();i++{
		// 	f:=t.Field(i)
		// 	val:=v.Field(i).Interface()
		// 	logs.Info(f.Type,val)
		// } 

		//logs.Info(msg.Strategy_name)
		//fmt.Println(msg)
		//fmt.Println(mt)
		//fmt.Println(message)
		//fmt.Printf("%T", message)
		//fmt.Println(GetGID())
		///////////////////////////////////////////////////////
		// var strategy_info models.Strategy_info
		// db := models.DB
		// //db.Where("strategy_name = ?", msg.Strategy_name).First(&strategy_info)
		// //fmt.Println("strategy_info=", strategy_info.Strategy_status)

		// var status string
		// if msg.Op == "start" {
		// 	status = "Proceeding"
		// } else if msg.Op == "stop" {
		// 	status = "Canceled"
		// }

		// if err := db.Model(&strategy_info).Where("strategy_name = ?", msg.Strategy_name).Update("update_time", time.Now()).Update("strategy_status", status).Error; err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	var st models.Strategy_info = models.Strategy_info{Strategy_name: msg.Strategy_name, Update_time: time.Now(), Strategy_status: status}
		// 	s, err := json.Marshal(st)
		// 	err = ws.WriteMessage(mt, []byte(s))
		// 	if err != nil {
		// 		break
		// 	}
		// }
		///////////////////////////////////////////////////////

		//var st Senddata = Senddata{Id: 1, Pong: "pong", Msg: "hello,girl!"}
		//s, err := json.Marshal(st)

		// 如果客户端发送ping就返回pong,否则数据原封不动返还给客户端
		//if string(message) == "ping" {
		//	message = []byte("pong")
		//}
		// 写入ws数据 二进制返回
		//err = ws.WriteMessage(mt, []byte(s))
		// 返回JSON字符串，借助gin的gin.H实现
		// v := gin.H{"message": msg}
		// err = ws.WriteJSON(v)
		//if err != nil {
		//	break
		//}
	}



}
