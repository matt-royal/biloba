package failing_test

import (
	"github.com/matt-royal/biloba"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFailing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Failing Suite", []Reporter{
		biloba.NewGoTestCompatibleReporter(),
	})
}
