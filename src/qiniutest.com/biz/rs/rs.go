package rs

import (
	"net/http"
	. "qiniutest.com/configs"
	. "qiniutest.com/lib/util"
)

func DeleteKey(bucket, key, ak, sk string) (resp *Response) {
	mac := NewMac(ak, sk)
	entry, _ := Encode(EncodeType_Base64URL, bucket+":"+key)
	url := ENV("rshost") + "/delete/" + entry

	token := mac.Generate_Acctoken(url, "", "POST")

	session := &Session{}
	session.Header = &http.Header{}
	println("QBox " + token)
	session.Header.Add("Authorization", "QBox "+token)

	resp, _ = session.Post(url, nil, nil, nil)
	return
}
