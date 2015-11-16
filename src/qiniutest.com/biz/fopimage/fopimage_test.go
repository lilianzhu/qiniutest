package fopimage_test

import (
	"fmt"
	"os"
	"qiniutest.com/biz/fopimage"
	"qiniutest.com/lib/util"
)

var _ = Describe("Fopimage", func() {
	It("test", func() {
		fmt.Println(os.Getenv("HOME"))
		out, _ := util.Encode(util.EncodeType_Base64, "hello")
		fmt.Println(out)
		env := util.GetEnv()
		fmt.Println(env)
		image := fopimage.ImageView2Param{"1", "2", "3", "4", "5", "6", "&"}
		fmt.Println(image.To_url())
		fmt.Println(util.GetRand(8))
		test := util.EnvParam{}
		test.GetHost()
		fmt.Println("-----", test.Username)
	})
})
