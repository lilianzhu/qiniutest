package acc_test

import (
	. "qiniutest.com/lib/def"
)

var _ = Describe("Acc", func() {
	It("获取用户acctoken成功，返回200", func() {
		resp := GetAccToken(ENV("username"), ENV("password"))
		result := AccTokenResult
		Expect(resp.Status()).To(Equal(200))
		err := resp.Unmarshal(&result)
		Expect(err).To(BeNil())
		println(result.AccToken)
		Expect(result.AccToken).NotTo(BeNil())
		Expect(result.RefToken).NotTo(BeNil())
	})
})
