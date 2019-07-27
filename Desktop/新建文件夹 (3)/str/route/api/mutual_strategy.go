//实现策略的增删改查功能，并和前端进行交互


package api

import(
	"net/http"
	"github.com/astaxie/beego/logs"
	"strategy_middle/models"
	//"reflect"
	//"gopkg.in/mgo.v2/bson"
	//"github.com/Unknwon/com"
	//"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func Get_Stra_DataAll() (strategy_infos  []models.StrategyInfo){
	c := models.ConnecToDB("strategy_info")
	err := c.Find(nil).All(&strategy_infos)
	if err != nil {

		logs.Error(err)
	}
	logs.Informational(strategy_infos)
	return
}

func GetALLInst(c *gin.Context) {
	/*
	   c.Query可用于获取?id=1这类URL参数，而c.DefaultQuery则支持设置一个默认值
	   code变量使用了e模块的错误编码，这正是先前规划好的错误码，方便排错和识别记录
	   util.GetPage保证了各接口的page处理是一致的
	   c *gin.Context是Gin很重要的组成部分，可以理解为上下文，它允许我们在中间件之间传递变量、管理流、验证请求的JSON和呈现JSON响应
	*/
	//id := c.Query("id")
	
	d := models.ConnecToDB("strategy_info")
	strategy_infos := make([]models.StrategyInfo, 20)
	err := d.Find(nil).All(&strategy_infos)
	if err != nil {
		logs.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": strategy_infos,
	})

}
func AddInst(c *gin.Context){
	// d := models.ConnecToDB("TdInfo")
	// aa:=models.TdInfo{}


	// err := d.Insert(&aa)
	// if err != nil {

	// 	logs.Error(err)
	// }

	d := models.ConnecToDB("User")
	aa1 :=models.User{}
	err := d.Insert(&aa1)
	if err != nil {

		logs.Error(err)
	}

	// d = models.ConnecToDB("Onorder")
	// aa2 :=models.Onorder{}
	// err = d.Insert(&aa2)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// d = models.ConnecToDB("Base_comm")
	// aa3 :=models.Base_comm{}
	// err = d.Insert(&aa3)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// d = models.ConnecToDB("Mdlnfo")
	// aa4 :=models.Mdlnfo{}
	// err = d.Insert(&aa4)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// d := models.ConnecToDB("SymbolInfo")
	// aa5 :=models.SymbolInfo{}
	// err := d.Insert(&aa5)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// 	d = models.ConnecToDB("SymbolTradingPeriod")
	// aa6 :=models.SymbolTradingPeriod{}
	// err = d.Insert(&aa6)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// 	d = models.ConnecToDB("SymbolMarginDefault")
	// aa7 :=models.SymbolMarginDefault{}
	// err = d.Insert(&aa7)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// 	d = models.ConnecToDB("SymbolFeeDefault")
	// aa8 :=models.SymbolFeeDefault{}
	// err = d.Insert(&aa8)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// 	d = models.ConnecToDB("AccountInfo")
	// aa9 :=models.AccountInfo{}
	// err = d.Insert(&aa9)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// 	d = models.ConnecToDB("AccountPnlDaily")
	// aa10 :=models.AccountPnlDaily{}
	// err = d.Insert(&aa10)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// 		d = models.ConnecToDB("Strategy")
	// aa11 :=models.Strategy{}
	// err = d.Insert(&aa11)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// 		d = models.ConnecToDB("StrategySymbol")
	// aa12 :=models.StrategySymbol{}
	// err = d.Insert(&aa12)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// 		d = models.ConnecToDB("StrategySymbolMarginOverride")
	// aa13 :=models.StrategySymbolMarginOverride{}
	// err = d.Insert(&aa13)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// 		d = models.ConnecToDB("StrategySymbolFeeOverride")
	// aa14 :=models.StrategySymbolFeeOverride{}
	// err = d.Insert(&aa14)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// d = models.ConnecToDB("StrategyPnlDaily")
	// aa15 :=models.StrategyPnlDaily{}
	// err = d.Insert(&aa15)
	// if err != nil {

	// 	logs.Error(err)
	// }

	// d = models.ConnecToDB("StrategyArchive")
	// aa16 :=models.StrategyArchive{}
	// err = d.Insert(&aa16)
	// if err != nil {

	// 	logs.Error(err)
	// }


	// //result := make([]models.Strategy_info,20)
	 
	// mydata := d.Find(nil).Sort(bson.M{"id","1"})
	// t:=reflect.TypeOf(mydata)
	// logs.Info(t)
	 
	// // for i:=0;i<t.NumField();i++{
	// // 	f:=t.Field(i)
	// // 	val :=v.Field(i).Interface()
	// // 	logs.Info(f.Name,f.Type,val)
	// // }
	// //t := reflect.TypeOf(mydata) 
	// if mydata != nil {
	// 	logs.Info(mydata)
	// 	logs.Info(reflect.ValueOf(mydata))
	// }
	// 	c.JSON(http.StatusOK, gin.H{
	// 	"data": mydata,
	// })
	// //logs.Info(result)
	// // .Sort("-age")
	// // stra := models.Strategy_info{
	// // 	Instance_name:  "zhangsan",
	// // 	Phone: "13480989765",
	// // 	Email: "329832984@qq.com",
	// // 	Sex:   "F",
	// // }

}

// func InsertToMogo() {
// 	c := ConnecToDB("student2")
// 	stu1 := Student{
// 		Name:  "zhangsan",
// 		Phone: "13480989765",
// 		Email: "329832984@qq.com",
// 		Sex:   "F",
// 	}
// 	stu2 := Student{
// 		Name:  "liss",
// 		Phone: "13980989767",
// 		Email: "12832984@qq.com",
// 		Sex:   "M",
// 	}
// 	err := c.Insert(&stu1, &stu2)
// 	if err != nil {

// 		logs.Fatal(err)
// 	}
// }