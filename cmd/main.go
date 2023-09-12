package main

import (
	"fmt"
	"os"

	"github.com/GabrielMaSosa/test_funcional/cmd/handlers"
	product "github.com/GabrielMaSosa/test_funcional/internal/products"
	"github.com/GabrielMaSosa/test_funcional/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("TOKEN", "123456")
	path := os.Getenv("MYPATH")
	// inyectamos las dependencias
	fmt.Println("mivar", path)

	slice, err := pkg.InitilizeBD("../products.json")
	if err != nil {
		panic(err)
	}

	repo := product.NewRepository(slice)
	servi := product.NewServiceProduct(&repo)
	hdler := handlers.NewHandlerProduct(servi)
	//fin de las inyecciones

	//inicio server
	server := gin.Default()
	productsrout := server.Group("/products")
	handlers.Rutas(productsrout, hdler)

	server.Run(":8080")

}
