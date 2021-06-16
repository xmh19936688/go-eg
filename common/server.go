package common

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"xmh.go-eg/setting"
)

func RunServer(gs *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":" + setting.Config.App.Port,
		Handler: gs,
	}

	go func(srv *http.Server) {
		log.Println("listening:", setting.Config.App.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln("failed to serve:" + err.Error())
		}
	}(srv)

	return srv
}

func StopServer(srv *http.Server) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	srv.Shutdown(ctx)
}
