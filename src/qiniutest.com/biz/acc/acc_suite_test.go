package acc_test

import (
	"qiniutest.com/biz/acc"
	. "qiniutest.com/lib/def"
	"testing"
)

func TestAcc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acc Suite")
}

//Declarations for acc
var GetAccessToken = acc.GetAccessToken
