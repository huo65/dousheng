package main

import (
	"douSheng/config"
	"douSheng/router"
	"fmt"
)

func main() {
	//fmt.Println("非常抱歉我们学校开学就有各种考试，压力太大，今年暑假我们一定报名并且交出一个完整的项目")
	r := router.InitDouyinRouter()
	err := r.Run(fmt.Sprintf(":%d", config.Info.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}
