package learning1_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLearning1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Learning1 Suite")
}
