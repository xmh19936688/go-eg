package ui

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"xmh.go-eg/setting"

	"github.com/gin-gonic/gin"
)

var cacheMap sync.Map

func auth(c *gin.Context) {
	u, _ := url.Parse(setting.Config.Auth.Url)
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL = u
			fmt.Println("req.URL:", req.URL)
			fmt.Println("req:", *req)
			req.Host = u.Host
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			c.JSON(http.StatusUnauthorized, "auther error: "+err.Error())
		},
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func set(c *gin.Context) {
	u, _ := url.Parse(setting.Config.Cache.Url)
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = u.Host
			req.URL.Path = c.Request.URL.Path
			req.Host = u.Host
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			c.JSON(http.StatusInternalServerError, "cacher error: "+err.Error())
		},
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func get(c *gin.Context) {
	u, _ := url.Parse(setting.Config.Cache.Url)
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = u.Host
			req.URL.Path = c.Request.URL.Path
			req.Host = u.Host
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			c.JSON(http.StatusInternalServerError, "cacher error: "+err.Error())
		},
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func Init(gs *gin.Engine) {
	gs.GET("/", index)
	gs.POST("/auth", auth)
	gs.POST("/set", set)
	gs.POST("/get", get)
}
