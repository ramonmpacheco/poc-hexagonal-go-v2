package productrepository

import (
	"context"
	"fmt"
	"log"

	"github.com/jinzhu/copier"
	"github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/output/persistence/product/model"
	productdomain "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/domain/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type findProductRepository struct {
	mongoClient *mongo.Client
}

func NewFindProductRepository(mongoClient *mongo.Client) productdomain.FindProductDataOutputPort {
	return &findProductRepository{
		mongoClient: mongoClient,
	}
}

func (fpr *findProductRepository) ById(product *productdomain.Product) error {
	objId, _ := primitive.ObjectIDFromHex(product.ID)
	filter := bson.D{{Key: "_id", Value: objId}}

	var producModel model.ProductModel

	collection := fpr.mongoClient.Database("poc_hexagonal_db").Collection("products")

	result := collection.FindOne(context.Background(), filter).Decode(&producModel)
	if result == mongo.ErrNoDocuments {
		log.Default().Println(result.Error())
		return fmt.Errorf("no document found for id: %s", product.ID)
	}
	if result != nil {
		log.Default().Println(result.Error())
		return fmt.Errorf("error trying to find document with id: %s", result.Error())
	}

	copier.Copy(&product, &producModel)

	return nil
}
