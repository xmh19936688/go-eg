package cacher

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"xmh.go-eg/model"
	"xmh.go-eg/setting"
)

var cacheMap sync.Map

func auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	log.Println("auth-header:", token)
	var body = model.Identity{
		Token: token,
	}
	bs, err := json.Marshal(body)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, "invalid token")
		c.Abort()
		return
	}

	resp, err := http.Post(setting.Config.Auth.Url, "application/json", bytes.NewReader(bs))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "auther error: "+err.Error())
		c.Abort()
		return
	}
	if resp.StatusCode > 200 {
		c.JSON(http.StatusUnauthorized, "invalid token")
		c.Abort()
		return
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func set(c *gin.Context) {
	var form model.KV
	if err := binding.JSON.Bind(c.Request, &form); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	cacheMap.Store(form.Key, form.Value)
	c.JSON(http.StatusOK, "success-beta")
}

func get(c *gin.Context) {
	var form model.KV
	if err := binding.JSON.Bind(c.Request, &form); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	value, ok := cacheMap.Load(form.Key)
	if !ok {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	c.JSON(http.StatusOK, value.(string)+"-beta")
}

func Init(gs *gin.Engine) {
	gs.GET("/", index)
	gs.POST("/set", auth, set)
	gs.POST("/get", auth, get)
}
