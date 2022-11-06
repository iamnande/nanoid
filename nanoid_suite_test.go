package nanoid_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNanoID(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NanoID Suite")
}
