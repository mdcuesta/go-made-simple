package data

import (
	."first-api/data/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Get(id uint) (*Customer, error)
	Insert(customer *Customer) (uint, error)
	Update(customer *Customer) (bool, error)
	All() ([]*Customer, error)
}

type SqlCustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository() *SqlCustomerRepository {
	return &SqlCustomerRepository{
		db: getDb(),
	}
}

func (repository *SqlCustomerRepository) Get(id uint) (*Customer, error) {
	customer := &Customer{}
	err := repository.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (repository *SqlCustomerRepository) Insert(customer *Customer) (uint, error) {
	err := repository.db.Create(customer).Error
	if err != nil {
		return 0, err
	}

	return customer.ID, nil
}

func (repository *SqlCustomerRepository) Update(customer *Customer) (bool, error) {
	err := repository.db.Save(customer).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *SqlCustomerRepository) All() ([]*Customer, error) {
	var customers []*Customer
	err := repository.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, err
}