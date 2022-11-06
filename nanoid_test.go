package nanoid

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NanoID", func() {

	Describe("New", func() {
		Context("When called", func() {
			It("Should generate a valid value without panic", func() {

				// test: execution
				actual := New()

				// test: validation
				Expect(actual).ShouldNot(BeEmpty())

			})
		})
	})

	Describe("NewID", func() {
		Context("When called 500,000 times", func() {
			It("Should generate values that do not collide", func() {

				// test: setup
				count := 500_000
				values := make(map[string]bool)

				// test: iterate the desired count
				for i := 0; i < count; i++ {

					// test: execution
					actual, err := NewID()

					// test: validation(s)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(values[actual.String()]).Should(BeFalse())

				}

			})
		})
	})

	Describe("NewString", func() {
		Context("When called", func() {
			It("Should generate and return a string value", func() {

				// test: execution
				foo := ""
				actual := NewString()

				// test: validation
				Expect(actual).ShouldNot(BeEmpty())
				Expect(actual).Should(BeAssignableToTypeOf(foo))

			})
		})
	})

})
