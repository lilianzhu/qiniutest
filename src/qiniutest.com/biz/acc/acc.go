package acc

import (
	"net/http"
	. "qiniutest.com/configs"
	. "qiniutest.com/lib/util"
)

func GetAccToken(username, password string) *Response {
	url := ENV("acchost") + "/oauth2/token"
	payload := "grant_type=password&username=" + username + "&password=" + password
	session := &Session{}
	session.Header = &http.Header{}
	session.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := session.Post(url, payload, nil, nil)

	return resp
}
