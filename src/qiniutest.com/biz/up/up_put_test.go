package up_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"qiniutest.com/biz/up"
	. "qiniutest.com/configs"
	"qiniutest.com/lib/util"

	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var _ = Describe("PUT上传测试", func() {

	It("Put上传成功，返回200", func() {
		key := "upload" + util.GetRand(8)
		bucket := ENV("bucket")
		file := "../../data/source/upload.jpg"

		cfg := up.NewDefaultConfig()
		clt := kodo.New(0, &cfg)
		bkt := clt.Bucket(bucket)
		ctx := context.Background()

		var ret kodocli.PutRet
		err := bkt.PutFile(ctx, &ret, key, file, nil)

		Expect(err).To(BeNil())
		Expect(ret.Key).To(Equal(key))

		Expect(bkt.Delete(ctx, key)).To(BeNil())
	})
})
