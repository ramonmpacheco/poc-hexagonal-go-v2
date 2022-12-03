package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	productcontrollerfactory "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/input/web/controller/factory"
)

type Routes struct {
	Router *gin.Engine
}

func (r *Routes) LoadRoutes() (err error) {
	r.Router.Use(CORSMiddleware())
	rest := r.Router.Group("/hex-go")
	loadProductControllers(rest)
	return
}

func loadProductControllers(rest *gin.RouterGroup) {
	err := productcontrollerfactory.AddProductControllers(rest)
	if err != nil {
		log.Fatalf(err.Error())
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
