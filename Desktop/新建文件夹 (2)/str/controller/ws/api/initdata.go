package api
import(
	//"net/http"
	//"runtime"
	"github.com/astaxie/beego/logs"
	//"strategy_middle/logs"
	"strategy_middle/models"
	"strategy_middle/controller/interflow"
	"gopkg.in/mgo.v2/bson"
	//"time"
	//"Goroutine"
	//"os"
	//"bytes"
	"encoding/json"
	//"strconv"
	//"reflect"
	//"github.com/gin-gonic/gin"
	//"github.com/gorilla/websocket"
	"strategy_middle/constant"
)



/*首先根据登录名查找 策略配置表StrategyInfo 的所拥有 正在使用的策略，找到它的策略id和账户id，
然后根据账户id的账户资金，持仓pnl信息
接着根据策略id查找该策略所对应的 创建时间最大的 手续费，保证金，策略每天的信息表以及策略涉及的合约，
最后根据策略涉及的合约表的strategyId查找交易合约配置SymbolInfo，策略交易时间段SymbolTradingPeriod
*/
func init_data(name string)(strategyinfo []models.StrategyInfo,
						  strategysymbolmarginoverride [][]models.StrategySymbolMarginOverride,
						  strategysymbolfeeoverride [][]models.StrategySymbolFeeOverride,
						  strategypnldaily [][]models.StrategyPnlDaily,
						  strategysymbol [][]models.StrategySymbol,
						  accountinfo []models.AccountInfo,
						  accountpnldaily []models.AccountPnlDaily,
						  symbolinfo [][]models.SymbolInfo,
						  symboltradingperiod [][]models.SymbolTradingPeriod){

	st := models.ConnecToDB("StrategyInfo")
	us :=models.ConnecToDB("User")


	if name =="admin"{

		err := st.Find(nil).All(&strategyinfo)

		user :=models.User{}
		err = us.Find(bson.M{"user": name}).One(&user)

		inf := models.ConnecToDB("AccountInfo")
		pnl := models.ConnecToDB("AccountPnlDaily")

		err = inf.Find(nil).All(&accountinfo)
		err =pnl.Find(nil).All(&accountpnldaily)

		if err != nil {
		// 	logs.Fatal(err)
			logs.Error(err)
		}
	}else{

		err := st.Find(bson.M{"strategyArchiveOwner": name}).All(&strategyinfo)


		user :=models.User{}
		err = us.Find(bson.M{"user": name}).One(&user)

		inf := models.ConnecToDB("AccountInfo")
		pnl := models.ConnecToDB("AccountPnlDaily")

		err = inf.Find(bson.M{"accountId": user.AccountId}).All(&accountinfo)
		err =pnl.Find(bson.M{"accountId": user.AccountId}).All(&accountpnldaily)

		if err != nil {
		// 	logs.Fatal(err)
			logs.Error(err)
		}

	}



	marginoverride :=[]models.StrategySymbolMarginOverride{}
	feeoverride :=[]models.StrategySymbolFeeOverride{}
	str_pnldaily :=[]models.StrategyPnlDaily{}
	symbol:=[]models.StrategySymbol{}

	sinfo:=[]models.SymbolInfo{}
	tradingperiod:=[]models.SymbolTradingPeriod{}





	for _,e_str := range strategyinfo{
		// logs.Info("StrategyId:",e_str.StrategyId)
		// logs.Info("AccountId:",e_str.AccountId)

		mar := models.ConnecToDB("StrategySymbolMarginOverride")
		fee := models.ConnecToDB("StrategySymbolFeeOverride")
		daily := models.ConnecToDB("StrategyPnlDaily")
		sym := models.ConnecToDB("StrategySymbol")

		//.Sort("-createTime").Limit(100)
		err := mar.Find(bson.M{"strategyId": e_str.StrategyId}).Sort("-createTime").Limit(1).All(&marginoverride)
		err = fee.Find(bson.M{"strategyId": e_str.StrategyId}).Sort("-createTime").Limit(1).All(&feeoverride)
		err = daily.Find(bson.M{"strategyId": e_str.StrategyId}).All(&str_pnldaily)
		err = sym.Find(bson.M{"strategyId": e_str.StrategyId}).All(&symbol)





		//根据策略涉及的合约表的strategyId查找交易合约配置SymbolInfo，策略交易时间段SymbolTradingPeriod
		for _,sid := range symbol{
			sinf := models.ConnecToDB("SymbolInfo")
			per := models.ConnecToDB("SymbolTradingPeriod")

			err = sinf.Find(bson.M{"symbolId": sid.SymbolId}).All(&sinfo)
			err =per.Find(bson.M{"symbolId": sid.SymbolId}).All(&tradingperiod)

		}



		strategysymbolmarginoverride =append(strategysymbolmarginoverride,marginoverride)
		strategysymbolfeeoverride =append(strategysymbolfeeoverride,feeoverride)
		strategypnldaily =append(strategypnldaily,str_pnldaily)
		strategysymbol =append(strategysymbol,symbol)
		// accountinfo =append(accountinfo,acc_info)
		// accountpnldaily =append(accountpnldaily,acc_pnldaily)
		symbolinfo =append(symbolinfo,sinfo)
		symboltradingperiod =append(symboltradingperiod,tradingperiod)

		if err != nil {
			logs.Error(err)
		}

		// logs.Info("strategysymbolmarginoverride:", strategysymbolmarginoverride)
		// logs.Info("strategysymbolfeeoverride:", strategysymbolfeeoverride)
		// logs.Info("accountpnldaily:",accountpnldaily)
		// logs.Info("strategysymbol:",strategysymbol)
		// logs.Info("accountinfo:",accountinfo)
		// logs.Info("accountpnldaily:",accountpnldaily)

		// logs.Info("symbolinfo:",symbolinfo)
		// logs.Info("symboltradingperiod:",symboltradingperiod)

	}
	return
}

