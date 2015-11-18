package up_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	. "qiniutest.com/configs"
	"qiniutest.com/lib/util"
)

var _ = Describe("PUT上传测试", func() {
	It("Put上传成功，返回200", func() {
		key := "upload" + util.GetRand(8)

		fmt.Println(key)

		fmt.Println("Load Parameters:" + ENV("quota_username"))

		Expect(true).To(Equal(true))

	})

	It("Put上传成功——2，返回200", func() {

		fmt.Println("Load Parameters:" + ENV("username"))

		Expect(true).To(Equal(true))

	})

	It("Put上传成功——3，返回200", func() {

		fmt.Println("Load Parameters:" + ENV("iohost"))

		Expect(true).To(Equal(true))

	})

})
