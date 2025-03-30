package main

import (
	"fmt"
	"log"

	"dynamicWeb/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFile("/orders", "./static/orders.html")
	router.StaticFile("/products", "./static/products.html")
	router.StaticFile("/", "./static/index.html") // Home page
	router.StaticFile("/js/main.js", "./static/js/main.js")

	router.Use(cors.Default())

	routes.SetupConfigRoutes(router)
	routes.SetupSpecificRoutes(router)

	fmt.Println("Server running at http://127.0.0.1:8080")
	log.Fatal(router.Run(":8080"))
}
