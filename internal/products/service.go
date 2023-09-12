package product

import "github.com/GabrielMaSosa/otherverbs/internal/domain"

type ProductService interface {
	ServiceGetAll() (dt []domain.Product, err error)
	UpdateItem(id int, data domain.Product) (dt *domain.Product, err error)
	Delete(id int) (dt *domain.Product, err error)
	UpdatePartial(id int, data map[string]interface{}) (dt *domain.Product, err error)
	ServiGetById(id int) (dt *domain.Product, err error)
	ServiGetPriceMayor(price float64) (dt []domain.Product, err error)
}
