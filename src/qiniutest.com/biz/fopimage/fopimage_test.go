package fopimage_test

import (
	// "os"
	// "qiniutest.com/biz/fopimage"
	. "qiniutest.com/configs"
	"qiniutest.com/lib/util"
)

var _ = Describe("Fopimage", func() {
	It("test", func() {
		url := util.MakeBaseUrl(ENV("privatebucket"), "upload.jpg")
		mac := util.NewMac(ENV("access_key"), ENV("secret_key"))
		str := mac.MakePrivateUrl(url, nil)
		println(str)
	})
})
