package product_test

import (
	"log"
	"testing"

	productrepository "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/output/persistence/product"
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
	db_test "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/test/config/db"
	seeds_test "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/test/db/seeds"
	"github.com/stretchr/testify/assert"
)

func Test_Find_A_Product_By_Id(t *testing.T) {
	ids, err := (&seeds_test.Seeds{DbClient: *db_test.Db}).Product(1)
	if err != nil {
		log.Panic(err)
	}

	findProductRepository := productrepository.NewFindProductRepository(db_test.GetProductCollection())
	productDomain := &productdomain.Product{ID: ids[0]}
	findErr := findProductRepository.ById(productDomain)

	assert.Nil(t, findErr)
	assert.NotNil(t, productDomain.Name)

	assert.EqualValues(t, ids[0], productDomain.ID)
}

func Test_Return_Error_When_Not_Find_A_Product_By_Id(t *testing.T) {
	_, err := (&seeds_test.Seeds{DbClient: *db_test.Db}).Product(1)
	if err != nil {
		log.Panic(err)
	}

	missingId := "6385e614198b9a59e233fe1f"
	findProductRepository := productrepository.NewFindProductRepository(db_test.GetProductCollection())
	productDomain := &productdomain.Product{ID: missingId}
	findErr := findProductRepository.ById(productDomain)

	assert.NotNil(t, findErr)
	assert.EqualValues(t, "no document found for id: "+missingId, findErr.Error())
}
