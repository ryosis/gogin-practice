package main

import (
	"net/http"
	"path/filepath"
	"runtime"

	"gogin-practice/controller"
	"gogin-practice/middleware"
	"gogin-practice/service"

	"github.com/gin-gonic/gin"
)

func main() {
	service.Init(getPath())
	service.InitEnv()

	router := gin.Default()
	router.Use(middleware.CORS())
	controller.GetData(router, service.CFG.Db)
	controller.SetData(router, service.CFG.Db)

	http.ListenAndServe(":8080", router)
}

func getPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
