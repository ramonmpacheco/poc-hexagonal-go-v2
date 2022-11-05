package routes

import "github.com/gin-gonic/gin"

type Routes struct {
	Router *gin.Engine
}

func (r *Routes) LoadRoutes() (err error) {
	return
}
