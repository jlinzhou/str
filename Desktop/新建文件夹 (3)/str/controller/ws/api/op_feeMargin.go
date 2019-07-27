package api

import(
	//"net/http"
	//"runtime"
	"github.com/astaxie/beego/logs"
	"strategy_middle/models"
	"strategy_middle/rabbitmq"
	"strategy_middle/setting"
	"strategy_middle/controller/interflow"
	//"strategy_middle/rabbitmq"
	//"gopkg.in/mgo.v2/bson"
	"time"
	//"Goroutine"
	//"os"
	//"bytes"
	"encoding/json"
	//"strconv"
	//"reflect"
	//"github.com/gin-gonic/gin"
	"strategy_middle/constant"
	//"time"
)

// type ModifyOverrideFunc interface {
	
// 	ChooseMethod()
// 	//修改保证金
// 	ModifyMarginOverride()
// 	//修改手续费
// 	ModifyFeeOverride()
// }

// type ModifyOverrideOp struct {
// 	Op string
// 	Message []uint8
// 	Ws *websocket.Conn
// 	Mt int
// }


// func (this ModifyOverrideOp) ChooseMethod(){
// 	if this.Op =="modifyMarginOverride"{
// 		this.ModifyMarginOverride()
// 	}else if this.Op =="modifyFeeOverride"{
// 		this.ModifyFeeOverride()
// 	}
// }

//直接创建个保证金记录
func (this *ApiOp) ModifyMarginOverride()(err error){

	//strategyId	symbolId	initialMargin	maintenanceMargin
	msg :=&struct{
		Id int `json:"id"`
		Op string `json:"op"`
		Args *struct{
			NodeId int `json:"nodeId"`
			StrategyId uint64 `json:"strategyId"`
			SymbolId string `json:"symbolId"`
			InitialMargin float64 `json:"initialMargin"`
			MaintenanceMargin float64 `json:"maintenanceMargin"`
		} `json:"args"`
	}{}

	err=json.Unmarshal(this.Message, msg)
	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1001,"fail",constant.Errmap[1001],err)
		return
	}	
	marginoverride:=models.StrategySymbolMarginOverride{
		NodeId:msg.Args.NodeId,
		StrategyId:msg.Args.StrategyId,
		SymbolId:msg.Args.SymbolId,
		InitialMargin:msg.Args.InitialMargin,
		MaintenanceMargin:msg.Args.MaintenanceMargin,
		CreateTime:time.Now(),
	}

	st := models.ConnecToDB("StrategySymbolMarginOverride")
	err =st.Insert(&marginoverride)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1500,"fail",constant.Errmap[1500],err)
    	return
	}

	thisargs,err:= interflow.StrategySymbolMarginOverride(msg.Args.NodeId,msg.Args.StrategyId)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1020,"fail",constant.Errmap[1020],err)
    	return
	}
	info:=&struct{
		Op string `json:"op"`
		Args models.StrategySymbolMarginOverride `json:"args"`
	}{
		"StrategySymbolMarginOverride",
		thisargs,
	}
	infojson,err:=json.Marshal(info)
	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1002,"fail",constant.Errmap[1002],err)
		return
	}

	rabbitmq.Send(setting.UpChannel, string(infojson))
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(2000,"fail",constant.Errmap[2000],err)
	}else{
		this.Send_callback_success()
	}
	return
}

func (this *ApiOp) ModifyFeeOverride()(err error){
	//strategyId	symbolId	initialMargin	maintenanceMargin
	msg :=&struct{
		Id int `json:"id"`
		Op string `json:"op"`
		Args *struct{
			NodeId int `json:"nodeId"`
			StrategyId uint64 `json:"strategyId"`
			SymbolId string `json:"symbolId"`
			Fee float64 `json:"fee"`
			FeeType string `json:"feeType"`
			FeeFormat string `json:"feeFormat"`

		} `json:"args"`
	}{}

	err=json.Unmarshal(this.Message, msg)
	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1001,"fail",constant.Errmap[1001],err)
		return
	}	

	feeoverride:=models.StrategySymbolFeeOverride{
		NodeId:msg.Args.NodeId,
		StrategyId:msg.Args.StrategyId,
		SymbolId:msg.Args.SymbolId,
		Fee:msg.Args.Fee,
		FeeType:msg.Args.FeeType,
		FeeFormat:msg.Args.FeeFormat,
		CreateTime:time.Now(),
	}

	st := models.ConnecToDB("StrategySymbolFeeOverride")
	err =st.Insert(&feeoverride)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1500,"fail",constant.Errmap[1500],err)
    	return
	}

	thisargs,err:= interflow.StrategySymbolFeeOverride(msg.Args.NodeId,msg.Args.StrategyId)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1020,"fail",constant.Errmap[1020],err)
    	return
	}
	info:=&struct{
		Op string `json:"op"`
		Args models.StrategySymbolFeeOverride `json:"args"`
	}{
		"StrategySymbolFeeOverride",
		thisargs,
	}
	infojson,err:=json.Marshal(info)
	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1002,"fail",constant.Errmap[1002],err)
		return
	}
	rabbitmq.Send(setting.UpChannel, string(infojson))
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(2000,"fail",constant.Errmap[2000],err)
	}else{
		this.Send_callback_success()
	}
	return
}