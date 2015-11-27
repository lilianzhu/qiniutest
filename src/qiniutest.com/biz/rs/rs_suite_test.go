package rs_test

import (
	. "qiniutest.com/lib/def"
	"testing"
)

func TestRs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rs Suite")
}
