package fopimage_test

import (
	. "qiniutest.com/lib/def"
)

var _ = Describe("Fopimage", func() {
	It("test", func() {
		url := util.MakeBaseUrl(ENV("privatebucket"), "upload.jpg")
		mac := util.NewMac(ENV("access_key"), ENV("secret_key"))
		str := mac.MakePrivateUrl(url, nil)
		println(str)
	})
})
