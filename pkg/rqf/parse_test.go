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
				filter, err := Parse(rawJson, nil)
				Expect(err).To(BeNil())
				Expect(filter).NotTo(BeNil())
			}

			{
				filter, err := Parse(rawJson, joi.AnySchema().Allow(""))
				Expect(err).To(BeNil())
				Expect(filter).NotTo(BeNil())
			}
		})
	})
})
