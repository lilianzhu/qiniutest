package main_test

import (
	"./biz/exec"
	"fmt"
	"github.com/jmcvetta/napping"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("test4", func() {

	Context("=4", func() {
		It("cd:should be 4", func() {
			Expect(4).To(Equal(4))
			val, _ := util.Decode(util.EncodeType_Base64, "MzMz")
			fmt.Println(val)

			// val,_:=util.Qiniu_base64(true,false,"333")
			// fmt.Println(val)
		})
	})

	Context("3", func() {

		It("z1:should be 3", func() {
			url := "http://7xiuga.com1.z0.glb.clouddn.com/upload.jpg?avinfo"
			resp, _ := napping.Get(url, nil, nil, nil)
			Expect(resp.Status()).To(Equal(200))
			Expect(resp.HttpResponse().Header["Content-Type"][0]).To(Equal("image/jpeg"))
		})

		PIt("smoke:p1 should be 2", func() {

			By("Finding 2")

			Expect(2).To(Equal(2))

			By("Finding 3")

			Expect(3).To(Equal(3))
		})
	})

})
