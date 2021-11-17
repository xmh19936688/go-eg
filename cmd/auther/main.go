package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"

	autherApi "xmh.go-eg/api/auther"
	"xmh.go-eg/common"
	"xmh.go-eg/setting"
)

func main() {
	log.Println("starting auther, version:", setting.Version)
	defer log.Println("bye")
	defer log.Println("tear down")

	gin.SetMode(gin.DebugMode)
	gs := gin.Default()

	autherApi.Init(gs)
	srv := common.RunServer(gs)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, os.Kill)
	<-interrupt

	common.StopServer(srv)
}
