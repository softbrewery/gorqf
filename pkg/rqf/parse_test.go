package rqf_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/softbrewery/gojoi/pkg/joi"

	. "github.com/softbrewery/gorqf/pkg/rqf"
)

var _ = Describe("Parse", func() {

	Describe("NewParser", func() {

		parser := NewParser()

		It("Should not be nil", func() {
			Expect(parser).ToNot(BeNil())
		})
	})

	Describe("Parse", func() {

		It("Should parse a filter as url", func() {
			filter := `http://app.myapp.com/api/v1/books?filter={"fields":["id","title","desc"]}`
			parser, err := NewParser().Parse(filter)

			Expect(parser).NotTo(BeNil())
			Expect(err).To(BeNil())
		})

		It("Should parse a filter as escaped url", func() {
			filter := `http://app.myapp.com/api/v1/books?filter=%7B%22fields%22%3A%5B%22id%22%2C%22title%22%2C%22desc%22%5D%7D`
			parser, err := NewParser().Parse(filter)

			Expect(parser).NotTo(BeNil())
			Expect(err).To(BeNil())
		})

		It("Should fail if a filter has invalid escaped url", func() {
			filter := `http://app.myapp.com/api/v1/books?filter=%f%22fields%22%3A%5B%22id%22%2C%22title%22%2C%22desc%22%5D%7D`
			parser, err := NewParser().Parse(filter)

			Expect(parser).To(BeNil())
			Expect(err).NotTo(BeNil())
		})

		It("Should parse a filter as url query", func() {
			filter := `filter={"fields":["id","title","desc"]}`
			parser, err := NewParser().Parse(filter)

			Expect(parser).NotTo(BeNil())
			Expect(err).To(BeNil())
		})

		It("Should parse a filter as url query (with question mark)", func() {
			filter := `?filter={"fields":["id","title","desc"]}`
			parser, err := NewParser().Parse(filter)

			Expect(parser).NotTo(BeNil())
			Expect(err).To(BeNil())
		})

		It("Should parse a filter as url query parameter", func() {
			filter := `filter={"fields":["id","title","desc"]}`
			parser, err := NewParser().Parse(filter)

			Expect(parser).NotTo(BeNil())
			Expect(err).To(BeNil())
		})

		It("Should parse a filter as json string", func() {
			filter := `{"fields":["id","title","desc"]}`
			parser, err := NewParser().Parse(filter)

			Expect(parser).NotTo(BeNil())
			Expect(err).To(BeNil())
		})

		It("Should fail to parse a filter as url if filter not present", func() {
			filter := `http://app.myapp.com/api/v1/books?data={"fields":["id","title","desc"]}`
			parser, err := NewParser().Parse(filter)

			Expect(parser).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("FieldSchema", func() {

		It("Should succeed if fieldschema is not set", func() {
			parser := NewParser()

			jsonFilter := `{"fields":["id","title"]}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).To(BeNil())
			Expect(filter.Fields).To(Equal([]string{"id", "title"}))
		})

		It("Should succeed if fields are matching", func() {
			schema := joi.String().Allow("id", "title")

			parser := NewParser()
			parser.FieldSchema(schema)

			jsonFilter := `{"fields":["id","title"]}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).To(BeNil())
			Expect(filter.Fields).To(Equal([]string{"id", "title"}))
		})

		It("Should fail if fields are not matching", func() {
			schema := joi.String().Allow("id", "title")

			parser := NewParser()
			parser.FieldSchema(schema)

			jsonFilter := `{"fields":["id","desc"]}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).NotTo(BeNil())
			Expect(filter).To(BeNil())
		})
	})

	Describe("OrderSchema", func() {

		It("Should succeed if orderschema is not set", func() {
			parser := NewParser()

			jsonFilter := `{"order":["id","title"]}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).To(BeNil())
			Expect(filter.Order).To(Equal([]string{"id", "title"}))
		})

		It("Should succeed if order are matching", func() {
			schema := joi.String().Allow("id", "title")

			parser := NewParser()
			parser.OrderSchema(schema)

			jsonFilter := `{"order":["id","title"]}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).To(BeNil())
			Expect(filter.Order).To(Equal([]string{"id", "title"}))
		})

		It("Should fail if order are not matching", func() {
			schema := joi.String().Allow("id", "title")

			parser := NewParser()
			parser.OrderSchema(schema)

			jsonFilter := `{"order":["id","desc"]}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).NotTo(BeNil())
			Expect(filter).To(BeNil())
		})
	})

	Describe("LimitSchema", func() {

		It("Should succeed if limit is not set", func() {
			parser := NewParser()

			jsonFilter := `{"limit": 100}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).To(BeNil())
			Expect(filter.Limit).To(Equal(100))
		})

		It("Should succeed if limit is matching", func() {
			schema := joi.Int().Max(100)

			parser := NewParser()
			parser.LimitSchema(schema)

			jsonFilter := `{"limit": 100}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).To(BeNil())
			Expect(filter.Limit).To(Equal(100))
		})

		It("Should fail if limit is not matching", func() {
			schema := joi.Int().Max(100)

			parser := NewParser()
			parser.LimitSchema(schema)

			jsonFilter := `{"limit": 101}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).ToNot(BeNil())
			Expect(filter).To(BeNil())
		})
	})

	Describe("OffsetSchema", func() {

		It("Should succeed if offset is not set", func() {
			parser := NewParser()

			jsonFilter := `{"offset": 3}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).To(BeNil())
			Expect(filter.Offset).To(Equal(3))
		})

		It("Should succeed if offset is matching", func() {
			schema := joi.Int().Max(100)

			parser := NewParser()
			parser.OffsetSchema(schema)

			jsonFilter := `{"offset": 3}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).To(BeNil())
			Expect(filter.Offset).To(Equal(3))
		})

		It("Should fail if offset is not matching", func() {
			schema := joi.Int().Max(10)

			parser := NewParser()
			parser.OffsetSchema(schema)

			jsonFilter := `{"offset": 11}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).ToNot(BeNil())
			Expect(filter).To(BeNil())
		})
	})

	Describe("WhereSchema", func() {

		It("Should succeed if where is not set", func() {
			parser := NewParser()

			jsonFilter := `{"where":{"isbn": "A_ISBN"}}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).To(BeNil())
			Expect(filter.Where).To(Equal(map[string]interface{}{"isbn": "A_ISBN"}))
		})

		It("Should succeed if where is matching", func() {
			schema := joi.Any()

			parser := NewParser()
			parser.WhereSchema(schema)

			jsonFilter := `{"where":{"isbn": "A_ISBN"}}`
			filter, err := parser.Parse(jsonFilter)

			Expect(err).To(BeNil())
			Expect(filter.Where).To(Equal(map[string]interface{}{"isbn": "A_ISBN"}))
		})

		It("Should fail if where is not matching", func() {
			schema := joi.Any().Forbidden()

			parser := NewParser()
			parser.WhereSchema(schema)

			jsonFilter := `{"where":{"isbn": "A_ISBN"}}`
			filter, err := parser.Parse(jsonFilter)

			fmt.Println(filter)

			Expect(err).NotTo(BeNil())
			Expect(filter).To(BeNil())
		})
	})
})
