package learning1_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Parallelsecond", func() {

	BeforeEach(func() {
	})

	It("parallel test1", func() {  
		time.Sleep(2*time.Second)
      })

	It("parallel test2", func() {  
		time.Sleep(1*time.Second)
      })
	AfterEach(func() {
	})
})
