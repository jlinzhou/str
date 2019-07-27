package interflow
import(
	"strategy_middle/models"
	//"github.com/gorilla/websocket"
	"github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2/bson"
	//"encoding/json"
)


//var Allws = make(map[*websocket.Conn]uint64)

//收到数据进行解析，然后根据情况发到前端或者后台或者存到数据库
//
//发到前端的
//合约pnl
//账户pnl
//报单回报
//仓位变化
//



//发送到前端，根据策略id找出对应的账户id，根据map表里面的账户id发送到对应的客户端
func SendWs(strategyid uint64,message string){

	st:=models.ConnecToDB("StrategyInfo")
	strategyinfo :=models.StrategyInfo{}
	err := st.Find(bson.M{"strategyId":  strategyid}).One(strategyinfo)
	if err!=nil{
		logs.Error(err)
	}

	// for ws,accoundid:= range Allws{
	// 	// if len(Allws)>0{
	// 	if accoundid ==strategyinfo.AccountId{
	// 		err = ws.WriteMessage(1, []byte(message))
	// 		if err!=nil{
	// 			logs.Info(err)
	// 		}
	// 	}
	// }
			
}

//接入成交回报，计算仓位，发送到前端
func calculatePosition(){

}


//接入实时行情，计算pnl，发送到前端
func calculatePnl(){

}



