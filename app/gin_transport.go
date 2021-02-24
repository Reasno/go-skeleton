package app

import (
	"github.com/DoNewsCode/core/ginmw"
	"github.com/gin-gonic/gin"
	"github.com/nfangxu/core-skeleton/app/handlers"
	"net/http"
)

type GinTransport struct {
	http.Handler
}

func NewGinTransport(b handlers.UsersHandler) GinTransport {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(ginmw.WithContext())
	r.GET("/", b.Find)

	return GinTransport{Handler: r}
}
