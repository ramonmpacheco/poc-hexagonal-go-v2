package productusecaseimpl

import (
	"testing"

	productusecaseimpl "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/port/input/product/impl"
	mocks_product "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/test/mocks/product"
	"github.com/stretchr/testify/assert"
)

func Test_Return_Error_When_Id_Is_Invalid(t *testing.T) {
	mockFindProductData := mocks_product.NewMockFindProductDataOutputPort(nil)
	productUseCaseImpl := productusecaseimpl.NewFindProductUseCaseImpl(mockFindProductData)

	product, err := productUseCaseImpl.ById("abc")
	assert.Nil(t, product)
	assert.EqualValues(t, "invalid id", err.Error())
}