type initstruct struct{
	Strategyinfo []models.StrategyInfo `json:"StrategyInfo"`
	Strategysymbolmarginoverride [][]models.StrategySymbolMarginOverride `json:"StrategySymbolMarginOverride"`
	Strategysymbolfeeoverride [][]models.StrategySymbolFeeOverride `json:"StrategySymbolFeeOverride"`
	Strategypnldaily [][]models.StrategyPnlDaily `json:"StrategyPnlDaily"`
	Strategysymbol [][]models.StrategySymbol `json:"StrategySymbol"`
	Accountinfo []models.AccountInfo `json:"AccountInfo"`
	Accountpnldaily []models.AccountPnlDaily `json:"AccountPnlDaily"`
	Symbolinfo [][]models.SymbolInfo `json:"SymbolInfo"`
	Symboltradingperiod [][]models.SymbolTradingPeriod `json:"SymbolTradingPeriod"`
}


type user_login struct{
	User string `json:"user"`
	Password string `json:"password"`
}

type userlogin struct{
	Id int `json:"id"`
	Op string `json:"op"`
	Args  *user_login `json:"args"`
}
func (this *ApiOp)Send_init()(err error){//msg *models.Ws_send
	//logs.Info(reflect.TypeOf(message))
	//logs.Info(msg.Args)
	msg:=&userlogin{}
	err=json.Unmarshal(this.Message, msg)
	if err!=nil{
		this.Send_callback_fail(1001,"fail",constant.Errmap[1001],err)
		return
	}

	//把每个websocket客户端对应的用户id记录下来
	//根据User查找accountid
	//
	us := models.ConnecToDB("User")
	user :=&models.User{}
	err = us.Find(bson.M{"user":  msg.Args.User}).One(user)
	if err != nil {
    	this.Send_callback_fail(1003,"fail",constant.Errmap[1003],err)
    	return
	}
	
		//将所有客户端保存到字典中
	//////////////////////////////////
	interflow.ClientMgr.AddClient(this.Ws,user.AccountId)
	//interflow.Allws[this.Ws] =user.AccountId
	//////////////////////////////////
	strategyinfo,marginoverride,feeoverride,strategypnldaily,strategysymbol,
	accountinfo,accountpnldaily,symbolinfo,symboltradingperiod:=init_data(msg.Args.User)
    send_initdata:=models.Ws_recv{
    	Id:1,
    	Op:"initdata",
    	Code:0,
    	Msg:"success",
    	Data:initstruct{
    		Strategyinfo:strategyinfo,
    		Strategysymbolmarginoverride:marginoverride,
    		Strategysymbolfeeoverride:feeoverride,
    		Strategypnldaily:strategypnldaily,
    		Strategysymbol:strategysymbol,
    		Accountinfo:accountinfo,
    		Accountpnldaily:accountpnldaily,
    		Symbolinfo:symbolinfo,
    		Symboltradingperiod:symboltradingperiod,
    	},
    }

    s, err := json.Marshal(send_initdata)
    err = this.Ws.WriteMessage(this.Mt, []byte(s))
   
    if err!=nil{
   		logs.Error(err)
    }
	return
	////一下可以匹配到map[string]interface{}
	// for key, value := range msg.Args.(map[string]interface{}) {
	// 	if key =="user"{
	// 		logs.Info(reflect.TypeOf(value))

	// 	    switch vty := value.(type) {
	// 	    case string:
	// 	            val :=vty
	// 	            strategy,marginoverride,feeoverride,strategypnldaily,strategysymbol,
	//                 accountinfo,accountpnldaily,symbolinfo,symboltradingperiod:=init_data(val)

	// 	            send_initdata:=models.Ws_recv{
	// 	            	Id:1,
	// 	            	Op:"initdata",
	// 	            	Code:200,
	// 	            	Msg:"success",
	// 	            	Data:initstruct{
	// 	            		Strategy:strategy,
	// 	            		Strategysymbolmarginoverride:marginoverride,
	// 	            		Strategysymbolfeeoverride:feeoverride,
	// 	            		Strategypnldaily:strategypnldaily,
	// 	            		Strategysymbol:strategysymbol,
	// 	            		Accountinfo:accountinfo,
	// 	            		Accountpnldaily:accountpnldaily,
	// 	            		Symbolinfo:symbolinfo,
	// 	            		Symboltradingperiod:symboltradingperiod,
	// 	            	},
	// 	            }

	// 	            s, err := json.Marshal(send_initdata)
	// 	            err = ws.WriteMessage(mt, []byte(s))
		           
	// 	            if err!=nil{
	// 	           		logs.Info(err)
	// 	            }
		            
	// 	    default:
	// 	    		logs.Error("def:", vty)
	// 	    }

	// 		//val:=string(value)
	// 		//strategy,marginoverride,feeoverride,strategypnldaily,strategysymbol,
	// 		//accountinfo,accountpnldaily,symbolinfo,symboltradingperiod:=initdata(value)
	// 		//strategy,_,_,_,_,_,_,_,_:=initdata(val)
	// 		//RECV: {“id”:0, ”op”:”auth”, “code”:0, ”msg”:”success”, “data”:”” }
	// 		//s, err := json.Marshal(strategy)
	// 		//err = ws.WriteMessage(mt, []byte(s))
	// 	}

	// }


}