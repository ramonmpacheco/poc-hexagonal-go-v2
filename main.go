package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ramonmpacheco/poc-hexagonal-go-v2/src/application/input/web/routes"
)

func main() {
	router := gin.Default()
	r := routes.Routes{
		Router: router,
	}

	if err := r.LoadRoutes(); err != nil {
		log.Fatalf("Error loading routes %s", err.Error())
	}

	router.Run()
}
