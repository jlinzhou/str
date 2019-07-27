package constant


type ApiName  struct{
	Id int
	Name string
}
var Allapi=map[string]*ApiName{
	"auth":{0,"auth"},
	"createStrategy":{1,"createStrategy"},
	"startStrategy":{2,"startStrategy"},
	"pauseStrategy":{3,"pauseStrategy"},
	"deleteStrategy":{4,"deleteStrategy"},
	"modifyStrategy":{5,"modifyStrategy"},
	"modifyMarginOverride":{6,"modifyMarginOverride"},
	"modifyFeeOverride":{7,"modifyFeeOverride"},
}


var RabbitApi=map[string]string{
	"SymbolInfo":"SymbolInfo",
	"StrategyInitReq":"StrategyInitReq",
	"StrategyInfo":"StrategyInfo",
	"StrategyParam":"StrategyParam",
	"StrategyInitFinished":"StrategyInitFinished",

	"StrategySwitch":"StrategySwitch",
	"UpdateStrategyParams":"UpdateStrategyParams",

	"Init":"Init",
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

}