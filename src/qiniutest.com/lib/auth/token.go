package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	// . "qiniutest.com/configs"
)

func GetUptoken(bucket, ak, sk string) string {
	policy := PutPolicy{
		Scope: bucket,
	}
	return policy.MakeUptoken(ak, sk)
}

func GetPrivateUrl(baseUrl, ak, sk string) string {
	//需不需要设置默认值？？
	if ak == "" && sk == "" {
		println("ak & sk can't be null")
	}
	policy := PutPolicy{}

	return policy.MakePrivateUrl(baseUrl, ak, sk)
}

// QBox Authorization
// DOC: https://github.com/qbox/bs-apidoc/blob/master/apidoc/v6/auths/Qbox.md
func GetQboxToken(uri, body, ak, sk string) string {
	u, err := url.Parse(uri)
	if err != nil {
		println("Parse url failed, url = %d", uri)
	}

	data := u.Path

	if u.RawQuery != "" {
		data += "?" + u.RawQuery
	}
	data += "\n"

	if body != "" {
		data += body
	}

	h := hmac.New(sha1.New, []byte(sk))
	h.Write([]byte(data))
	sign := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return ak + ":" + sign
}

//QUser Authorization
//DOC: https://github.com/qbox/bs-apidoc/blob/master/apidoc/v6/auths/QUser.md
// func GetQuserToken(method, uri, body, ak, sk string) string {

// }
