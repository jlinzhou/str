package logging

import(
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)


func init() {

	// logConf := logConfT{
    //     Filename: "D:/go/go_demo/logs/beego.log",
    //     Maxdays:  30,
    //     //Maxsize:  30 * mb,
    // }
	// configStr, err := json.Marshal(logConf)
	// //logs.SetLogger(logs.AdapterFile, string(b))
	// logs.SetLogger(logs.AdapterConsole)

    config := make(map[string]interface{})
    config["filename"] = "D:/go/go_demo/logs/middle.log"
	config["level"] = logs.LevelDebug
	config["maxsize"]=1024*500//相当于500kb
	config["maxdays"]=365
	/*

// "filename":"log/project.log" ：将日志保存到当前目录下的log目录下的project.log文件中
// "level":7 ：将日志级别设为7，也就是LevelDebug
// "maxlines":0 ：设置日志文件分割条件，若文件超过maxlines，则将日志保存到下个文件中，为0表示不设置
// "maxsize":0 ：设置日志文件分割条件，若文件超过maxsize，则将日志保存到下个文件中，为0表示不设置
// "daily":true：设置日志日否每天分割一次
// "maxdays":10：设置保存最近几天的日志文件，超过天数的日志文件被删除，为0表示不设置

	*/
    configStr, err := json.Marshal(config)
    if err != nil {
        fmt.Println("marshal failed,err:", err)
        return
	}
	
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	/*
	    logs.Debug("this is a test,my name is %s", "stu01")
    logs.Trace("this is a trace,my name is %s", "stu02")
	logs.Warn("this is a warn,my name is %s", "stu03")
	
Emergency          white
Alert              cyan
Critical           magenta
Error              red
Warning            yellow
Notice             green
Informational      blue
Debug       
	*/

}