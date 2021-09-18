package routers

type routes []route  //单路由
//type routesGroups []route_group // 路由组集合

var Routes = routes{
	route{
		Path:       "/",
		Target:     "IndexController@Index",
		Method:     "get",
		Alias:      "index.index",
	},
	route{
		Path:       "/error",
		Target:     "IndexController@Error",
		Method:     "get",
		Alias:      "index.error",
	},
	route{
		Path:       "/msg",
		Target:     "IndexController@Msg",
		Method:     "get",
		Alias:      "index.msg",
	},
}

//路由分组 暂时不做
//var RoutesGroups = routesGroups{
//	route_group{
//		root_path:   "/",
//		root_target: "",
//		alias:       "",
//		middleware:  nil,
//		routes:	routes{
//			route{
//				path:       "/",
//				target:     "IndexController@index",
//				method:     "post",
//				alias:      "index.index",
//				middleware: nil,
//				controller: "IndexController",
//				function:   "",
//			},
//		},
//	},
//}







