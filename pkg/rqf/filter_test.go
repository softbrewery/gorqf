package rqf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/softbrewery/gorqf/pkg/rqf"
)

var _ = Describe("Filter", func() {

	Describe("NewFilter", func() {

		filter := NewFilter()

		It("Should not be nil", func() {
			Expect(filter).ToNot(BeNil())
		})

		It("Fields should be empty", func() {
			Expect(filter.Fields).To(BeEmpty())
		})

		It("Order should be empty", func() {
			Expect(filter.Order).To(BeEmpty())
		})

		It("Limit should be 0", func() {
			Expect(filter.Limit).To(Equal(0))
		})

		It("Offset should be 0", func() {
			Expect(filter.Offset).To(Equal(0))
		})

		It("Where should be nil", func() {
			Expect(filter.Where).To(BeNil())
		})
	})

	Describe("IsEmpty", func() {

		It("Should be empty when data is not set", func() {
			filter := NewFilter()
			Expect(filter.IsEmpty()).To(BeTrue())
		})

		It("Should not be empty when fields is set", func() {
			filter := Filter{Fields: []string{"name"}}
			Expect(filter.IsEmpty()).To(BeFalse())
		})

		It("Should not be empty when order is set", func() {
			filter := Filter{Order: []string{"name ASC"}}
			Expect(filter.IsEmpty()).To(BeFalse())
		})

		It("Should not be empty when limit is set", func() {
			filter := Filter{Limit: 100}
			Expect(filter.IsEmpty()).To(BeFalse())
		})

		It("Should not be empty when offset is set", func() {
			filter := Filter{Offset: 10}
			Expect(filter.IsEmpty()).To(BeFalse())
		})

		It("Should not be empty when where is set", func() {
			filter := Filter{Where: map[string]string{}}
			Expect(filter.IsEmpty()).To(BeFalse())
		})
	})
})
