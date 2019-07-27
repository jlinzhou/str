package api
import(
	"github.com/astaxie/beego/logs"
	"fmt"
	"github.com/gorilla/websocket"
	"strategy_middle/constant"
)

type ApiOp struct {
	Op string
	Message []uint8
	Ws *websocket.Conn
	Mt int
}

func (p *ApiOp)Send_callback_fail(code int, msg string,reason error,err error){
	
	s:=`{"id":%d,"op":"%s","code":%d,"msg":"%s","data":{"reason":"%s","error":"%s"}}`
	s =fmt.Sprintf(s,constant.Allapi[p.Op].Id,p.Op,code,msg,reason,err)
	err = p.Ws.WriteMessage(p.Mt, []byte(s))
	if err!=nil{
		logs.Error(err)
	}
}
func (p *ApiOp)Send_callback_success(){
	s:=`{"id":%d,"op":"%s","code":%d,"msg":"%s"}`
	s =fmt.Sprintf(s,constant.Allapi[p.Op].Id,p.Op,0,"success")
	err := p.Ws.WriteMessage(p.Mt, []byte(s))
	if err!=nil{
		logs.Error(err)
	}
	
}