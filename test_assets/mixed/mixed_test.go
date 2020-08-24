package mixed_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("level 1", func() {
	Describe("A", func() {
		It("test 1 fails", func() {
			Expect(true).To(Equal(false))
		})

		It("test 2 passes", func() {
			Expect(true).To(Equal(true))
		})
	})

	Describe("B", func() {
		It("test 1 fails", func() {
			Expect(true).To(Equal(false))
		})

		It("test 2 passes", func() {
			Expect(true).To(Equal(true))
		})
	})
})
