package main

import (
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"net/http"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := routers.InitRouter()

	s := &http.Server{
		Addr : fmt.Sprintf(":%d", setting.HTTPPORT),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}