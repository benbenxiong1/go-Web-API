package controllers

import (
	"blog/app/models"
	"fmt"
	"net/http"
)

type IndexController struct {
	Controller
}

type Todos []models.UserInfo

func (index *IndexController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	todos := Todos{
		models.UserInfo{Name: "Write presentation",Email: "123@qq.com"},
		models.UserInfo{Name: "Host meetup",Email: "456@qq.com"},
	}
	models.DB.Create(&todos)
	index.success(w,todos)
	return
	index.msg(w,"1111")

}

func (index *IndexController) Error(w http.ResponseWriter, r *http.Request) {
	index.error(w,"错误拉",1001)
}

func (index *IndexController) Msg(w http.ResponseWriter, r *http.Request) {
	index.msg(w,"1111")
}