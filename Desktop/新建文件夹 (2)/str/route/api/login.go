package api

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"strategy_middle/setting"
	"github.com/astaxie/beego/logs"
	 "net"
    "strings"
    //"reflect"
	"strategy_middle/models"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

func getip()(ip string) {
    conn, err := net.Dial("udp", "google.com:80")
    if err != nil {
        logs.Error(err.Error())
        return
    }
    defer conn.Close()
    ip =strings.Split(conn.LocalAddr().String(), ":")[0]
    return ip
}


func accountmatch(name string,password string)(err error){
	us:=models.ConnecToDB("User")
	user:=&models.User{}
	err=us.Find(bson.M{"user":name}).One(user)
	if err!=nil{
		return
	}
	if password!=user.Password {
		err=errors.New("not match")
	}
	return
}
func Login(c *gin.Context) {
	//name := c.Request.FormValue("Name")
	//passwd := c.Request.FormValue("Passwd")
	name := c.PostForm("Name")
	password := c.PostForm("Passwd")
	err:=accountmatch(name,password)
	// logs.Informational(reflect.TypeOf(name),name)
	// logs.Informational(password)

	if err !=nil{
		c.String(502,"账户或密码错误!")
	}else{

		c.HTML(200, "index.html", gin.H{
			"PORT":         setting.HTTPPort,
			"IP":           getip(),
			"username":         name,
			"password":     password,

		})
	}

	// //用于临时存储用户登录信息的Map
	// var State = make(map[string]interface{})
	// var user []User

	// if DB.Where("username = ? AND password=?", name, password).First(&user).RowsAffected == 0 {
	// 	State["state"] = 0
	// 	State["text"] = "账号或密码错误！"
	// 	c.String(http.StatusOK, "%v", State)
	// } else {
	// 	//var onorder []Onorder
	// 	//onorderdata := DB.Find(&onorder)

	// 	//fmt.Printf("%t\n", onorderdata)
	// 	//for a, b := range onorderdata {
	// 	//	fmt.Println(a)
	// 	//}
	// 	var strategy_info []Strategy_info
	// 	strategydata := DB.Find(&strategy_info)

	// 	c.HTML(200, "index.html", gin.H{
	// 		"strategydata": strategydata,
	// 	})
	// }
	//if err := DB.Where("username = ?", name).First(&user).Error; err != nil {
	//	State["state"] = 0
	//	State["text"] = "密码错误！"
	//	c.String(http.StatusOK, "%v", State)
	//} else {
	//	State["state"] = 1
	//	State["text"] = "登录成功！"
	//dd := DB.Where("username = ?", name).First(&user)
	//fmt.Println(typeof(DB.Where("username = ?", "jinzhu").First(&user)))
	//	fmt.Println(DB.Where("username = ?", name).First(&user).RecordNotFound())
	//	fmt.Println(DB.Where("username = ?", "jlzhou").First(&user).Value)
	//	var onorder []Onorder
	//	onorderdata := DB.Find(&onorder)
	//c.JSON(200, user)
	//c.String(http.StatusOK, "%v", State)
	//	c.HTML(200, "index.html", gin.H{
	//		"onorder": onorderdata,
	//	})
	//}

}
