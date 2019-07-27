package models

import (
	"time"
)


type Ws_send struct{
	Id int `json:"id"`
	Op string `json:"op"`
	//Args string  `json:"args"`
	Args interface{} `json:"args"`
}

type Ws_recv struct{
	Id int `json:"id"`
	Op string `json:"op"`
	Code int `json:"code"`
	Msg string `json:"msg"`
	//Data string `json:"data"`
	Data interface{} `json:"data"`
}
type User struct {
 	User string    `bson:"user"  json:"user"` 
 	AccountId uint64 `bson:"accountId"  json:"accountId"`
 	Username string `bson:"username"  json:"username"`
 	Password string `bson:"password"  json:"password"`
 	PowerLevel int `bson:"poweLevel"  json:"powerLevel"`
 	Parytof string `bson:"parytof"  json:"parytof"`
 	CreateTime time.Time `bson:"createTime"  json:"createtime"`
 	UpdateTime time.Time `bson:"updateTime"  json:"updatetime"`
}

// type Onorder struct {
// 	Investorid       string `bson:"investorid"  json:"investorid"`
// 	Userid           string `bson:"userid"  json:"userid"`
// 	Ordersysid       string `bson:"ordersysid"  json:"ordersysid"`
// 	Instrumentid     string `bson:"instrumentid"  json:"instrumentid"`
// 	Direction        string `bson:"direction"  json:"direction"`
// 	Offsetflag       string `bson:"offsetflag"  json:"offsetflag"`
// 	Limitprice       float64 `bson:"limitprice"  json:"limitprice"`
// 	Volume           uint `bson:"volume"  json:"volume"`
// 	Orderstatus      string `bson:"orderstatus"  json:"orderstatus"`
// 	Volumetraded     uint `bson:"volumetraded"  json:"volumetraded"`
// 	Volumeremain     uint `bson:"volumeremain"  json:"volumeremain"`
// 	Volumecancled    uint `bson:"volumecancled"  json:"volumecancled"`
// 	Inserttime       string `bson:"inserttime"  json:"inserttime"`
// 	Localdate        time.Time `bson:"localdate"  json:"localdate"`
// 	Exchangeid       string `bson:"exchangeid"  json:"exchangeid"`
// 	Userorderlocalid string `bson:"userorderlocalid"  json:"userorderlocalid"`
// }

type LogInfo struct{
	//AccountId uint64 `bson:"accountId"  json:"accountId"`
	StrategyId uint64  `bson:"strategyId"  json:"strategyId"`
	// LogTime int `bson:"logTime"  json:"logTime"`
	// LogType string `bson:"logType"  json:"logType"`
	// EpochTime uint64 `bson:"epochTime"  json:"epochTime"`
	LogText string `bson:"logText"  json:"logText"`
}

type NodeInitConfig struct{
	NodeId int `bson:"nodeId"  json:"nodeId"`
	FieldName string `bson:"fieldName"  json:"fieldName"`
	FieldValue string `bson:"fieldValue"  json:"fieldValue"`
	FieldType string `bson:"fieldTyped"  json:"fieldTyped"`
}


//交易前置信息
type TdInfo struct{
	NickName string `bson:"nickName"  json:"nickName"`
	InterfaceType string `bson:"interfaceType"  json:"interfaceType"`
	Id  int `bson:"id"  json:"id"`
	NodeId int `bson:"nodeId"  json:"nodeId"`
	InvestorId string `bson:"investorId"  json:"investorId"`
	Password string `bson:"password"  json:"password"`
	BrokerId string `bson:"brokerId"  json:"brokerId"`
	AppId string `bson:"appId"  json:"appId"`
	AuthenCode string `bson:"authenCode"  json:"authenCode"`
	FrontAddr string `bson:"frontAddr"  json:"frontAddr"`
}

