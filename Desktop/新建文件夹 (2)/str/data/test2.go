package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Student struct {
	//Id_   bson.ObjectId `bson:"_id"`
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
	Email string `bson:"email"`
	Sex   string `bson:"sex"`
}

//数据库连接主要用到了mgo中的Dial()函数,连接形式如mgo.Dial(url1,url2,url3)
func ConnecToDB() *mgo.Collection {

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"192.168.18.16:27017"}, //数据库地址 dbhost: mongodb://user@123456:127.0.0.1:27017
		Timeout:   60 * time.Second,                // 连接超时时间 timeout: 60 * time.Second
		Source:    "test",                          // 设置权限的数据库 authdb: admin
		Username:  "jlzhou",                        // 设置的用户名 authuser: user
		Password:  "jlzhou",                        // 设置的密码 authpass: 123456
		PoolLimit: 1024,                            // 连接池的数量 poollimit: 100
	}

	session, err := mgo.DialWithInfo(dialInfo)

	//session, err := mgo.Dial("192.168.18.16:27017")
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	//session设置的模式分别为:
	//1.Strong
	//session 的读写一直向主服务器发起并使用一个唯一的连接，因此所有的读写操作完全的一致。
	//2.Monotonic
	//session 的读操作开始是向其他服务器发起（且通过一个唯一的连接），只要出现了一次写操作，
	//session 的连接就会切换至主服务器。由此可见此模式下，能够分散一些读操作到其他服务器，但是读操作不一定能够获得最新的数据。
	//3.Eventual
	//session 的读操作会向任意的其他服务器发起，多次读操作并不一定使用相同的连接，也就是读操作不一定有序。session 的写操作总是向主服务器发起，
	//但是可能使用不同的连接，也就是写操作也不一定有序。

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("student2")
	return c
}

//插入主要用到了函数 func (c *Collection) Insert(docs ...interface{}) error
func InsertToMogo() {
	c := ConnecToDB()
	stu1 := Student{
		Name:  "zhangsan",
		Phone: "13480989765",
		Email: "329832984@qq.com",
		Sex:   "F",
	}
	stu2 := Student{
		Name:  "liss",
		Phone: "13980989767",
		Email: "12832984@qq.com",
		Sex:   "M",
	}
	err := c.Insert(&stu1, &stu2)
	if err != nil {
		log.Fatal(err)
	}
}

//查询单个主要用到了func (c *Collection) Find(query interface{}) *Query函数，查询单个和多个主要用到了One()和Many()函数，条件组合可以查看mongDB数据库使用
func GetDataViaSex() {
	c := ConnecToDB()
	result := Student{}
	err := c.Find(bson.M{"sex": "M"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("student", result)
	students := make([]Student, 20)
	err = c.Find(nil).All(&students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(students)
}

//查询所有形如：c.Find(nil).Many(&results)
//另外，方法中也有个根据id来查询的方法 func (c *Collection) FindId(id interface{}) *Query，
func GetDataViaId() {
	id := bson.ObjectIdHex("5d1482570ddd3ee182df0359")
	c := ConnecToDB()
	stu := &Student{}
	err := c.FindId(id).One(stu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stu)
}

//更新通过函数
//*func (c *Collection) Update(selector interface{}, update interface{}) error
//*func (c *Collection) UpdateAll(selector interface{}, update interface{}) (info *ChangeInfo, err error)
//*func (c *Collection) UpdateId(id interface{}, update interface{}) error
func UpdateDBViaId() {
	//id := bson.ObjectIdHex("5a66a96306d2a40a8b884049")
	c := ConnecToDB()
	err := c.Update(bson.M{"email": "12832984@qq.com"}, bson.M{"$set": bson.M{"name": "haha", "phone": "37848"}})
	if err != nil {
		log.Fatal(err)
	}
}

//删除对应的方法
//
//func (c *Collection) Remove(selector interface{}) error]
//func (c *Collection) RemoveAll(selector interface{}) (info *ChangeInfo, err error)
//func (c *Collection) RemoveId(id interface{}) error
func RemoveFromMgo() {
	c := ConnecToDB()
	_, err := c.RemoveAll(bson.M{"phone": "13480989765"})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	InsertToMogo()
	//GetDataViaSex()
	//GetDataViaId()
	//UpdateDBViaId()
	//RemoveFromMgo()
}
