package e2e

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Apps")
}
