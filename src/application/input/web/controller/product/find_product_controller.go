package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/input/web/controller/converters"
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
	product, err := fpc.findProductUseCase.ById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response := converters.ProductDomainToResponse(product)
	c.JSON(http.StatusOK, response)
}
