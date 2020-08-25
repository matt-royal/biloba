package mixed_test

import (
	"github.com/matt-royal/biloba"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMixed(t *testing.T) {
	if os.Getenv("BILOBA_INTEGRATION_TEST") == "" {
		return
	}
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Mixed Suite", []Reporter{
		biloba.NewGoTestCompatibleReporter(),
	})
}
