package interflow
import(
	"strategy_middle/models"
	"github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2/bson"
)


func StrategyInitFinished(strategyid uint64)(nodeid int,err error){
	st:=models.ConnecToDB("StrategyInfo")
	strategyinfo :=&models.StrategyInfo{}
	err = st.Find(bson.M{"strategyId":  strategyid}).One(strategyinfo)
	if err!=nil{
		logs.Error(err)
	}
	nodeid =strategyinfo.NodeId
	return
}


//SymbolInfo、StrategyInfo、StrategyParam、StrategyInitFinished
//根据(nodeid)strategyid查找StrategySymbol表找到SymbolInfo,再查找SymbolInfo表信息

func SymbolInfo(nodeid int,strategyid uint64)(allSymbolinfo []models.SymbolInfo,err error){

	//存储所有合约信息
	//var allSymbolinfo =make([]models.SymbolInfo{})
	st:=models.ConnecToDB("StrategySymbol")
	strategysymbol:=[]models.StrategySymbol{}
	err = st.Find(bson.M{"strategyId":strategyid,"nodeId":nodeid}).All(&strategysymbol)
	if err!=nil{
		logs.Error(err)
	}
	for _,ssymbol := range strategysymbol{
		sy:=models.ConnecToDB("SymbolInfo")
		symbolinfo:=&models.SymbolInfo{}
		err = sy.Find(bson.M{"symbolId":  ssymbol.SymbolId}).One(symbolinfo)
		if err!=nil{
			logs.Error(err)
		}
		allSymbolinfo =append(allSymbolinfo,*symbolinfo)
	}
	return 
}

func SymbolTradingPeriod(symbolid string)(this models.SymbolTradingPeriod,err error){
	st:=models.ConnecToDB("SymbolTradingPeriod")
	err = st.Find(bson.M{"symbolId":  symbolid}).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}

func SymbolMarginDefault(symbolid string)(this models.SymbolMarginDefault,err error){
	st:=models.ConnecToDB("SymbolMarginDefault")
	err = st.Find(bson.M{"symbolId":  symbolid}).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}

func SymbolFeeDefault(symbolid string)(this models.SymbolFeeDefault,err error){
	st:=models.ConnecToDB("SymbolFeeDefault")
	err = st.Find(bson.M{"symbolId":  symbolid}).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}

func AccountInfo(symbolid string)(this models.AccountInfo,err error){
	st:=models.ConnecToDB("AccountInfo")
	err = st.Find(bson.M{"symbolId":  symbolid}).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}


func StrategyInfo(nodeid int,strategyid uint64)(this models.StrategyInfo,err error){

	st:=models.ConnecToDB("StrategyInfo")
	err = st.Find(bson.M{"strategyId":  strategyid,"nodeId":nodeid}).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}

func StrategySymbol(nodeid int,strategyid uint64)(this models.StrategySymbol,err error){

	st:=models.ConnecToDB("StrategySymbol")
	err = st.Find(bson.M{"strategyId":  strategyid,"nodeId":nodeid}).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}

func StrategySymbolMarginOverride(nodeid int,strategyid uint64)(this models.StrategySymbolMarginOverride,err error){

	st:=models.ConnecToDB("StrategySymbolMarginOverride")
	err = st.Find(bson.M{"strategyId":  strategyid,"nodeId":nodeid}).Sort("-createTime").Limit(1).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}

func StrategySymbolFeeOverride(nodeid int,strategyid uint64)(this models.StrategySymbolFeeOverride,err error){

	st:=models.ConnecToDB("StrategySymbolFeeOverride")
	err = st.Find(bson.M{"strategyId":  strategyid,"nodeId":nodeid}).Sort("-createTime").Limit(1).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}

func StrategyParam(nodeid int,strategyid uint64)(this models.StrategyParam,err error){
	st:=models.ConnecToDB("StrategyParam")
	err = st.Find(bson.M{"strategyId":  strategyid,"nodeId":nodeid}).Sort("-createTime").Limit(1).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}


func StrategySwitch(nodeid int,strategyid uint64)(this models.StrategySwitch,err error){
	st:=models.ConnecToDB("StrategySwitch")
	err = st.Find(bson.M{"strategyId":  strategyid,"nodeId":nodeid}).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}



func StrategyArchive(archiveowner string)(this models.StrategyArchive,err error){
	st:=models.ConnecToDB("StrategyParam")
	err = st.Find(bson.M{"archiveOwner":  archiveowner}).One(&this)
	if err!=nil{
		logs.Error(err)
	}
	return 
}

