package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	productcontrollerfactory "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/input/web/controller/product/factory"
)

type Routes struct {
	Router *gin.Engine
}

func (r *Routes) LoadRoutes() (err error) {
	rest := r.Router.Group("/hex-go")

	addFindProductController(rest)
	return
}

func addFindProductController(rest *gin.RouterGroup) {
	findProductController, err := productcontrollerfactory.GetFindProductController()
	if err != nil {
		log.Fatalf(err.Error())
	}
	rest.GET("/v1/products/:id", findProductController.ById)
}
