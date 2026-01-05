package product

import (
	"validation-api/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{Database: database}
}

func (repository *ProductRepository) GetAll() ([]Product, error) {
	var products []Product
	res := repository.Database.Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products, nil
}

func (repository *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	res := repository.Database.DB.First(&product, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &product, nil
}
func (repository *ProductRepository) Create(product *Product) (*Product, error) {
	res := repository.Database.DB.Create(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (repository *ProductRepository) Update(product *Product) (*Product, error) {
	res := repository.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (repository *ProductRepository) Delete(id uint) error {
	res := repository.Database.DB.Delete(&Product{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
