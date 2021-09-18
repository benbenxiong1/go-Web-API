package static

import (
	"blog/app/http/controllers"
	"blog/app/http/middleware"
	"blog/app/models"
	routers "blog/router"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type BenBen struct {

}

func (b *BenBen)Run(port string)  {
	http.ListenAndServe(port, b)
}

// ControllerMap 这里记录所有的应该注册的结构体
// 控制器map
var ControllerMap map[string]interface{}
// MiddlewareMap 中间件map
var MiddlewareMap map[string]interface{}

func init()  {
	//链接数据库
	models.DB = models.Link()

	ControllerMap = make(map[string]interface{})
	MiddlewareMap = make(map[string]interface{})
	// 给这两个map赋初始值 每次添加完一条路由或中间件，都要在此处把路由或者中间件注册到这里
	// 注册中间件
	MiddlewareMap["WebMiddleware"] =&middleware.WebMiddleware{}

	// 注册路由
	ControllerMap["Controller"] = &controllers.Controller{}
	ControllerMap["IndexController"] = &controllers.IndexController{}
}


func (b *BenBen)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	falg := false

	// 每一个http请求都会走到这里，然后在这里，根据请求的URL，为其分配所需要调用的方法
	params := []reflect.Value{reflect.ValueOf(w),reflect.ValueOf(r)}

	for _,v := range routers.RoutesList{
		fmt.Println(r.Method)
		fmt.Println(v.Method)
		if r.URL.Path == v.Path && r.Method == v.Method {
			falg = true
			//调用中间件 寻找路由及其调用中间件收尾功能

			//检测中间件是否存在
			for _,m := range v.Middleware{
				if mid,ok := MiddlewareMap[m]; ok{
					rmid := reflect.ValueOf(mid)
					params = rmid.MethodByName("Handle").Call(params) //执行中间件 返回value数据
					//判断中间件执行结果
					str := rmid.Elem().FieldByName("resString").String()
					if str != ""{
						status := rmid.Elem().FieldByName("status").Int()
						if status == 0 {
							status = 500
						}
						w.WriteHeader(int(status))
						fmt.Fprintf(w,str)
						return
					}
				}
			}
			// 检测成功，开始调用方法
			// 获取一个控制器包下的结构体
			if d,ok := ControllerMap[v.Controller];ok{
				reflect.ValueOf(d).MethodByName(v.Function).Call(params)
			}
			// 停止向后执行
			return
		}
	}
	// 如果路由列表中还是没有的话,去静态服务器中寻找
	if !falg {
		http.ServeFile(w,r,GetPublicPath()+ r.URL.Path)
	}
	return
}

// GetPublicPath 获取静态文件夹路径
func GetPublicPath() string {
	dir, err := os.Getwd()
	dir = strings.Replace(dir, "\\", "/", -1)
	Check(err)
	return dir + "/public/"
}

func Check(err error)  {
	if err != nil {
		fmt.Println(err)
		return
	}
}