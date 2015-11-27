package up_test

import (
	. "qiniutest.com/lib/def"
	"testing"
)

func TestUp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Up Suite")
}
