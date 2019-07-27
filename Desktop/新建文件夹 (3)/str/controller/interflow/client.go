package interflow
import(
	"github.com/gorilla/websocket"
)
type Client struct{
	onlineUsers map[*websocket.Conn]uint64 
}

var (
	ClientMgr *Client
)
 func init(){
	ClientMgr =&Client{
		onlineUsers:make(map[*websocket.Conn]uint64,1024),
	}
 }

 func (p *Client)AddClient(conn *websocket.Conn,userid uint64){
	p.onlineUsers[conn] = userid
 }
 func (p *Client)GetClient(userid uint64 )(conn *websocket.Conn,err error){
	err=nil
	for k,v := range p.onlineUsers{
		if userid==v{
			conn=k
			break
		}
	}
	return
 }

func (p *Client)GetAllUsers()(map[*websocket.Conn]uint64){
	return p.onlineUsers
}

 func (p *Client)DelClient(conn *websocket.Conn){
	delete(p.onlineUsers, conn)
 }
// func (p *Client)Addclient()(err error){

// }