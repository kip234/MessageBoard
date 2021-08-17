package DC

import (
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var DB *gorm.DB

const Addr="localhost:8081"

func Run(){
	MessageService:=new(Message)
	UserService:=new(User)
	rpc.Register(MessageService)
	rpc.Register(UserService)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", Addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}

func NewClient()*rpc.Client{
	client, err := rpc.DialHTTP("tcp",Addr)
	if err != nil {
		log.Fatal("dialing:", err)
		return nil
	}
	return client
}