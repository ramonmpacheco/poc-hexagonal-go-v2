package productusecaseimpl

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
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

func Test_Forward_Error_When_Get_Error_From_Persistence(t *testing.T) {
	mockFindProductData := mocks_product.NewMockFindProductDataOutputPort(gomock.NewController(t))
	productUseCaseImpl := productusecaseimpl.NewFindProductUseCaseImpl(mockFindProductData)
	errorReturn := errors.New("no document found for id: 6385e614198b9a59e233fe1f")
	mockFindProductData.EXPECT().ById(gomock.Any()).Return(errorReturn)

	product, err := productUseCaseImpl.ById("6385e614198b9a59e233fe1f")
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.EqualValues(t, errorReturn.Error(), err.Error())
}
