package api

import(
	//"net/http"
	//"runtime"
	"github.com/astaxie/beego/logs"
	"strategy_middle/models"
	"strategy_middle/controller/interflow"
	"strategy_middle/rabbitmq"
	"gopkg.in/mgo.v2/bson"
	"strategy_middle/constant"
	"strategy_middle/setting"
	//"time"
	//"Goroutine"
	//"os"
	//"bytes"
	"encoding/json"
	//"strconv"
	//"reflect"
	//"github.com/gin-gonic/gin"
	//"github.com/gorilla/websocket"
	
	"time"
)

// 定义interface，interface是一组method签名的组合
// interface可以被任意对象实现，一个对象也可以实现多个interface
// 任意类型都实现了空interface（也就是包含0个method的interface）
// 空interface可以存储任意类型的值
// interface定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。
// type StrategyFunc interface {





//StrategyStatus
//0 正在执行 1暂停 2已经执行完成 3删除
////创建策略步骤，先修改StrategyInfo表，加创建信息之类的，然后是StrategyParam表，加参数和id
//创建合约:先找到User表中的AccountId,再统计StrategyInfo表中策略最大值，计算出strategyid，
//插入StrategyInfo、StrategyParam、StrategySymbol、StrategySwitch，发送到rabbitmq
func (this *ApiOp) CreateStrategy()(err error){

	//msg :=&strategy_op{}
	//
	msg :=&struct{
		Id int `json:"id"`
		Op string `json:"op"`
		Args *struct{
			NodeId int `json:"nodeId"`
			SymbolInfo string `json:"symbolInfo"`
			StrategyName string `json:"strategyName"`
			StrategyArchiveOwner string `json:"strategyArchiveOwner"`
			StrategyArchiveName string `json:"strategyArchiveName"`
			Params string `json:"params"`
			// ValueName string `json:"valueName"`
			// ValueInt string `json:"valueInt"`
			// ValueDouble string `json:"valueDouble"`
			// ValueDate string `json:"valueDate"`
			// ValueString string `json:"valueString"`
			// ValueType string `json:"valueType"`
		} `json:"args"`
	}{}
	//
	err =json.Unmarshal(this.Message, msg)

	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1001,"fail",constant.Errmap[1001],err)
		return
	}

	//查找账户id
	us := models.ConnecToDB("User")
	user :=&models.User{}
	err = us.Find(bson.M{"user":  msg.Args.StrategyArchiveOwner}).One(user)
	if err != nil {
		logs.Error(err)
    	this.Send_callback_fail(1003,"fail",constant.Errmap[1003],err)
    	return
	}

	//查找同一个NodId下的策略最大数
	st:=models.ConnecToDB("StrategyInfo")
	//这应该查找最大值，不是计数
	max,err:=st.Find(bson.M{"nodeId":msg.Args.NodeId}).Count()
	strategyid:=uint64(max+1)

	if err != nil {
		logs.Error(err)
    	this.Send_callback_fail(1004,"fail",constant.Errmap[1004],err)
    	return
	}


	//插入StrategyInfo数据
	strategyinfo:=models.StrategyInfo{
		NodeId:msg.Args.NodeId,
		StrategyName:msg.Args.StrategyName,
		StrategyId:strategyid,
		AccountId:user.AccountId,
		StrategyArchiveName: msg.Args.StrategyArchiveName,
		StrategyArchiveOwner:msg.Args.StrategyArchiveOwner,
		CreateTime:time.Now(),
	}

	err =st.Insert(&strategyinfo)

	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1500,"fail",constant.Errmap[1500],err)
    	return
	}

	//插入StrategyParam
	stp := models.ConnecToDB("StrategyParam")

	strategyparams:=models.StrategyParam{
		NodeId:msg.Args.NodeId,
		StrategyId:strategyid,
		Params:msg.Args.Params,
		// ValueName:msg.Args.ValueName,
		// ValueInt:msg.Args.ValueInt,
		// ValueDouble:msg.Args.ValueDouble,
		// ValueDate:msg.Args.ValueDate,
		// ValueString:msg.Args.ValueStrings,
		// ValueType:msg.Args.ValueType,
		CreateTime:time.Now(),
	}
	err =stp.Insert(&strategyparams)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1500,"fail",constant.Errmap[1500],err)
    	return
	}

	//插入StrategySwitch
	sts := models.ConnecToDB("StrategySwitch")
	strategyswitch:=models.StrategySwitch{
		NodeId:msg.Args.NodeId,
		StrategyId:strategyid,
		StrategyStatus:1,
		UpdateTime:time.Now(),
	}

	err =sts.Insert(&strategyswitch)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1500,"fail",constant.Errmap[1500],err)
    	return
	}
	//插入StrategySymbol数据
	stsym := models.ConnecToDB("StrategySymbol")
	strategysymbol:=models.StrategySymbol{
		NodeId:msg.Args.NodeId,
		StrategyId:strategyid,
		SymbolId:msg.Args.SymbolInfo,
		CreateTime:time.Now(),
	}
	err =stsym.Insert(&strategysymbol)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1500,"fail",constant.Errmap[1500],err)
    	return
	}

	//发送rabbitmq
	strategyargs,err:= interflow.StrategyInfo(msg.Args.NodeId,strategyid)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1020,"fail",constant.Errmap[1020],err)
    	return
	}
	strategy_info :=&struct{
		Op string `json:"op"`
		Args models.StrategyInfo `json:"args"`
	}{
		constant.RabbitApi["StrategyInfo"],
		strategyargs,
	}
	strategyinfojson,err:=json.Marshal(strategy_info)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1002,"fail",constant.Errmap[1002],err)
    	return
	}
	//根据nodeid选择channel
	err=rabbitmq.Send(setting.UpChannel, string(strategyinfojson))
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(2000,"fail",constant.Errmap[2000],err)
	}else{
		this.Send_callback_success()
	}
	return
}

