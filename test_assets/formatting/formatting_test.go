package passing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FORMATTING", func() {
	Describe("this (level) has parenthesis", func() {
		It("test 1 passes", func() {
			Expect(true).To(Equal(true))
		})

		It("test 2 passes", func() {
			Expect(true).To(Equal(true))
		})
	})

	Describe("this /level/ has slashes", func() {
		It("test 1 passes", func() {
			Expect(true).To(Equal(true))
		})

		It("test 2 passes", func() {
			Expect(true).To(Equal(true))
		})
	})
})