//行情前置信息
type Mdlnfo struct{
	NickName string `bson:"nickName"  json:"nickName"`
	InterfaceType string `bson:"interfaceType"  json:"interfaceType"`
	Id  int `bson:"id"  json:"id"`
	NodeId int `bson:"nodeId"  json:"nodeId"`
	InvestorId string `bson:"investorId"  json:"investorId"`
	Password string `bson:"password"  json:"password"`
	BrokerId string `bson:"brokerId"  json:"brokerId"`
	AppId string `bson:"appId"  json:"appId"`
	AuthenCode string `bson:"authenCode"  json:"authenCode"`
	FrontAddr string `bson:"frontAddr"  json:"frontAddr"`
	Background string `bson:"background"  json:"background"`
	BackgroundFrontAddr string `bson:"backgroundFrontAddr"  json:"backgroundFrontAddr"`
	UdpAddr string `bson:"udpAddr"  json:"udpAddr"`
}
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
//交易合约配置,对外
type SymbolInfo struct{

	SymbolId string `bson:"symbolId"  json:"symbolId"`
	Symbol string `bson:"symbol"  json:"symbol"`
	ExchangeCode string `bson:"exchangeCode"  json:"exchangeCode"`
	NodeId int `bson:"nodeId"  json:"nodeid"`
	Multiplier int `bson:"multiplier"  json:"multiplier"`
	TickSize float64 `bson:"tickSize"  json:"tickSize"`
	////////////////////////////////////////////////
	SymbolType string `bson:"symbolType"  json:"symbolType"`
	////////////////////////////////////////////////
	RoundLot int `bson:"roundLot"  json:"roundLot"`
	Sides int `bson:"sides"  json:"sides"`

}
//策略交易时间段
type SymbolTradingPeriod struct{
	SymbolId string `bson:"symbolId"  json:"symbolId"`
	TradingPeriodType int `bson:"tradingPeriodType"  json:"tradingperiodType"`
	StartTime int `bson:"startTime"  json:"startTime"`
	EndTime int `bson:"endTime"  json:"endTime"`
}
//保证金
type SymbolMarginDefault struct{
	  SymbolId string `bson:"symbolId"  json:"symbolId"`
	  InitialMargin float64 `bson:"initialMargin"  json:"initialMargin"`
	  MaintenanceMargin float64 `bson:"maintenanceMargin"  json:"maintenanceMargin"`
}
//手续费
type SymbolFeeDefault struct{
      SymbolId string `bson:"symbolId"  json:"symbolId"`
      Fee float64 `bson:"fee"  json:"fee"`
      FeeType string `bson:"feeType"  json:"feeType"`
      FeeFormat string `bson:"feeFormat"  json:"feeFormat"`
}
//账户配置,对外
type AccountInfo struct{
	AccountId uint64 `bson:"accountId"  json:"accountId"`
	Currency string `bson:"currency"  json:"currency"`
	CashBalance float64 `bson:"cashBalance"  json:"cashBalance"`
}
//每天的持仓信息
type AccountPnlDaily struct{
	AccountId uint64 `bson:"accountId"  json:"accountId"`
	SymbolId string `bson:"symbolId"  json:"symbolId"`
    TradingDay string `bson:"tradingDay"  json:"tradingDay"`
    AvgBuyPrice float64 `bson:"avgBuyPrice"  json:"avgBuyPrice"`
    AvgSellPrice float64 `bson:"avgSellPrice"  json:"avgSellPrice"`
    BuyQuantity float64 `bson:"buyQuantity"  json:"buyQuantity"`
    SellQuantity float64 `bson:"sellQuantity"  json:"sellQuantity"`
    TodayLong float64 `bson:"todayLong"  json:"todayLong"`
    TodayShort float64 `bson:"todayShort"  json:"todayShort"`
    YesterdayLong float64 `bson:"yesterdayLong"  json:"yesterdayLong"`
    YesterdayShort float64 `bson:"yesterdayShort"  json:"yesterdayShort"`
    NetPosition float64 `bson:"netPosition"  json:"netPosition"`
	Turnover float64 `bson:"turnover"  json:"turnover"`
	AggregatedFee float64 `bson:"aggregatedFee"  json:"aggregatedFee"`
		
		
}
//策略info

type StrategyInfo struct{
	NodeId int `bson:"nodeId"  json:"nodeId"`
	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
	StrategyName string `bson:"strategyName"  json:"strategyName"`
	AccountId uint64 `bson:"accountId"  json:"accountId"`
	StrategyOrderType string `bson:"strategyOrderType"  json:"strategyOrderType"`
	StrategyArchiveType  string `bson:"strategyArchiveType"  json:"strategyArchiveTypee"`
	StrategyArchiveName string `bson:"strategyArchiveName"  json:"strategyArchiveName"`
	StrategyArchiveOwner string `bson:"strategyArchiveOwner"  json:"strategyArchiveOwner"`
	//Tradeable int `bson:"tradeable"  json:"tradeable"`
	//StrategyStatus int `bson:"strategyStatus"  json:"strategystatus"`
	CreateTime time.Time `bson:"createTime"  json:"-"`
	//UpdateTime time.Time `bson:"updateTime"  json:"updatetime"`
	//UpdateAccountId uint64 `bson:"updateAccountId"  json:"updateaccountid"`
}