// //开始策略
func (this *ApiOp) StartStrategy()(err error){
	msg :=&struct{
		Id int `json:"id"`
		Op string `json:"op"`
		Args *struct{
			NodeId int `json:"nodeId"`
			StrategyId uint64 `json:"strategyId"`
		} `json:"args"`
	}{}
	err=json.Unmarshal(this.Message, msg)
	if err!=nil{
		this.Send_callback_fail(1001,"fail",constant.Errmap[1001],err)
		return
	}

	st:=models.ConnecToDB("StrategySwitch")
	err =st.Update(bson.M{"strategyId": msg.Args.StrategyId,"nodeId":msg.Args.NodeId}, bson.M{"$set": bson.M{"strategyStatus": 0,"updateTime":time.Now()}})
	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1010,"fail",constant.Errmap[1010],err)
		return
	}
	//发送rabbitmq
	thisargs,err:= interflow.StrategySwitch(msg.Args.NodeId,msg.Args.StrategyId)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1020,"fail",constant.Errmap[1020],err)
    	return
	}
	info:=&struct{
		Op string `json:"op"`
		Args models.StrategySwitch `json:"args"`
	}{
		constant.RabbitApi["StrategySwitch"],
		thisargs,
	}
	infojson,err:=json.Marshal(info)
	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1002,"fail",constant.Errmap[1002],err)
		return
	}

	err=rabbitmq.Send(setting.UpChannel, string(infojson))
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(2000,"fail",constant.Errmap[2000],err)
	}else{
		this.Send_callback_success()
	}
	return
}

