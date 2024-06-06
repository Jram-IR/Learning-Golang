package helper

import "github.com/gin-gonic/gin"

func HandleError(c *gin.Context, err error, code int, emsg string) {
	c.String(code, emsg)
	return
}
