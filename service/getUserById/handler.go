package getUserById

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

func (h handler) GetUserById(c *gin.Context) {
	req := Request{}
	req.Id = c.Param("id")
	res, err := h.service.Execute(c, req)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	c.JSON(200, res)
}
