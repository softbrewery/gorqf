package rqf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/softbrewery/gojoi/pkg/joi"

	. "github.com/softbrewery/gorqf/pkg/rqf"
)

var _ = Describe("Parse", func() {

	Describe("sssParse", func() {

		It("fffParse", func() {

			rawJson := `{
				"fields": [
					"id",
					"title",
					"isbn"
				]
			}`

			{
				parser := NewParser()
				filter, err := parser.Parse(rawJson)

				Expect(err).To(BeNil())
				Expect(filter).NotTo(BeNil())
			}

			{
				parser := NewParser()
				parser.FieldsSchema(
					joi.String().Allow("id", "title", "isbn"),
				)
				filter, err := parser.Parse(rawJson)

				Expect(err).To(BeNil())
				Expect(filter).NotTo(BeNil())
			}

			{
				parser := NewParser()
				parser.FieldsSchema(
					joi.String().Allow("id", "title"),
				)
				filter, err := parser.Parse(rawJson)

				Expect(err).ToNot(BeNil())
				Expect(filter).To(BeNil())
			}
		})
	})
})
