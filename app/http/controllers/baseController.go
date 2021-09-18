package controllers

import (
	"encoding/json"
	"net/http"
)

// Controller 控制器，处理数据的逻辑
type Controller struct{}

type BingoStruct struct {
	Status int  `json:"status"` // 响应的状态码
	Msg string `json:"msg"`
	Data interface{} `json:"data"` // 响应信息
}

//返回信息
func (c *Controller)success(w http.ResponseWriter,data interface{})  {
	b := BingoStruct{Status: 200,Data:data,Msg: "操作成功"}
	json.NewEncoder(w).Encode(b)
}

//错误返回
func (c *Controller)error(w http.ResponseWriter,msg string,status int)  {
	b := BingoStruct{Status: status,Data:nil,Msg: msg}
	json.NewEncoder(w).Encode(b)
}

//返回提示信息
func (c *Controller)msg(w http.ResponseWriter,msg string)  {
	b := BingoStruct{Status: 200,Data:nil,Msg: msg}
	json.NewEncoder(w).Encode(b)
}


