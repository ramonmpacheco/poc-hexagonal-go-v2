package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/input/web/routes"
	mongoclient "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/config/db"
)

func main() {
	client, ctx, cancel, err := mongoclient.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer mongoclient.Close(client, ctx, cancel)
	mongoclient.Ping(client, ctx)

	router := gin.Default()
	r := routes.Routes{
		Router: router,
	}

	if err := r.LoadRoutes(); err != nil {
		log.Fatalf("Error loading routes %s", err.Error())
	}

	router.Run()
}
