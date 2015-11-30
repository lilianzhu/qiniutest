package acc_test

import (
	. "qiniutest.com/lib/def"
)

var _ = Describe("Acc", func() {
	It("获取用户acctoken成功，返回200", func() {
		resp, acctoken := GetAccessToken(Env("username"), Env("password"))
		Expect(resp.Status()).To(Equal(200))
		Expect(acctoken).NotTo(BeNil())
	})
})
