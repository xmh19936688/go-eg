package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"xmh.go-eg/api"
	"xmh.go-eg/setting"

	"github.com/gin-gonic/gin"
)

func run(gs *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":" + setting.Config.App.Port,
		Handler: gs,
	}

	go func(srv *http.Server) {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln("failed to serve:" + err.Error())
		}
	}(srv)

	return srv
}

func stop(srv *http.Server) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	srv.Shutdown(ctx)
}

func main() {
	fmt.Println("starting cacher")
	defer fmt.Println("bye")
	defer fmt.Println("tear down")

	gin.SetMode(gin.DebugMode)
	gs := gin.Default()
	gs.LoadHTMLGlob("static/*")
	gs.Static("/static", "static")

	api.Init(gs)
	srv := run(gs)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, os.Kill)
	<-interrupt

	stop(srv)
}
