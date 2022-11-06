package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productusecase "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/port/input/product"
)

type findProductController struct {
	findProductUseCase productusecase.FindProductUseCase
}

func NewFindProductController(findProductUseCase productusecase.FindProductUseCase) *findProductController {
	return &findProductController{
		findProductUseCase: findProductUseCase,
	}
}

func (fpc *findProductController) ById(c *gin.Context) {
	id := c.Param("id")
	product, _ := fpc.findProductUseCase.ById(id)
	c.JSON(http.StatusOK, product)
}
