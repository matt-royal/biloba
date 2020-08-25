package passing_test

import (
	"github.com/matt-royal/biloba"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFormatting(t *testing.T) {
	if os.Getenv("BILOBA_INTEGRATION_TEST") == "" {
		return
	}

	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Formatting Suite", []Reporter{
		biloba.NewGoTestCompatibleReporter(),
	})
}
