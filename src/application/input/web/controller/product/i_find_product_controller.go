package productcontroller

import "github.com/gin-gonic/gin"

type IFindProductController interface {
	ById(c *gin.Context)
}
