package up

import (
	"net/http"
	. "qiniutest.com/configs"
	. "qiniutest.com/lib/util"
)

func FormUp(key, filepath, token string) (res *Response) {
	body := map[string]string{"Key": key, "filepath": filepath, "token": token}
	session := &Session{}
	session.Header = &http.Header{}
	session.Header.Add("Content-Type", "application/octet-stream")
	resp, err := session.Post(ENV("uphost"), body, nil, nil)
	if err != nil {
		println(err)
	}
	return resp
}
