package productrepository

import (
	"fmt"
	"testing"

	productrepository "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/output/persistence/product"
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func Test_Find_Product_By_Id(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("no document found", func(mt *mtest.T) {
		var productCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "poc_hexagonal_db.products", mtest.FirstBatch))

		repo := productrepository.NewFindProductRepository(productCollection)
		product := productdomain.Product{ID: "637911e88e5647cccad9dc2c"}
		result := repo.ById(&product)

		assert.EqualValues(t, "no document found for id: 637911e88e5647cccad9dc2c", result.Error())
	})

	mt.Run("failure", func(mt *mtest.T) {
		var productCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		repo := productrepository.NewFindProductRepository(productCollection)
		product := productdomain.Product{ID: "637911e88e5647cccad9dc2c"}
		result := repo.ById(&product)

		assert.EqualValues(t, fmt.Sprintf("error trying to find document with id: %s", product.ID), result.Error())
	})

	mt.Run("success", func(mt *mtest.T) {
		var productCollection = mt.Coll

		expectedProduct := productdomain.Product{ID: "637911e88e5647cccad9dc2c", Name: "Test"}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "poc_hexagonal_db.products", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedProduct.ID},
			{Key: "name", Value: expectedProduct.Name},
		}))

		repo := productrepository.NewFindProductRepository(productCollection)
		product := productdomain.Product{ID: "637911e88e5647cccad9dc2c"}
		result := repo.ById(&product)
		assert.Nil(t, result)
	})
}
