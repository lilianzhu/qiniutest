package acc

import (
	"net/http"
	. "qiniutest.com/configs"
	"qiniutest.com/lib/util"
)

type Gettoken struct {
	Username   string
	Password   string
	Grant_type string
}

func GetAccToken(username, password string) (*util.Response, string) {
	url := ENV("acchost") + "/oauth2/token"
	payload := "grant_type=password&username=" + username + "&password=" + password
	session := &util.Session{}
	session.Header = &http.Header{}
	session.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := session.Post(url, payload, nil, nil)

	acctoken, _ := util.ConvertToMap(resp.RawText())
	token, ok := acctoken["access_token"]
	if ok {
		return resp, token.(string)
	} else {
		return resp, "acctoken not found"
	}
}