//Tradeable  StrategyStatus UpdateTime UpdateAccountId
//策略切换，NodeId?
type StrategySwitch struct{
	NodeId int `bson:"nodeId"  json:"-"`
	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
	//Tradeable int `bson:"tradeable"  json:"tradeable"`
	StrategyStatus int `bson:"strategyStatus"  json:"strategyStatus"`
	UpdateTime time.Time `bson:"updateTime"  json:"-"`
}

//策略参数
type StrategyParam struct{
	NodeId int `bson:"nodeId"  json:"-"`
	StrategyId  uint64 `bson:"strategyId"  json:"strategyId"`
	Params string `bson:"params"  json:"params"`
	ValueName string `bson:"valueName"  json:"valueName"`
	ValueInt int `bson:"valueInt"  json:"valueInt"`
	ValueDouble float64 `bson:"valueDouble"  json:"valueDouble"`
	ValueDate int `bson:"valueDate"  json:"valueDate"`
	ValueString string `bson:"valueString"  json:"valueString"`
	ValueType string `bson:"valueType"  json:"valueType"`
	CreateTime time.Time `bson:"createTime"  json:"-"`
}

type UpdateStrategyParams struct{
	NodeId int `bson:"nodeId"  json:"-"`
	StrategyId  uint64 `bson:"strategyId"  json:"strategyId"`
	Params string `bson:"params"  json:"params"`
	ValueName string `bson:"valueName"  json:"valueName"`
	ValueInt int `bson:"valueInt"  json:"valueInt"`
	ValueDouble float64 `bson:"valueDouble"  json:"valueDouble"`
	ValueDate int `bson:"valueDate"  json:"valueDate"`
	ValueString string `bson:"valueString"  json:"valueString"`
	ValueType string `bson:"valueType"  json:"valueType"`
}
//
//策略涉及的合约
type StrategySymbol struct{
	NodeId int `bson:"nodeId"  json:"-"`
	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
	SymbolId string `bson:"symbolId"  json:"symbolId"`
	TdNickName string `bson:"tdNickName"  json:"tdNickName"`
	CreateTime time.Time `bson:"createTime"  json:"-"`
}

type StrategySymbolMarginOverride struct{
	NodeId int `bson:"nodeId"  json:"-"`
	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
	SymbolId string `bson:"symbolId"  json:"symbolId"`
	InitialMargin float64 `bson:"initialMargin"  json:"initialMargin"`
	MaintenanceMargin float64 `bson:"maintenanceMargin"  json:"maintenanceMargin"`
	CreateTime time.Time `bson:"createTime"  json:"-"`
}
//StrategyId  SymbolId Fee FeeType FeeFormat
type StrategySymbolFeeOverride struct{
	NodeId int `bson:"nodeId"  json:"-"`
	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
	SymbolId string `bson:"symbolId"  json:"symbolId"`
	Fee float64 `bson:"fee"  json:"fee"`
	FeeType string `bson:"feeType"  json:"feeType"`
	FeeFormat string `bson:"feeFormat"  json:"feeFormat"`
	CreateTime time.Time `bson:"createTime"  json:"-"`
}
//策略每天的信息表
type StrategyPnlDaily struct{
	NodeId int `bson:"nodeId"  json:"-"`
	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
	SymbolId string `bson:"symbolId"  json:"symbolId"`
	TradingDay string `bson:"tradingDay"  json:"tradingDay"`
	AvgBuyPrice float64  `bson:"avgBuyPrice"  json:"avgBuyPrice"`
	AvgSellPrice float64  `bson:"avgSellPrice"  json:"avgSellPrice"`
	BuyQuantity float64  `bson:"buyQuantity"  json:"buyQuantity"`
	SellQuantity float64  `bson:"sellQuantity"  json:"sellQuantity"`
	TodayLong float64  `bson:"todayLong"  json:"todayLong"`
	TodayShort float64  `bson:"todayShort"  json:"todayShort"`
	YesterdayLong float64  `bson:"yesterdayLong"  json:"yesterdayLong"`
	YesterdayShort float64  `bson:"yesterdayShort"  json:"yesterdayShort"`
	NetPosition float64  `bson:"netPosition"  json:"netPosition"`
	Turnover float64  `bson:"turnover"  json:"turnover"`
	AggregatedFee float64  `bson:"aggregatedFee"  json:"aggregatedFee"`
	  
}
//策略上传相关
type StrategyArchive struct{
	ArchiveName string `bson:"archiveName"  json:"archiveName"`
    ArchiveOwner string `bson:"archiveOwner"  json:"archiveOwner"`
    ArchiveType string `bson:"archiveType"  json:"archiveType"`
}


