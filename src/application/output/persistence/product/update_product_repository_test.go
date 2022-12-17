package productrepository

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func Test_Update_Product_By_Id(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// mt.Run("no document to update", func(mt *mtest.T) {
	// 	var productCollection = mt.Coll

	// 	mt.AddMockResponses(bson.D{
	// 		{Key: "ok", Value: 1},
	// 		{Key: "acknowledged", Value: true},
	// 		{Key: "n", Value: 0},
	// 	})

	// 	repo := NewFindProductRepository(productCollection)
	// 	f := productdomain.Product{ID: "637911e88e5647cccad9dc2c", Name: "Test"}

	// })

	// mt.Run("failed", func(mt *mtest.T) {
	// 	var productCollection = mt.Coll
	// 	mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
	// 		Index:   1,
	// 		Code:    11000,
	// 		Message: "product not found",
	// 	}))
	// })

}
