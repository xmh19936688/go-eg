package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"

	uiApi "xmh.go-eg/api/ui"
	"xmh.go-eg/common"
)

func main() {
	log.Println("starting cacher")
	defer log.Println("bye")
	defer log.Println("tear down")

	gin.SetMode(gin.DebugMode)
	gs := gin.Default()
	gs.LoadHTMLGlob("static/*")
	gs.Static("/static", "static")

	uiApi.Init(gs)
	srv := common.RunServer(gs)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, os.Kill)
	<-interrupt

	common.StopServer(srv)
}
