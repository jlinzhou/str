package models

// import (
// 	//"fmt"
// 	"net/http"
// 	//"reflect"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jinzhu/gorm"
// 	"strategy_middle/setting"
// )

// var DB *gorm.DB
// var err error

// func InitDB() (*gorm.DB, error) {

// 	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
// 	DB.SingularTable(true)
// 	//注意gorm查找struct名对应数据库中的表名的时候会默认把你的struct中的大写字母转换为小写并加上“s”，
// 	//所以可以加上DB.SingularTable(true) 让grom转义struct名字的时候不用加上s。
// 	return nil, err
// }

// //登录
// func Login(c *gin.Context) {
// 	//name := c.Request.FormValue("Name")
// 	//passwd := c.Request.FormValue("Passwd")
// 	name := c.PostForm("Name")
// 	password := c.PostForm("Passwd")
// 	//fmt.Println(name)
// 	//fmt.Println(password)
// 	logging.Info(name)
// 	logging.Info(password)
// 	//用于临时存储用户登录信息的Map
// 	var State = make(map[string]interface{})
// 	var user []User

// 	if DB.Where("username = ? AND password=?", name, password).First(&user).RowsAffected == 0 {
// 		State["state"] = 0
// 		State["text"] = "账号或密码错误！"
// 		c.String(http.StatusOK, "%v", State)
// 	} else {
// 		//var onorder []Onorder
// 		//onorderdata := DB.Find(&onorder)

// 		//fmt.Printf("%t\n", onorderdata)
// 		//for a, b := range onorderdata {
// 		//	fmt.Println(a)
// 		//}
// 		var strategy_info []Strategy_info
// 		strategydata := DB.Find(&strategy_info)

// 		c.HTML(200, "index.html", gin.H{
// 			"strategydata": strategydata,
// 		})
// 	}
// 	//if err := DB.Where("username = ?", name).First(&user).Error; err != nil {
// 	//	State["state"] = 0
// 	//	State["text"] = "密码错误！"
// 	//	c.String(http.StatusOK, "%v", State)
// 	//} else {
// 	//	State["state"] = 1
// 	//	State["text"] = "登录成功！"
// 	//dd := DB.Where("username = ?", name).First(&user)
// 	//fmt.Println(typeof(DB.Where("username = ?", "jinzhu").First(&user)))
// 	//	fmt.Println(DB.Where("username = ?", name).First(&user).RecordNotFound())
// 	//	fmt.Println(DB.Where("username = ?", "jlzhou").First(&user).Value)
// 	//	var onorder []Onorder
// 	//	onorderdata := DB.Find(&onorder)
// 	//c.JSON(200, user)
// 	//c.String(http.StatusOK, "%v", State)
// 	//	c.HTML(200, "index.html", gin.H{
// 	//		"onorder": onorderdata,
// 	//	})
// 	//}

// }

// //func GetProjects(c *gin.Context) {
// //	var onorder []Onorder
// //	if err := DB.Find(&onorder).Error; err != nil {
// //		c.AbortWithStatus(404)
// //		fmt.Println(err)
// //	} else {
// //		//c.JSON(200, people)
// //		gin.H(onorder)
// //	}
// //}

// func GetPerson(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var person Person
// 	if err := DB.Where("id = ?", id).First(&person).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		logging.Error(err)
// 	} else {
// 		c.JSON(200, person)
// 	}
// }

// func CreatePerson(c *gin.Context) {
// 	var person Person
// 	c.BindJSON(&person)
// 	DB.Create(&person)
// 	c.JSON(200, person)
// }

// func UpdatePerson(c *gin.Context) {
// 	var person Person
// 	id := c.Params.ByName("id")
// 	if err := DB.Where("id = ?", id).First(&person).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		logging.Error(err)
// 		//fmt.Println(err)
		
// 	}
// 	c.BindJSON(&person)
// 	DB.Save(&person)
// 	c.JSON(200, person)
// }

// func DeletePerson(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var person Person
// 	d := DB.Where("id = ?", id).Delete(&person)
// 	logging.Info(d)
// 	c.JSON(200, gin.H{"id #" + id: "deleted"})
// }
