package failing_test

import (
	"github.com/matt-royal/biloba"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFailing(t *testing.T) {
	if os.Getenv("BILOBA_INTEGRATION_TEST") == "" {
		return
	}
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Failing Suite", []Reporter{
		biloba.NewGoTestCompatibleReporter(),
	})
}
