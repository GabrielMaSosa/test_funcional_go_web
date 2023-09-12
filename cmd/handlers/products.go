package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/GabrielMaSosa/test_funcional/internal/domain"
	product "github.com/GabrielMaSosa/test_funcional/internal/products"
	"github.com/gin-gonic/gin"
)

// agrego la interfaz para la composicion
type HandlerProduct struct {
	service product.ProductService
}

// Constructor del Handler
func NewHandlerProduct(sv *product.ProductServiceImpl) *HandlerProduct {

	return &HandlerProduct{service: sv}
}

// Definicion de las rutas
func Rutas(g *gin.RouterGroup, h *HandlerProduct) {
	g.GET("", h.GetAll())
	g.PUT("/:id", h.Update())
	g.DELETE("/:id", h.Delete())
	g.PATCH("/:id", h.PartialSave())
	g.GET("/:id", h.GetProductByIdPatch())
	g.GET("/search", h.SearchProduct())
}

var (
	ErrNotFound = errors.New("Not found")
)

func (h *HandlerProduct) PartialSave() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if mitoken := ctx.GetHeader("token"); mitoken != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorize"})
			return
		}
		id := ctx.Param("id")
		idn, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": "Bad request1"})
			return
		}
		if idn <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": "Bad request2"})
			return
		}

		dta := map[string]interface{}{}
		if err := ctx.ShouldBindJSON(&dta); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		if err2 := ValidateRequest(dta); err2 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": err2.Error()})
			fmt.Println(err2)
			return
		}

		dta1, err5 := h.service.UpdatePartial(idn, dta)
		if err5 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Error internal"})
			return
		} else {
			ctx.JSON(http.StatusOK, dta1)
		}

	}
}

func (h *HandlerProduct) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if mitoken := ctx.GetHeader("token"); mitoken != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorize"})
			return
		}
		dta, err := h.service.ServiceGetAll()
		if err != nil {
			ctx.String(http.StatusNotFound, ErrNotFound.Error())
			return
		}
		ctx.JSON(http.StatusOK, dta)
		return
	}
}

func (h *HandlerProduct) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if mitoken := ctx.GetHeader("token"); mitoken != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorize"})
			return
		}
		var datain domain.Product

		id := ctx.Param("id")
		idn, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": "Bad request1"})
			return
		}
		if idn <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": "Bad request2"})
			return
		}
		if err2 := ctx.ShouldBindJSON(&datain); err2 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": "Bad request3"})
			return

		}
		if err3 := ValidateData(datain); err3 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": "Bad request4"})
			return
		}
		val2, err5 := h.service.UpdateItem(idn, datain)
		if err5 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "internal error"})
			return
		}
		ctx.JSON(http.StatusOK, val2)
		return
	}
}

func (h *HandlerProduct) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if mitoken := ctx.GetHeader("token"); mitoken != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorize"})
			return
		}
		id := ctx.Param("id")
		idn, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": "Bad request1"})
			return
		}
		if idn <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": "Bad request2"})
			return
		}

		_, err5 := h.service.Delete(idn)
		if err5 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "internal error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"msg": "Succesful delete"})
		return
	}

}

// GetProduct traer un producto por id
func (h *HandlerProduct) GetProductByIdPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if mitoken := ctx.GetHeader("token"); mitoken != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorize"})
			return
		}
		id := ctx.Param("id")
		idn, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": "Bad request1"})
			return
		}
		if idn <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": "Bad request2"})
			return
		}

		valore, err := h.service.ServiGetById(idn)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
			return
		}

		ctx.JSON(200, valore)
		return
	}
}

// SearchProduct traer un producto por nombre o categoria
func (h *HandlerProduct) SearchProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if mitoken := ctx.GetHeader("token"); mitoken != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorize"})
			return
		}
		query := ctx.Query("priceGt")
		priceGt, err := strconv.ParseFloat(query, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
			return
		}
		valore, err := h.service.ServiGetPriceMayor(priceGt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		ctx.JSON(200, valore)
		return
	}
}
