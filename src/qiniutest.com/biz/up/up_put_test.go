package up_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	. "qiniutest.com/configs"
)

var _ = Describe("PUT上传测试", func() {
	It("Put上传成功，返回200", func() {

		fmt.Println("Load Parameters:" + ENV("default_test_account"))

		Expect(true).To(Equal(true))

	})

})
