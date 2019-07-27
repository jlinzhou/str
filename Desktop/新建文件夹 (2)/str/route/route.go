package route

import (
	//"fmt"
	"net/http"
	"strategy_middle/controller/ws"
	//"strategy_middle/models"
	"strategy_middle/route/api"
	"strategy_middle/setting"
	//"runtime"

	"github.com/gin-gonic/gin"
	//"strategy_middle/models"
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	//"github.com/gin-contrib/sessions/mongo"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
)

//var allws

// Init 路由初始化

// func Init() {

// 	db, err := models.InitDB()

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()

// 	//db.AutoMigrate(&Person{})

// 	router := gin.Default()
// 	router.GET("/ping", ws.Ping)
// 	//router.GET("g/", models.GetProjects)
// 	//router.LoadHTMLFiles("templates/index.html")

// 	router.LoadHTMLGlob("templates/*")

// 	// 静态资源加载，本例为css,js以及资源图片
// 	router.StaticFS("/public", http.Dir("D:/go/go_demo/src/go-websocket/static"))

// 	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML(200, "login.html", nil)
// 	})
// 	router.POST("/index", models.Login)
// 	router.GET("index/", func(c *gin.Context) {
// 		c.HTML(200, "index.html", nil)
// 	})
// 	router.Run(":3000")
// }

func InitRouter() *gin.Engine {

	r := gin.New()

	//r.Use(gin.Logger())

	r.Use(gin.Recovery())




	// store := cookie.NewStore([]byte("secret"))
	// r.Use(sessions.Sessions("mysession", store))




	gin.SetMode(setting.RunMode)

	r.GET("/ping", ws.Ping)
	//router.GET("g/", models.GetProjects)
	//router.LoadHTMLFiles("templates/index.html")

	r.LoadHTMLGlob("templates/*")

	// 静态资源加载，本例为css,js以及资源图片

	r.StaticFS("/public", http.Dir(setting.Static_file))

	r.StaticFile("/favicon.ico", "./resources/favicon.ico")




	r.GET("/", func(c *gin.Context) {
		// session := sessions.Default(c)
		// var count int
		// v := session.Get("count")
		// if v == nil {
		// 	count = 0
		// } else {
		// 	count = v.(int)
		// 	count++
		// }
		// session.Set("count", count)
		// session.Save()
		c.HTML(200, "login.html", nil)
	})
	r.POST("/index", api.Login)
	// r.GET("index/", func(c *gin.Context) {
	// 	c.HTML(200, "index.html", nil)
	// })

	//路由设置
	apiv1 := r.Group("/instance")
	{
		//获取所有
		apiv1.GET("/", api.GetALLInst)
		// //新建
		apiv1.GET("/aa", api.AddInst)
		// //更新
		// apiv1.PUT("/instance/:id", v1.EditInst)
		// //删除
		// apiv1.DELETE("/instance/:id", v1.DeleteInst)
	}

	return r
}
