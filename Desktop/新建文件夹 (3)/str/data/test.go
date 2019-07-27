package main

import (
	"fmt"
	"io/ioutil"
	"strategy_middle/rabbitmq"
	"strings"
)

func mutual_channel() (string, string) {
	a := "strategy_send"
	b := "json_file.txt"
	return a, b
}

type TestPro struct {
	msgContent string
}

// 实现发送者
func (t *TestPro) MsgContent() string {
	return t.msgContent
}

// 实现接收者
func (t *TestPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	return nil
}

func main() {
	msg := fmt.Sprintf("这是测试任务")
	t := &TestPro{
		msg,
	}
	queueExchange := &rabbitmq.QueueExchange{
		"test.rabbit",
		"rabbit.key",
		"test.rabbit.mq",
		"direct",
	}
	mq := rabbitmq.New(queueExchange)
	mq.RegisterProducer(t)
	mq.RegisterReceiver(t)
	mq.Start()
	//sendchanneL_name, jsonfile := mutual_channel()
	//Ioutil(sendchanneL_name, jsonfile)
}

func Ioutil(sendchanneL_name string, jsonfile string) {
	if contents, err := ioutil.ReadFile(jsonfile); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		//fmt.Println(string(contents))
		//result := strings.Replace(string(contents), "\n", "", 1)
		//fmt.Println(result)
		json_list := strings.Split(string(contents), "\n")
		for _, msg := range json_list {
			fmt.Println(msg)

			//rabbitmq_test.Send(sendchanneL_name, msg)
		}
	}

}
