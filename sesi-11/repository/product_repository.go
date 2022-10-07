package repository

import "sesi-11/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
