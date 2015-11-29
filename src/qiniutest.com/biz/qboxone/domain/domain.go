package domain

import (
	"net/http"
	"qiniutest.com/biz/auth"
	. "qiniutest.com/configs"
	. "qiniutest.com/lib/util"
)

func PublishDomain(bucket, domain, uid string) (res *Response, err error) {
	url := ENV("rshost") + "/admin/publish/" + EncodeURL(bucket) + "/from/" + bucket + "/uid/" + uid
	token := auth.SignRequest(url, "", ENV("admin_access_key"), ENV("admin_secret_key"))

	session := &Session{}
	session.Header = &http.Header{}
	session.Header.Add("Authorization", "QBox "+token)
	session.Header.Add("User-Agent", "qiniu_qa")

	res, err = session.Post(url, "", nil, nil)
	return
}

func RepublishDomain(domain, srcOwner, dstOwner, dstTbl string) (res *Response, err error) {
	url := ENV("apihost") + "/v6/domain/republish/"
	payload := "domain=" + domain + "&srcOwner=" + srcOwner + "&dstOwner=" + dstOwner + "&dstTbl=" + dstTbl

	token := auth.SignRequest(url, payload, ENV("admin_access_key"), ENV("admin_secret_key"))

	session := &Session{}
	session.Header = &http.Header{}
	session.Header.Add("Authorization", "QBox "+token)
	session.Header.Add("User-Agent", "qiniu_qa")

	res, err = session.Post(url, payload, nil, nil)
	return
}