//根据StrategyId来暂停
func (this *ApiOp) PauseStrategy()(err error){
	msg :=&struct{
		Id int `json:"id"`
		Op string `json:"op"`
		Args *struct{
			NodeId int `json:"nodeId"`
			StrategyId uint64 `json:"strategyId"`
		} `json:"args"`
	}{}
	err=json.Unmarshal(this.Message, msg)
	if err!=nil{
		this.Send_callback_fail(1001,"fail",constant.Errmap[1001],err)
		return
	}

	st:=models.ConnecToDB("StrategySwitch")
	err =st.Update(bson.M{"strategyId": msg.Args.StrategyId,"nodeId":msg.Args.NodeId}, bson.M{"$set": bson.M{"strategyStatus": 1,"updateTime":time.Now()}})
	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1010,"fail",constant.Errmap[1010],err)
		return
	}
	//发送rabbitmq
	thisargs,err:= interflow.StrategySwitch(msg.Args.NodeId,msg.Args.StrategyId)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1020,"fail",constant.Errmap[1020],err)
    	return
	}
	info:=&struct{
		Op string `json:"op"`
		Args models.StrategySwitch `json:"args"`
	}{
		//"StrategySwitch",
		constant.RabbitApi["StrategySwitch"],
		thisargs,
	}
	infojson,err:=json.Marshal(info)

	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1002,"fail",constant.Errmap[1002],err)
		return
	}

	err=rabbitmq.Send(setting.UpChannel, string(infojson))
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(2000,"fail",constant.Errmap[2000],err)
	}else{
		this.Send_callback_success()
	}
	return

}
//根据StrategyId来删除
func (this *ApiOp) DeleteStrategy()(err error){
	msg :=&struct{
		Id int `json:"id"`
		Op string `json:"op"`
		Args *struct{
			NodeId int `json:"nodeId"`
			StrategyId  uint64 `json:"strategyId"`
		} `json:"args"`
	}{}
	err=json.Unmarshal(this.Message, msg)
	if err!=nil{
		this.Send_callback_fail(1001,"fail",constant.Errmap[1001],err)
		return
	}		
	st:=models.ConnecToDB("StrategySwitch")
	err =st.Update(bson.M{"strategyId": msg.Args.StrategyId,"nodeId":msg.Args.NodeId}, bson.M{"$set": bson.M{"strategyStatus": 3,"updateTime":time.Now()}})
	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1010,"fail",constant.Errmap[1010],err)
		return
	}

	//发送rabbitmq
	thisargs,err:= interflow.StrategySwitch(msg.Args.NodeId,msg.Args.StrategyId)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1020,"fail",constant.Errmap[1020],err)
    	return
	}
	info:=&struct{
		Op string `json:"op"`
		Args models.StrategySwitch `json:"args"`
	}{
		//"StrategySwitch",
		constant.RabbitApi["StrategySwitch"],
		thisargs,
	}
	infojson,err:=json.Marshal(info)
	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1002,"fail",constant.Errmap[1002],err)
		return
	}
	err=rabbitmq.Send(setting.UpChannel, string(infojson))
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(2000,"fail",constant.Errmap[2000],err)
	}else{
		this.Send_callback_success()
	}
	return

}

//根据id来修改，把旧的数据的tradeable改为3，新的策略参数新建一个StrategyId
//修改要把原来不变的数据放进来，旧的数据改变
//先修改StrategyInfo表，再修改StrategyParam表
//先把StrategySwitch表中的strategyStatus改为3，在重新生成新的策略

