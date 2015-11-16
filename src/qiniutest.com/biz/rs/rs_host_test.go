package rs

import (
	//"github.com/jmcvetta/napping"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get /host/<bucket>", func() {

	Describe("正向用例", func() {

		Context("对于z0的空间", func() {
			It("应该能得到华北的UP和IO的Host地址", func() {

				Expect(true).To(Equal(true))

			})
		})

		Context("对于z1的空间", func() {
			It("应该能得到华东UP和IO的Host的地址", func() {

				Expect(true).To(Equal(true))

			})

		})

	})

})