package data

import (
	"first-api/data/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Get(id uint) (*models.Product, error)
	Insert(product *models.Product) (uint, error)
	Update(product *models.Product) (bool, error)
	All() ([]*models.Product, error)
}

type SqlProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() *SqlProductRepository {
	return &SqlProductRepository{
		db: getDb(),
	}
}

func (repository *SqlProductRepository) Get(id uint) (*models.Product, error) {
	product := &models.Product{}
	err := repository.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (repository *SqlProductRepository) Insert(product *models.Product) (uint, error) {
	err := repository.db.Create(product).Error
	if err != nil {
		return 0, err
	}

	return product.ID, nil
}

func (repository *SqlProductRepository) Update(product *models.Product) (bool, error) {
	err := repository.db.Save(product).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *SqlProductRepository) All() ([]*models.Product, error) {
	var products []*models.Product
	err := repository.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, err
}