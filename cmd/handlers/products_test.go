package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/GabrielMaSosa/test_funcional/internal/domain"
	product "github.com/GabrielMaSosa/test_funcional/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	os.Setenv("TOKEN", "123456")
	var slice = []domain.Product{
		{
			ID:           1,
			Name:         "Oil - Margarine",
			Quantity:     439,
			Code_value:   "S82254D",
			Is_published: true,
			Expiration:   "15/12/2021",
			Price:        71.42,
		},
		{
			ID:           2,
			Name:         "Pineapple - Canned, Rings",
			Quantity:     345,
			Code_value:   "M4637",
			Is_published: true,
			Expiration:   "09/08/2021",
			Price:        352.79,
		},
	}

	repo := product.NewRepositoryTesting(slice)
	servi := product.NewServiceProduct(&repo)
	hdler := NewHandlerProduct(servi)
	//inicio server
	server := gin.New()
	productsrout := server.Group("/products")

	Rutas(productsrout, hdler)
	return server
}

func TestFuncionalHandlerProductGetByID(t *testing.T) {

	t.Run("Happy path get by id", func(t *testing.T) {
		//arrange
		//levanto servrer
		expectedStatusCode := 200
		expectBodty := `{"id":1,"name":"Oil - Margarine","quantity":439,"code_value":"S82254D","is_published":true,"expiration":"15/12/2021","price": 71.42}`

		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		//--> input
		request := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		request.Header.Add("token", "123456")
		response := httptest.NewRecorder()
		r := createServer()
		//act
		r.ServeHTTP(response, request)

		//assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, string(expectBodty), response.Body.String(), "error")
		assert.Equal(t, expectedHeaders, response.Header())
		//assert

	})
	t.Run("One error id negative get by id", func(t *testing.T) {
		//arrange
		//levanto servrer
		expectedStatusCode := http.StatusBadRequest
		expectBodty := `{"mesage": "Bad request2"}`

		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		//--> input
		request := httptest.NewRequest(http.MethodGet, "/products/-1", nil)
		request.Header.Add("token", "123456")
		response := httptest.NewRecorder()
		r := createServer()
		//act
		r.ServeHTTP(response, request)

		//assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, string(expectBodty), response.Body.String(), "error")
		assert.Equal(t, expectedHeaders, response.Header())
		//assert

	})

}
