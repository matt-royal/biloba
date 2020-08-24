package biloba_test

import (
	"github.com/matt-royal/biloba"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBiloba(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Biloba Suite", biloba.DefaultReporters())
}
