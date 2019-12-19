package biloba_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBiloba(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Biloba Suite")
}
