package rabbitmq
import (
	"strategy_middle/controller/interflow"
	"github.com/astaxie/beego/logs"
	"strategy_middle/models"
	"strategy_middle/constant"
	"encoding/json"
	"strategy_middle/setting"
	"gopkg.in/mgo.v2/bson"
)



//创建策略初始化
//策略初始化流程，收到策略创建信息后，先发送SymbolInfo，收到StrategyInitReq后发送StrategyInfo、
//StrategyParam等信息，最后结束发送StrategyInitFinished信息
func RecvDownInfo(data string){

	pattern:=&struct{
		Op string `json:"op"`
		Args interface{} `json:"args"`
	}{}

	err:=json.Unmarshal([]byte(data),pattern)
	if err!=nil{
		logs.Error(err)
	}

	
	switch pattern.Op {

	case constant.RabbitApi["Init"]:
		systeminit(data)
	case constant.RabbitApi["StrategyHeartbeat"]:
		
	case constant.RabbitApi["StrategyInitReq"]:
		strategyinit(data)
	case constant.RabbitApi["LogInfo"]:
	case constant.RabbitApi["TdInfo"]:
	case constant.RabbitApi["MdInfo"]:
	case constant.RabbitApi["AccountInfo"]:
	case constant.RabbitApi["AccountPnlDaily"]:
	case constant.RabbitApi["StrategyPnlDaily"]:
	}
	//"StrategyHeartbeat":"StrategyHeartbeat",
	/*
		"StrategyHeartbeat":"StrategyHeartbeat",
	"LogInfo":"LogInfo",
	"NodeInitConfig":"NodeInitConfig",
	"TdInfo":"TdInfo",
	"MdInfo":"MdInfo",
	"SymbolTradingPeriod":"SymbolTradingPeriod",
	"SymbolMarginDefault":"SymbolMarginDefault",
	"SymbolFeeDefault":"SymbolFeeDefault",
	"AccountInfo":"AccountInfo",
	"AccountPnlDaily":"AccountPnlDaily",
	"StrategySymbol":"StrategySymbol",
	"StrategySymbolMarginOverride":"StrategySymbolMarginOverride",
	"StrategySymbolFeeOverride":"StrategySymbolFeeOverride",
	"StrategyPnlDaily":"StrategyPnlDaily",
	"StrategyArchive":"StrategyArchive",
	*/

	//要收到的api
	/*
	StrategyHeartbeat
	StrategyInitReq
	LogInfo
	TdInfo
	MdInfo
	AccountInfo
	AccountPnlDaily
	StrategyPnlDaily
	*/

	//要发送的api
	/*
	SymbolInfo
	StrategyInfo
	StrategyParam
	StrategyInitFinished
	StrategySwitch
	UpdateStrategyParams
	StrategyHeartbeat

	NodeInitConfig
	SymbolTradingPeriod
	SymbolMarginDefault
	StrategySymbol
	StrategySymbolMarginOverride
	StrategySymbolFeeOverride
	*/

}
//收到后台初始化请求Init 如果stage:0，就发送stage:1，如果stage:2，就发送初始化数据
func systeminit(data string){

	init:=&struct{
		Op string `json:"op"`
		Args *struct{
			Stage int `json:"stage"`
		}`json:"args"`
	}{}

	err :=json.Unmarshal([]byte(data),init)
	if err!=nil{
		logs.Error(err)
	}
	if init.Args.Stage==0{
		sendinit:=&struct{
			Op string `json:"op"`
			Args *struct{
				Stage int `json:"stage"`
			}`json:"args"`
		}{}
		sendinit.Op=constant.RabbitApi["Init"]
		sendinit.Args.Stage=1
		sendinitjson,err:=json.Marshal(sendinit)
		if err!=nil{
			logs.Error(err)
		}
		Send(setting.UpChannel,string(sendinitjson))
	}
	//bson.M{"x": bson.M{"$ne": 3}}
	//将所有未删除的策略数据发送过去
	//先查找StrategySwitch的state!=3的所有策略，然后将获取strageyinfo
	if init.Args.Stage==2{
		strategySwitch:=[]models.StrategySwitch{}
		strswitch :=models.ConnecToDB("StrategySwitch")

		strswitch.Find(bson.M{"strategyStatus": bson.M{"$ne": 3}}).All(strategySwitch)
		for _,mSwitch := range strategySwitch{
			//interflow.StrategyInfo()
			//mSwitch.NodeId
			//mSwitch.StrategyId
		}

	}
}

func strategyinit(data string){
	//发送StrategyInfo
	////////////解析出StrategyId
	//strategyinfoargs,err:=interflow.StrategyInfo(1,1)

	// if err!=nil{
	// 	logs.Error(err)
	// 	return
	// }

	//发送StrategyParam

	//解析出id
	initReq:=&struct{
		Op string `json:"op"`
		Args models.StrategyInitReq `json:"args"`
	}{}
	err:=json.Unmarshal([]byte(data),initReq)
	if err!=nil{
		logs.Error(err)
		return
	}
	//models.StrategyInitReq

	strategyparamargs,err:=interflow.StrategyParam(initReq.Args.NodeId,initReq.Args.StrategyId)
	if err!=nil{
		logs.Error(err)
		return
	}
	strategyparam:=&struct{
		Op string `json:"op"`
		Args *models.StrategyParam `json:"args"`
	}{
		"StrategyParam",
		&strategyparamargs,
	}
	strategyparamjson,err:=json.Marshal(strategyparam)
	Send(setting.UpChannel,string(strategyparamjson))
	if err!=nil{
		logs.Error(err)
	}

	//发送StrategyInitFinished
	//strategyinitfinishedNodeid:=interflow.StrategyInitFinished(1)
	//

	//initReq.Args.NodeId,initReq.Args.StrategyId
	strategyinitfinished:=&struct{
		Op string `json:"op"`
		Args models.StrategyInitFinished `json:"args"`
	}{
		"StrategyInitFinished",
		models.StrategyInitFinished{StrategyId:initReq.Args.StrategyId,NodeId:initReq.Args.NodeId,ErrorId:0,ErrorMsg:"0",},
	}
	strategyinitfinishedjson,err:=json.Marshal(strategyinitfinished)
	
	Send(setting.UpChannel,string(strategyinitfinishedjson))
	if err!=nil{
		logs.Error(err)
	}
}