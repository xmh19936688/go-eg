package api

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var cacheMap sync.Map

type kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func set(c *gin.Context) {
	var form kv
	if err := binding.JSON.Bind(c.Request, &form); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	cacheMap.Store(form.Key, form.Value)
	c.JSON(http.StatusOK, "success")
}

func get(c *gin.Context) {
	var form kv
	if err := binding.JSON.Bind(c.Request, &form); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	value, ok := cacheMap.Load(form.Key)
	if !ok {
		c.JSON(http.StatusBadRequest, "not found")
		return
	}

	c.JSON(http.StatusOK, value)
}

func Init(gs *gin.Engine) {
	gs.GET("/", index)
	gs.POST("/set", set)
	gs.POST("/get", get)
}
