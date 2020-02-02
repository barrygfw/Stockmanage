package main

import (
	"fmt"
	"graduationProjectPeng/db"
	"graduationProjectPeng/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"

	"graduationProjectPeng/pkg/setting"
	"graduationProjectPeng/routers"
)

func main() {
	setting.Setup()
	db.Setup()
	logging.Setup()

	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	routers.InitRouter(server)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        server,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
