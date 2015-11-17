package up_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Up Suite")
}
