package seeds_test

import (
	"context"
	"log"
	"strconv"

	"github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/output/persistence/product/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Seeds struct {
	DbClient mongo.Client
}

func (s *Seeds) Product(qtd int) ([]string, error) {
	log.Default().Println("seeding product...")
	colllection := s.DbClient.Database("poc_hexagonal_db").Collection("products")
	var ids []string
	for i := 0; i < qtd; i++ {
		if result, err := colllection.InsertOne(
			context.Background(),
			model.ProductModel{Name: "Product_" + strconv.Itoa(i)},
		); err != nil {
			return nil, err
		} else {
			id := result.InsertedID.(primitive.ObjectID).Hex()
			log.Default().Printf("inserted document with _id: %v\n", id)
			ids = append(ids, id)
		}
	}
	log.Default().Println("seeding product, done")
	return ids, nil
}
