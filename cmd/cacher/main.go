package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"

	cacherApi "xmh.go-eg/api/cacher"
	"xmh.go-eg/common"
	"xmh.go-eg/setting"
)

func main() {
	log.Println("starting cacher, version:", setting.Version)
	defer log.Println("bye")
	defer log.Println("tear down")

	gin.SetMode(gin.DebugMode)
	gs := gin.Default()
	gs.LoadHTMLGlob("static/*")
	gs.Static("/static", "static")

	cacherApi.Init(gs)
	srv := common.RunServer(gs)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, os.Kill)
	<-interrupt

	common.StopServer(srv)
}