//策略初始化流程
//先发送SymbolInfo，收到StrategyInitReq后发送初始化数据StrategyInfo、StrategyParam等数据，结束发送StrategyInitFinished
//StrategyInitReq
type StrategyInitReq struct{
	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
	NodeId int `bson:"nodeId"  json:"nodeId"`
	Pid uint64 `bson:"pid"  json:"pid"`
	ErrorId int `bson:"errorId"  json:"errorId"`
	ErrorMsg string `bson:"errorMsg"  json:"errorMsg"`
}

type StrategyHeartbeat struct{
	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
	NodeId int `bson:"nodeId"  json:"nodeId"`
}
type StrategyInitFinished struct{
	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
	NodeId int `bson:"nodeId"  json:"nodeId"`
	ErrorId int `bson:"errorId"  json:"errorId"`
	ErrorMsg string `bson:"errorMsg"  json:"errorMsg"`
}

type StrategyOrder struct{
	TdNickName string `bson:"tdNickName"  json:"tdNickName"`
	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
	SymbolId string `bson:"symbolId"  json:"symbolId"`
	InstrumentId string `bson:"instrumentId"  json:"instrumentId"`
	ExchangeCode string `bson:"exchangeCode"  json:"exchangeCode"`
	Direction string `bson:"direction"  json:"direction"`
	Offset string `bson:"offset"  json:"offset"`
	LimitPrice float64 `bson:"limitPrice"  json:"limitPrice"`
	HedgeFlag int `bson:"hedgeFlag"  json:"hedgeFlag"`
	OrderType int `bson:"orderType"  json:"orderType"`
	VolumeTotalOriginal float64 `bson:"volumeTotalOriginal"  json:"volumeTotalOriginal"`
	OrderSysId string `bson:"orderSysId"  json:"orderSysId"`
	VolumeFilled float64 `bson:"volumeFilled"  json:"volumeFilled"`
	OrderRef int `bson:"orderRef"  json:"orderRef"`
	OrderMsgType int `bson:"orderMsgType"  json:"orderMsgType"`
	Volume float64 `bson:"volume"  json:"volume"`
	VolumeRemained float64 `bson:"volumeRemained"  json:"volumeRemained"`
	TradingDay string `bson:"tradingDay"  json:"tradingDay"`
	OrderStatus string `bson:"orderStatus"  json:"orderStatus"`
	Price float64 `bson:"price"  json:"price"`
	TradeId string `bson:"tradeId"  json:"tradeId"`
	OrdRejReason int `bson:"ordRejReason"  json:"ordRejReason"`
	InsertTime string `bson:"insertTime"  json:"insertTime"`
	UpdateTime string `bson:"updateTime"  json:"updateTime"`
	CancelTime string `bson:"cancelTime"  json:"cancelTime"`
	FrontId int `bson:"frontId"  json:"frontId"`
	SessionId int `bson:"sessionId"  json:"sessionId"`
	StatusMsg string `bson:"statusMsg"  json:"statusMsg"`
	Fee float64 `bson:"fee"  json:"fee"`
	CounterType string `bson:"counterType"  json:"counterType"`
	CounterSysId uint64 `bson:"counterSysId"  json:"counterSysId"`
	CancelAttempts int `bson:"cancelAttempts"  json:"cancelAttempts"`
	TimeStamp uint64 `bson:"timeStamp"  json:"timeStamp"`
	EpochTimeReturn uint64 `bson:"epochTimeReturn"  json:"epochTimeReturn"`
	FuncName string `bson:"funcName"  json:"funcName"`
	EpochTimeReqBefore uint64 `bson:"epochTimeReqBefore"  json:"epochTimeReqBefore"`
	EpochTimeReqAfter uint64 `bson:"epochTimeReqAfter"  json:"epochTimeReqAfter"`
	BorrowFlag int `bson:"borrowFlag"  json:"borrowFlag"`
	UserId string `bson:"userId"  json:"userId"`
	BrokerId string `bson:"brokerId"  json:"brokerId"`
	InvestorId string `bson:"investorId"  json:"investorId"`
}

