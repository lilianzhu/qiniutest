package acc_test

import (
	. "qiniutest.com/lib/def"
	"testing"
)

func TestAcc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acc Suite")
}
