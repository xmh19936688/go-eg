package auther

import (
	"net/http"
	"xmh.go-eg/setting"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"xmh.go-eg/model"
)

func auth(c *gin.Context) {
	var form model.Identity
	if err := binding.JSON.Bind(c.Request, &form); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input")
		return
	}
	if len(form.Password) == 0 && len(form.Token) == 0 {
		c.JSON(http.StatusUnauthorized, "invalid input")
		return
	}

	if len(form.Password) > 0 && form.Password != setting.Config.Auth.Secret {
		c.JSON(http.StatusUnauthorized, "invalid password")
		return
	}
	if len(form.Token) > 0 && form.Token != setting.Config.Auth.Secret {
		c.JSON(http.StatusUnauthorized, "invalid token")
		return
	}

	c.JSON(http.StatusOK, model.Identity{Token: setting.Config.Auth.Secret})
}

func Init(gs *gin.Engine) {
	gs.POST("/auth", auth)
}
