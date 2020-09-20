package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFirstApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FirstApi Suite")
}
