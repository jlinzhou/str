package models

// import (
// 	"gopkg.in/mgo.v2/bson"
// 	"strategy_middle/models"
// 	"time"
// )

// type User struct {
// 	ID       bson.ObjectId `bson:"_id"`
// 	UserName string        `bson:"username"`
// 	Summary  string        `bson:"summary"`
// 	Age      int           `bson:"age"`
// 	Phone    int           `bson:"phone"`
// 	PassWord string        `bson:"password"`
// 	Sex      int           `bson:"sex"`
// 	Name     string        `bson:"name"`
// 	Email    string        `bson:"email"`
// }

// type Diary struct {
// 	Uid        bson.ObjectId `bson:"uid"`
// 	ID         bson.ObjectId `bson:"_id"`
// 	CreatTime  time.Time     `bson:"creattime"`
// 	UpdateTime time.Time     `bson:"updatetime"`
// 	Title      string        `bson:"title"`
// 	Content    string        `bson:"content"`
// 	Mood       int           `bson:'Mood"`
// 	Pic        []string      `bson:'pic'`
// }

// func Register(password string, username string) (err error) {
// 	con := mgo.GetDataBase().C("user")
// 	//可以添加一个或多个文档
// 	/* 对应mongo命令行
// 	   db.user.insert({username:"13888888888",summary:"code",
// 	   age:20,phone:"13888888888"})*/
// 	err = con.Insert(&User{ID: bson.NewObjectId(), UserName: username, PassWord: password})
// 	return
// }

// func FindUser(username string) (User, error) {
// 	var user User
// 	con := mgo.GetDataBase().C("user")
// 	//通过bson.M(是一个map[string]interface{}类型)进行
// 	//条件筛选，达到文档查询的目的
// 	/* 对应mongo命令行
// 	   db.user.find({username:"13888888888"})*/
// 	if err := con.Find(bson.M{"username": username}).One(&user); err != nil {
// 		if err.Error() != mgo.GetErrNotFound().Error() {
// 			return user, err
// 		}

// 	}
// 	return user, nil
// }

// //通过uid查找本作者文章，并且显示文章作者名字
// func FindDiary(uid string) ([]interface{}, error) {
// 	con := mgo.GetDataBase().C("diary")
// 	// 其中的lookup功能可以实现类似于mysql中的join操作，方便于关联查询。
// 	/*对应mongo命令行
// 	  db.diary.aggregate([{$match:{uid: ObjectId("58e7a1b89b5099fdc585d370")}},
// 	  {$lookup{from:"user",localField:"uid",foreignField:"_id",as:"user"}},
// 	  {$project:{"user.name":1,title:1,content:1,mood:1}}]).pretty()
// 	*/
// 	pipeline := []bson.M{
// 		bson.M{"$match": bson.M{"uid": bson.ObjectIdHex(uid)}},
// 		bson.M{"$lookup": bson.M{"from": "user", "localField": "uid", "foreignField": "_id", "as": "user"}},
// 		bson.M{"$project": bson.M{"user.name": 1, "title": 1, "content": 1, "mood": 1, "creattime": 1}},
// 	}
// 	pipe := con.Pipe(pipeline)
// 	var data []interface{}
// 	err := pipe.All(&data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil

// }

// func ModifyDiary(id, title, content string) (err error) {
// 	con := mgo.GetDataBase().C("diary")
// 	//更新
// 	/*对应mongo命令行
// 	  db.diary.update({_id:ObjectId("58e7a1b89b5099fdc585d370")},
// 	   {$set:{title:"modify title",content:"modify content",
// 	   updatetime:new Date()})*/
// 	err = con.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"title": title, "content": content, "updatetime": time.Now().Add(8 * time.Hour)}})
// 	return

// }
