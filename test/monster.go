package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// Monster 绑定方法Store 将一个Monstar变量序列化到文件中
func (monster *Monster) Store() bool {
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("marshal err=", err)
		return false
	}
	// 保存到文件
	filepath := "./json.txt"
	err = ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		fmt.Println("write file err = ", err)
		return false
	}
	return true
}

// 给Monster绑定方法ReStore 将一个序列化的Monster，从文件中读取
// 反序列化为Monster对象，检查反序列化 名字正确
func (monster *Monster) ReStore() bool {
	filepath := "./json.txt"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("ReadFile err=", err)
		return false
	}
	err = json.Unmarshal(data, monster)
	if err != nil {
		fmt.Println("UnMarshal err=", err)
		return false
	}
	return true
}
