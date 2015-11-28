package domain

import (
	. "qiniutest.com/configs"

	. "qiniutest.com/lib/util"

	"qiniutest.com/biz/auth"

	"net/http"
)

func PublishDomain(bucket, domain, uid string) {
	url := ENV("rshost") + "/admin/publish/" + EncodeURL(bucket) + "/from/" + bucket + "/uid/" + uid

	token := auth.SignRequest(url, "", ENV("admin_access_key"), ENV("admin_secret_key"))

	session := &Session{}
	session.Header = &http.Header{}

	println("Token=" + token)
	session.Header.Add("Authorization", "QBox "+token)
	session.Header.Add("User-Agent", "qiniu_qa")

	resp, _ := session.Post(url, nil, nil, nil)

	println("Status: " + string(resp.Status()))
	println("Body: " + resp.RawText())
}
