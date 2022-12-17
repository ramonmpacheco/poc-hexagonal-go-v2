package productcontroller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/input/web/controller/responses"
	mocks_product "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/test/mocks/product"
	producttemplates "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/test/templates/product_templates"
	"github.com/stretchr/testify/assert"
)

func Test_Find_A_Product_By_Id(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDomain := producttemplates.GetProductDomain("")
	mockUseCase := mocks_product.NewMockFindProductUseCase(mockCtrl)
	mockUseCase.EXPECT().ById(gomock.Any()).Return(mockDomain, nil)
	controller := NewFindProductController(mockUseCase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/test/v1/products/:id", controller.ById)

	req, err := http.NewRequest(http.MethodGet, "/test/v1/products/"+mockDomain.ID, nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response responses.ProductResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.EqualValues(t, mockDomain.ID, response.ID)
	assert.EqualValues(t, mockDomain.Name, response.Name)
}
