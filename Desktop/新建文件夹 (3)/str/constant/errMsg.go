package constant

import "errors"
var Success int



// var AlreadyConnected = errors.New("already connected")
// var NotConnected = errors.New("not connected")
// var ConnectionFailed = errors.New("connection failed")
// var OrderNotFound = errors.New("order not found")
// var InvalidConnector = errors.New("invalid connector")
// var UnknownInstrument = errors.New("unknown instrument")
// var UnsupportedOrderType = errors.New("unsupported order type")
// var DownloadFailed = errors.New("download failed")
//json解析失败,1001
//插入表失败
//查找失败

var Errmap=map[int]error{
	0:errors.New("Success"),
	1001:errors.New("Invalid json params"),
	1002:errors.New("Marshal json fail"),
	1003:errors.New("unknown StrategyArchiveOwner"),
	1004:errors.New("unknown NodeId"),
	
	1010:errors.New("update err"),

	1020:errors.New("not find"),

	1500:errors.New("insert table fail"),

	2001:errors.New("rabbitmq send err"),

	3001:errors.New("unknown Op"),
	3002:errors.New("unknown Pattern"),
}



// const(
// 	ErrInvalidParams = "Invalid params"
// )