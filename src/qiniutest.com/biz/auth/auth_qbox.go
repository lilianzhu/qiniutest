package auth

import (
	// . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
)

// QBox Authorization
// DOC: https://github.com/qbox/bs-apidoc/blob/master/apidoc/v6/auths/Qbox.md
func SignRequest(uri, body, ak, sk string) (token string) {
	u, err := url.Parse(uri)
	Expect(err).ToNot(Equal(BeNil()), "Parse url failed, url = %d", uri)

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
	token = ak + ":" + sign
	return
}
