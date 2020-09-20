package services

import (
	. "first-api/data"
	. "first-api/data/models"
	. "first-api/services/dtos"
)

type CustomerService interface {
	Get(id uint) (*CustomerDto, error)
	Create(customerDto *CustomerDto) (uint, error)
	Update(customerDto *CustomerDto) (bool, error)
}

type CustomerServiceImplementation struct {
	customerRepository CustomerRepository
}

func (service *CustomerServiceImplementation) Get(id uint) (*CustomerDto, error) {
	var (
		customer *Customer
		err error
	)
	customer, err = service.customerRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return &CustomerDto{
		ID: customer.ID,
		FirstName: customer.FirstName,
		MiddleName: customer.MiddleName,
		LastName: customer.LastName,
		Email: customer.Email,
	}, nil
}

func (service *CustomerServiceImplementation) Create(customerDto *CustomerDto) (uint, error) {
	customer := &Customer{
		FirstName: customerDto.FirstName,
		MiddleName: customerDto.MiddleName,
		LastName: customerDto.LastName,
		Email: customerDto.Email,
	}

	var (
		id uint
		err error
	)

	id, err = service.customerRepository.Insert(customer)

	return id, err
}

func (service *CustomerServiceImplementation) Update(customerDto *CustomerDto) (bool, error) {
	var (
		err error
		success bool
		customer *Customer
	)
	customer, err = service.customerRepository.Get(customerDto.ID)

	if err != nil {
		return false, err
	}

	customer.FirstName = customerDto.FirstName
	customer.MiddleName = customerDto.MiddleName
	customer.LastName = customerDto.LastName
	customer.Email = customerDto.Email

	success, err = service.customerRepository.Update(customer)

	return success, err
}