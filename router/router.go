package routers

import (
	"fmt"
	"strings"
)

var RoutesList []route         // 全部路由列表

type route struct {
	Path       string   // 路径
	Target     string   // 对应的控制器路径 Controller@index 这样的方法
	Method     string   // 访问类型 是get post 或者其他
	Alias      string   // 路由的别名
	Middleware []string // 中间件名称
	Controller string   // 控制器名称
	Function   string   // 挂载到控制器上的方法名称
}

//暂不做路由分组
//type route_group struct {
//	root_path   string   // 路径
//	root_target string   // 对应的控制器路径 Controller@index 这样的方法
//	alias       string   // 路由的别名
//	middleware  []string // 中间件名称
//	routes      []route  // 包含的路由
//}



func init() {
	for _,router1 := range Routes {
		router1.Method = strings.ToTitle(router1.Method)
		// 把target拆分成控制器和方法
		cf := strings.Split(router1.Target,"@")
		if len(cf)==2 {
			router1.Controller = cf[0]
			router1.Function = cf[1]
		}else{
			fmt.Println("Target格式错误！"+router1.Target)
			return
		}

		fmt.Println(router1)
		RoutesList = append(RoutesList, router1)
	}
}