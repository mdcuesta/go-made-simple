package data

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("CustomerRepository", func() {
	var repository *SqlCustomerRepository
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var (
			db *sql.DB
			err error
		)

		db, mock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())

		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{})
		Expect(err).Should(HaveOccurred())

		repository = &SqlCustomerRepository{db: gormDB}
		Expect(repository).Should(BeNil())
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet()
		Expect(err).ShouldNot(HaveOccurred())
	})
})
