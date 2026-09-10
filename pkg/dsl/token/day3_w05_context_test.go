package token

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Context Keyword Token Specs", func() {
	Context("Token initialization for rule context parameters", func() {
		It("should correctly identify CONTEXT keyword token", func() {
			tok := New(IDENT, "context", PositionInfo{LinePosition: 5, ColumnPosition: 1})
			Expect(tok.Type).To(Equal(IDENT))
			Expect(tok.Literal).To(Equal("context"))
			Expect(tok.PositionInfo.LinePosition).To(Equal(5))
		})
	})
})
