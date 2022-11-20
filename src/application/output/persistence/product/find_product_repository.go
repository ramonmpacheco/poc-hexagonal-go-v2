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
	mongoCollection *mongo.Collection
}

func NewFindProductRepository(mongoClient *mongo.Collection) productdomain.FindProductDataOutputPort {
	return &findProductRepository{
		mongoCollection: mongoClient,
	}
}

func (fpr *findProductRepository) ById(product *productdomain.Product) error {
	id, _ := primitive.ObjectIDFromHex(product.ID)
	filter := bson.D{{Key: "_id", Value: id}}

	var producModel model.ProductModel

	result := fpr.mongoCollection.FindOne(context.Background(), filter).Decode(&producModel)
	if result != nil {
		log.Default().Println(result.Error())
		if result == mongo.ErrNoDocuments {
			log.Default().Println(result.Error())
			return fmt.Errorf("no document found for id: %s", product.ID)
		}
		return fmt.Errorf("error trying to find document with id: %s", product.ID)
	}

	copier.Copy(&product, &producModel)

	return nil
}
