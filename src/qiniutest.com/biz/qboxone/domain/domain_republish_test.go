package domain_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "qiniutest.com/biz/qboxone/domain"
)

/*
* ISSUE: https://pm.qbox.me/issues/22085
* APIDoc: https://github.com/qbox/bs-apidoc/blob/master/apidoc/v6/qboxone/domain.md#republish-%E8%BF%81%E7%A7%BB%E5%9F%9F%E5%90%8D
 */
var _ = Describe("迁移域名 [apihost]/v6/domain/republish", func() {

	Describe("正向用例", func() {

		Context("使用Admin Token", func() {

			It("同一人，Publish 一个Bucket的Domain到另一个bucket", func() {
				PublishDomain("Test", "Test", "UID")

				Expect(true).To(Equal(true))
			})

			It("把一个账号的Domain Republish到另一个人的bucket上", func() {

			})

		})

		Context("使用一般用户Token", func() {
			It("同一人，Publish 一个Bucket的Domain到另一个bucket", func() {

			})

		})

	})

	Context("反向用例", func() {

	})
})
