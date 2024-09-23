package learning1_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Serialfirst",Serial,func() {
	BeforeEach(func() {
	})

	It("serial test1", func() {  
		time.Sleep(2*time.Second)
      })

	It("serial test2", func() {  
		time.Sleep(1*time.Second)
      })
	AfterEach(func() {
	})
})