//修改了合约
func (this *ApiOp) ModifyStrategy()(err error){

	//args:{"strategyid":1,"strategyname":"aaa","params":"ssss"}
	msg :=&struct{
		Id int `json:"id"`
		Op string `json:"op"`
		Args *struct{
			NodeId int `json:"nodeId"`
			StrategyId  uint64 `json:"strategyId"`
			ModifySymbol int `json:"modifySymbol"`
			Params  string `json:"params"`
			// ValueName string `json:"valueName"`
			// ValueInt string `json:"valueInt"`
			// ValueDouble string `json:"valueDouble"`
			// ValueDate string `json:"valueDate"`
			// ValueString string `json:"valueString"`
			// ValueType string `json:"valueType"`
		} `json:"args"`
	}{}
	err=json.Unmarshal(this.Message, msg)
	if err!=nil{
		logs.Error(err)
		this.Send_callback_fail(1001,"fail",constant.Errmap[1001],err)
		return
	}	
	// //更改StrategySwitch
	// sts:=models.ConnecToDB("StrategySwitch")
	// err :=sts.Update(bson.M{"strategyId": msg.Args.StrategyId,"nodeId":msg.Args.NodeId}, bson.M{"$set": bson.M{"strategyStatus": 3,"updateTime":time.Now()}})
	// if err!=nil{
	// 	logs.Error(err)
	// 	send_callback_fail(this.Ws,this.Mt,5,"modifyStrategy",500,"fail",err)
	// 	return
	// }

	// //查找StrategyInfo，找到相关参数
	// stra :=&models.StrategyInfo{}
	// st:=models.ConnecToDB("StrategyInfo")
	// err = st.Find(bson.M{"strategyId": msg.Args.StrategyId,"nodeId":msg.Args.NodeId}).One(stra)
	// if err != nil {
	// 	logs.Error(err)
	// 	send_callback_fail(this.Ws,this.Mt,5,"modifyStrategy",500,"fail",err)
	// 	return
	// }

	// us := models.ConnecToDB("User")
	// user :=&models.User{}
	// err = us.Find(bson.M{"user":msg.Args.Username}).One(user)
	// if err != nil {
	// 	logs.Error(err)
	// 	send_callback_fail(ws,mt,5,"modifyStrategy",500,"fail",err)
	// 	return
	// }



	// //找到最大值
	// max,err:=st.Find(bson.M{"nodeId":msg.Args.NodeId}).Count()
	// strategyid :=uint64(max+1)

	// //插入StrategyInfo
	// strategyinfo:=models.StrategyInfo{
	// 	NodeId:msg.Args.NodeId,
	// 	StrategyId:strategyid,
	// 	StrategyName:stra.StrategyName,
	// 	AccountId:stra.AccountId,
	// 	StrategyArchiveName: stra.StrategyArchiveName,
	// 	StrategyArchiveOwner:stra.StrategyArchiveOwner,
	// 	//Tradeable:0,
	// 	//StrategyStatus:0,
	// 	CreateTime:time.Now(),
	// 	//UpdateTime:time.Now(),
	// 	//UpdateAccountId:user.AccountId,
	// }

	// err =st.Insert(&strategyinfo)
	// if err!=nil{
	// 	logs.Error(err)
	// 	send_callback_fail(this.Ws,this.Mt,5,"modifyStrategy",500,"fail",err)
	// 	return
	// }
	// //插入StrategyParam
	// stp:=models.ConnecToDB("StrategyParam")
	// strategyparam:=models.StrategyParam{
	// 	NodeId:msg.Args.NodeId,
	// 	StrategyId:strategyid,
	// 	Params:msg.Args.Params,
	// }
	// err =stp.Insert(&strategyparam)
	// //插入StrategySwitch
	// strategyswitch:=models.StrategySwitch{
	// 	NodeId:msg.Args.NodeId,
	// 	StrategyId:strategyid,
	// 	StrategyStatus:1,
	// 	UpdateTime:time.Now(),
	// }
	// err =sts.Insert(&strategyswitch)

	// if err!=nil{
	// 	logs.Error(err)
	// 	send_callback_fail(this.Ws,this.Mt,5,"modifyStrategy",500,"fail",err)

	// }else{
	// 	send_callback_success(this.Ws,this.Mt,5,"modifyStrategy",200,"success")
	// }

	stp:=models.ConnecToDB("StrategyParam")
	strategyparam:=models.StrategyParam{
		NodeId:msg.Args.NodeId,
		StrategyId:msg.Args.StrategyId,
		Params:msg.Args.Params,
		// ValueName:msg.Args.ValueName,
		// ValueInt:msg.Args.ValueInt,
		// ValueDouble:msg.Args.ValueDouble,
		// ValueDate:msg.Args.ValueDate,
		// ValueString:msg.Args.ValueStrings,
		// ValueType:msg.Args.ValueType,
		CreateTime:time.Now(),
	}
	err =stp.Insert(&strategyparam)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1500,"fail",constant.Errmap[1500],err)
    	return
	}

	//修改参数问题
	//发送rabbitmq
	thisargs,err:= interflow.StrategyParam(msg.Args.NodeId,msg.Args.StrategyId)
	if err!=nil{
		logs.Error(err)
    	this.Send_callback_fail(1020,"fail",constant.Errmap[1020],err)
    	return
	}
	info:=&struct{
		Op string `json:"op"`
		Args models.StrategyParam `json:"args"`
	}{
		//"UpdateStrategyParams",
		constant.RabbitApi["UpdateStrategyParams"],
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
	// 
	//UpdateStrategyParams
}