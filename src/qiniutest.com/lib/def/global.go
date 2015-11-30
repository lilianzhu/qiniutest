package def

import (
	"qiniutest.com/biz/up"
	"qiniutest.com/configs"
	"qiniutest.com/lib/auth"
	"qiniutest.com/lib/util"
)

// Declarations for configs
var ENV = configs.ENV

//Declarations for util
var GetRand = util.GetRand
var GetEtag = util.GetEtag

//Declarations for token
var GetPrivateUrl = auth.GetPrivateUrl
var GetUptoken = auth.GetUptoken

//Declarations for up
var FormUp = up.FormUp
var UpResult up.UpResult
