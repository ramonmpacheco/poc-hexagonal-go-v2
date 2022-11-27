package db_test

import (
	"context"
	"fmt"
	"log"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Client
var pool *dockertest.Pool
var resource *dockertest.Resource

const MONGO_INITDB_ROOT_USERNAME = "root"
const MONGO_INITDB_ROOT_PASSWORD = "password"

func Connect() {
	var err error
	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	log.Default().Println("starting conection with db")
	environmentVariables := []string{
		"MONGO_INITDB_ROOT_USERNAME=" + MONGO_INITDB_ROOT_USERNAME,
		"MONGO_INITDB_ROOT_PASSWORD=" + MONGO_INITDB_ROOT_PASSWORD,
	}
	resource = starContainerPool(pool, environmentVariables)
	conectToContainerPool(pool, resource)
}

func starContainerPool(pool *dockertest.Pool, env []string) *dockertest.Resource {
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "5.0",
		Env:        env,
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true
		// so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	return resource
}

func conectToContainerPool(pool *dockertest.Pool, resource *dockertest.Resource) {
	err := pool.Retry(func() error {
		var err error
		Db, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(getMongoUri(resource)),
		)
		if err != nil {
			return err
		}
		return Db.Ping(context.TODO(), nil)
	})
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
}

func getMongoUri(resource *dockertest.Resource) string {
	return fmt.Sprintf(
		"mongodb://%s:%s@localhost:%s",
		MONGO_INITDB_ROOT_USERNAME,
		MONGO_INITDB_ROOT_PASSWORD, resource.GetPort("27017/tcp"))
}

func teardown(pool *dockertest.Pool, resource *dockertest.Resource) {
	log.Default().Println("tearing down: killing and removing the container")
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func Close() {
	teardown(pool, resource)
	log.Default().Println("closing conection with db")
	if err := Db.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func GetProductCollection() *mongo.Collection {
	return Db.Database("poc_hexagonal_db").Collection("products")
}
