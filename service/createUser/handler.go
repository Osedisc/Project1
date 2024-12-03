package createUser

import "github.com/gin-gonic/gin"

type Servicer interface {
	Execute(c *gin.Context, req Request) (*Response, error)
}

type handler struct {
	service Servicer
}

func NewHandler(servicer Servicer) *handler {
	return &handler{service: servicer}
}

func (h handler) CreateUser(c *gin.Context) {
	req := Request{}
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	res, err := h.service.Execute(c, req)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	c.JSON(200, res)
}
