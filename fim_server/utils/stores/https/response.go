package https

import "github.com/gin-gonic/gin"

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type baseResponse struct{ c *gin.Context }

func Response(c *gin.Context) *baseResponse { return &baseResponse{c: c} }

func (r *baseResponse) Error(data interface{}) {
	r.c.JSON(200, response{Code: 1000, Message: "error", Data: data})
}
func (r *baseResponse) Ok(data interface{}) {
	r.c.JSON(200, response{Code: 0, Message: "ok", Data: data})
}
