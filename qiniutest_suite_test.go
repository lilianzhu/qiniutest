package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

 	"github.com/onsi/ginkgo/reporters"
	"testing"
)

func TestQiniutest(t *testing.T) {
	RegisterFailHandler(Fail)
	// RunSpecs(t, "Qiniutest Suite")
    junitReporter := reporters.NewJUnitReporter("junit.xml")
    RunSpecsWithDefaultAndCustomReporters(t, "Foo Suite", []Reporter{junitReporter})
}
