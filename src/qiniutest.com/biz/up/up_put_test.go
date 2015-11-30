package up_test

import (
	. "qiniutest.com/lib/def"
)

var _ = Describe("表单上传测试", func() {

	It("表单上传成功，返回200", func() {
		path := "../../data/source/upload.jpg"
		qetag, _ := GetEtag(path)
		uptoken := GetUptoken(ENV("bucket"), ENV("access_key"), ENV("secret_key"))
		key := "formup" + GetRand(8)
		resp := FormUp(key, path, uptoken)

		Expect(resp.Status()).To(Equal(200))
		upResult := UpResult
		err := JsonToMap(resp.RawText(), &upResult)
		Expect(err).To(BeNil())
		Expect(upResult.Hash).To(Equal(qetag))
		Expect(upResult.Key).To(Equal(qetag))
	})

	It("test", func() {
		upResult := UpResult //必须重新定义
		println(upResult.Hash, "++++")
	})

})
