package acc_test

import (
	"qiniutest.com/biz/acc"
	. "qiniutest.com/configs"
	// "qiniutest.com/lib/util"
)

var _ = Describe("Acc", func() {
	It("获取用户acctoken成功，返回200", func() {
		resp, acctoken := acc.GetToken(ENV("username"), ENV("password"))
		Expect(resp.Status()).To(Equal(200))
		Expect(acctoken).NotTo(BeNil())
	})
})
