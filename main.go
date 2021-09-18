package main

import "blog/static"

func main()  {
	benben := new(static.BenBen)
	benben.Run(":9090")
	//http.Handle('/',IndexController.Index)
	//models2.DB = models2.Link()
	//router := routers.NewRouter()
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
}