package fopimage_test

import (
	. "qiniutest.com/lib/def"
	"testing"
)

func TestFopimage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fopimage Suite")
}
