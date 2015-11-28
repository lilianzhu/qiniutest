package auth

import (
	// . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
)

func GenerateAcctoken(url, body, ak, sk string) {

}

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
