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

func main() {

	sendchanneL_name, jsonfile := mutual_channel()
	Ioutil(sendchanneL_name, jsonfile)
	rabbitmq.Recv(sendchanneL_name)
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

			rabbitmq.Send(sendchanneL_name, msg)
		}
	}

}
