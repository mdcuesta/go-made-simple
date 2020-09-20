package services

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("CustomerService", func() {
	var target *CustomerServiceImplementation

	BeforeEach(func () {
		target = &CustomerServiceImplementation{
			customerRepository: nil,
		}
	})

	AfterEach(func () {

	})
})