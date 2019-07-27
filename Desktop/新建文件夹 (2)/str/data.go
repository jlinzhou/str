package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Auth_args struct {
// 	User string `json:user`
// 	Pwd  string `json:Pwd`
// }
// type Auth struct {
// 	Id   int        `json:id`
// 	Op   string     `json:op`
// 	Args *Auth_args `json:args`
// }

// type Strategy_args struct {
// 	Strategy_id   int    `json:strategy_id`
// 	Strategy_name string `json:strategy_name`
// 	Action        string `json:action`
// 	Account       string `json:account`
// }
// type Strategy struct {
// 	Id   int            `json:id`
// 	Op   string         `json:op`
// 	Args *Strategy_args `json:args`
// }

// var auth = &Auth{Id: 1, Op: "auth", Args: &Auth_args{User: "myid", Pwd: "123456"}}

// var strategy = &Strategy{Id: 1, Op: "strategy", Args: &Strategy_args{Strategy_id: 1, Strategy_name: "stop", Action: "m_stop", Account: "test"}}

// func main() {
// 	jsonauth, _ := json.Marshal(auth)
// 	jsonstrategy, _ := json.Marshal(strategy)
// 	fmt.Println(string(jsonauth))
// 	fmt.Println(string(jsonstrategy))

// }