// type InnerPluginRisOrderInsert struct{
// 	PluginOrderId int `bson:"pluginOrderId"  json:"pluginOrderId"`
// 	InstrumentId string `bson:"instrumentId"  json:"instrumentId"`
// 	ExchangeCode string `bson:"exchangeCode"  json:"exchangeCode"`
// 	Direction string `bson:"direction"  json:"direction"`
// 	LimitPrice float64 `bson:"limitPrice"  json:"limitPrice"`
// 	VolumeTotalOriginal float64 `bson:"volumeTotalOriginal"  json:"volumeTotalOriginal"`
// 	OrderType int `bson:"orderType"  json:"orderType"`
// 	MinVolume int `bson:"minVolume"  json:"minVolume"`
// }

// type InnerPluginRawOrderInsert struct{
// 	PluginOrderId int `bson:"pluginOrderId"  json:"pluginOrderId"`
// 	InstrumentId string `bson:"instrumentId"  json:"instrumentId"`
// 	ExchangeCode string `bson:"exchangeCode"  json:"exchangeCode"`
// 	Direction string `bson:"direction"  json:"direction"`
// 	LimitPrice float64 `bson:"limitPrice"  json:"limitPrice"`
// 	VolumeTotalOriginal float64 `bson:"volumeTotalOriginal"  json:"volumeTotalOriginal"`
// 	OrderType int `bson:"orderType"  json:"orderType"`
// 	MinVolume int `bson:"minVolume"  json:"minVolume"`
// 	Offset string `bson:"offset"  json:"offset"`
//   }
// type InnerPluginOrderAction struct{
// 	PluginOrderId int `bson:"pluginOrderId"  json:"pluginOrderId"`
// 	InstrumentId string `bson:"instrumentId"  json:"instrumentId"`
// 	ExchangeCode string `bson:"exchangeCode"  json:"exchangeCode"`
// 	OrderRef int `bson:"orderRef"  json:"orderRef"`
// 	OrderSysId string `bson:"orderSysId"  json:"orderSysId"`
// }  
// type InnerPluginOrder struct{
// 	PluginOrderId int `bson:"pluginOrderId"  json:"pluginOrderId"`
// 	OrderRef int `bson:"orderRef"  json:"orderRef"`
// 	InstrumentId string `bson:"instrumentId"  json:"instrumentId"`
// 	ExchangeCode string `bson:"exchangeCode"  json:"exchangeCode"`
// 	LimitPrice float64 `bson:"limitPrice"  json:"limitPrice"`
// 	Direction string `bson:"direction"  json:"direction"`
// 	Offset string `bson:"offset"  json:"offset"`
// 	OrderType int `bson:"orderType"  json:"orderType"`
// 	OrderSysId string `bson:"orderSysId"  json:"orderSysId"`
// 	VolumeTotalOriginal float64 `bson:"volumeTotalOriginal"  json:"volumeTotalOriginal"`
// 	VolumeFilled float64 `bson:"volumeFilled"  json:"volumeFilled"`
// 	VolumeRemained float64 `bson:"volumeRemained"  json:"volumeRemained"`
// 	OrderMsgType int `bson:"orderMsgType"  json:"orderMsgType"`
// 	OrderStatus string `bson:"orderStatus"  json:"orderStatus"`
// 	Price float64 `bson:"price"  json:"price"`
// 	Volume float64 `bson:"volume"  json:"volume"`
// 	OrdRejReason int `bson:"ordRejReason"  json:"ordRejReason"`
// 	StatusMsg string `bson:"statusMsg"  json:"statusMsg"`
// }
// type InnerPluginTimer struct{
// 	StrategyId uint64 `bson:"strategyId"  json:"strategyId"`
// 	TimerId int `bson:"timerId"  json:"timerId"`
// }