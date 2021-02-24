package handlers

import (
	"github.com/DoNewsCode/core/srvhttp"
	"github.com/DoNewsCode/core/unierr"
	"github.com/gin-gonic/gin"
	"github.com/nfangxu/core-skeleton/app/repositories"
)

type UsersHandler struct {
	Repository *repositories.UserRepository
}

func (h UsersHandler) Find(c *gin.Context) {
	encoder := srvhttp.NewResponseEncoder(c.Writer)
	id := c.Query("user_id")
	val, err := h.Repository.Find(c, id)
	if err != nil {
		encoder.EncodeError(unierr.NotFoundErr(err, "user %s not found", id))
		return
	}
	encoder.EncodeResponse(val)
}
