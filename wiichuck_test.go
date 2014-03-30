package gobotI2C

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wiichuck", func() {
	var (
			someAdaptor TestAdaptor
			someDriver *Wiichuck
		)

	BeforeEach(func() {
		someDriver = NewWiichuck(someAdaptor)
	})

	PIt("Should be able to Start", func() {
		Expect(true)
	})
})
