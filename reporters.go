package biloba

import (
	"fmt"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/ginkgo/reporters/stenographer"
	"github.com/onsi/ginkgo/reporters/stenographer/support/go-colorable"
	"os"
	"strings"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

type gotestCompatibleReporter struct {
}

func GoLandReporter() []ginkgo.Reporter {
	if !strings.Contains(os.Getenv("XPC_SERVICE_NAME"), "goland") {
		return []ginkgo.Reporter{}
	}
	return []ginkgo.Reporter{NewGoTestCompatibleReporter()}
}

func NewGoTestCompatibleReporter() *gotestCompatibleReporter {
	return new(gotestCompatibleReporter)
}

// deprecated: Use ginkgo.RunSpecsWithDefaultAndCustomReporters with GoLandReporter() instead
func DefaultReporters() []ginkgo.Reporter {
	s := stenographer.New(!config.DefaultReporterConfig.NoColor, config.GinkgoConfig.FlakeAttempts > 1, colorable.NewColorableStdout())
	defaultReporter := reporters.NewDefaultReporter(config.DefaultReporterConfig, s)

	return append(GoLandReporter(), defaultReporter)
}

func (r *gotestCompatibleReporter) SpecWillRun(specSummary *types.SpecSummary) {
	fmt.Printf("\n=== RUN   %s\n", testName(specSummary))
}

func (r *gotestCompatibleReporter) SpecDidComplete(spec *types.SpecSummary) {
	name := testName(spec)
	seconds := spec.RunTime.Milliseconds() / 1000
	milliseconds := spec.RunTime.Milliseconds() % 1000
	durationStr := fmt.Sprintf("%d.%ds", seconds, milliseconds)
	var state string
	switch {
	case spec.Passed():
		state = "PASS"
	case spec.HasFailureState():
		state = "FAIL"
	case spec.Skipped() || spec.Pending():
		state = "SKIP"
	default:
		panic("Unknown state")
	}
	fmt.Printf("\n--- %s: %s (%s)\n", state, name, durationStr)

}

func testName(spec *types.SpecSummary) string {
	return strings.ReplaceAll(
		strings.Join(spec.ComponentTexts[1:len(spec.ComponentTexts)], " "),
		"(",
		"\\(",
	)
}

// No-Op methods for compatibility with ginkgo.Reporter

func (r *gotestCompatibleReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {}

func (r *gotestCompatibleReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {}

func (r *gotestCompatibleReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {}

func (r *gotestCompatibleReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
}

// force compatibility
var _ ginkgo.Reporter = new(gotestCompatibleReporter)
