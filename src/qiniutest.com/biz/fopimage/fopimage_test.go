package fopimage_test

import (
	"fmt"
	"qiniutest.com/biz/fopimage"
	"qiniutest.com/lib/util"
)

var _ = Describe("Fopimage", func() {
	It("test", func() {
		// fmt.Println(os.Getenv("HOME"))
		// var out, _ = util.Encode(EncodeType_Base64, "hello")
		out := util.GetEnv()
		fmt.Println(out)
		image := fopimage.ImageView2Param{"1", "2", "3", "4", "5", "6", "&"}
		fmt.Println(image.To_url())
		fmt.Println(util.Getrand(8))
	})
})
