package biloba_test

import (
	"bufio"
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"os"
	"os/exec"
	"regexp"
	"time"
)

type testJsonEntry struct {
	Action string
	Test   string
	Output string
}

var _ = Describe("GoTestCompatibleReporter", func() {
	var projectRoot string

	BeforeEach(func() {
		projectRoot = os.Getenv("PWD")
	})

	When("the tests pass", func() {
		It("outputs them in a format that GoLand parses as nested", func() {
			lines := testOutputLines("./test_assets/passing")
			groups := groupByTest(lines)

			Expect(groups).To(HaveLen(5))
			Expect(groups[0]).To(Equal([]testJsonEntry{
				{Action: "output", Test: "", Output: "Running Suite: Passing Suite\n"},
				{Action: "output", Test: "", Output: "============================\n"},
				{Action: "output", Test: "", Output: "Random Seed: 1234\n"},
				{Action: "output", Test: "", Output: "Will run 4 of 4 specs\n"},
				{Action: "output", Test: "", Output: "\n"},
				{Action: "output", Test: "", Output: "\n"},
			}))
			Expect(groups[1]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 A test 1 passes", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 passes", Output: "=== RUN   level 1 A test 1 passes\n"},
				{Action: "output", Test: "level 1 A test 1 passes", Output: "level 1 A \n"},
				{Action: "output", Test: "level 1 A test 1 passes", Output: "  test 1 passes\n"},
				{Action: "output", Test: "level 1 A test 1 passes", Output: fmt.Sprintf("  %s/test_assets/passing/passing_test.go:10\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 passes", Output: "•\n"},
				{Action: "output", Test: "level 1 A test 1 passes", Output: "--- PASS: level 1 A test 1 passes (TIME)\n"},
				{Action: "output", Test: "level 1 A test 1 passes", Output: "\n"},
				{Action: "pass", Test: "level 1 A test 1 passes", Output: "\n"},
			}))
			Expect(groups[2]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 A test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "=== RUN   level 1 A test 2 passes\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "------------------------------\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "level 1 A \n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "  test 2 passes\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: fmt.Sprintf("  %s/test_assets/passing/passing_test.go:14\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "•\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "--- PASS: level 1 A test 2 passes (TIME)\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "\n"},
				{Action: "pass", Test: "level 1 A test 2 passes", Output: "\n"},
			}))

			Expect(groups[3]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 B test 1 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 passes", Output: "=== RUN   level 1 B test 1 passes\n"},
				{Action: "output", Test: "level 1 B test 1 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 passes", Output: "------------------------------\n"},
				{Action: "output", Test: "level 1 B test 1 passes", Output: "level 1 B \n"},
				{Action: "output", Test: "level 1 B test 1 passes", Output: "  test 1 passes\n"},
				{Action: "output", Test: "level 1 B test 1 passes", Output: fmt.Sprintf("  %s/test_assets/passing/passing_test.go:20\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 passes", Output: "•\n"},
				{Action: "output", Test: "level 1 B test 1 passes", Output: "--- PASS: level 1 B test 1 passes (TIME)\n"},
				{Action: "output", Test: "level 1 B test 1 passes", Output: "\n"},
				{Action: "pass", Test: "level 1 B test 1 passes", Output: "\n"},
			}))

			Expect(groups[4]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "=== RUN   level 1 B test 2 passes\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "------------------------------\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "level 1 B \n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "  test 2 passes\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: fmt.Sprintf("  %s/test_assets/passing/passing_test.go:24\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "•\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "--- PASS: level 1 B test 2 passes (TIME)\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "pass", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "PASS\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "Ran 4 of 4 Specs in TIME\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "SUCCESS! -- 4 Passed | 0 Failed | 0 Pending | 0 Skipped\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "PASS\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "Ginkgo ran 1 suite in TIME\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "Test Suite Passed\n"},
				{Action: "pass", Test: "level 1 B test 2 passes", Output: "Test Suite Passed\n"},
			}))
		})
	})

	When("the tests fail", func() {
		It("outputs them in a format that GoLand parses as nested", func() {
			lines := testOutputLines("./test_assets/failing")
			groups := groupByTest(lines)

			Expect(groups).To(HaveLen(6))

			Expect(groups[0]).To(Equal([]testJsonEntry{
				{Action: "output", Test: "", Output: "Running Suite: Failing Suite\n"},
				{Action: "output", Test: "", Output: "============================\n"},
				{Action: "output", Test: "", Output: "Random Seed: 1234\n"},
				{Action: "output", Test: "", Output: "Will run 4 of 4 specs\n"},
				{Action: "output", Test: "", Output: "\n"},
				{Action: "output", Test: "", Output: "\n"},
			}))

			Expect(groups[1]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "=== RUN   level 1 A test 1 fails\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "level 1 A \n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "  test 1 fails\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: fmt.Sprintf("  %s/test_assets/failing/failing_test.go:10\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "• Failure [TIME]\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "level 1\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: fmt.Sprintf("%s/test_assets/failing/failing_test.go:8\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "  A\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: fmt.Sprintf("  %s/test_assets/failing/failing_test.go:9\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "    test 1 fails [It]\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: fmt.Sprintf("    %s/test_assets/failing/failing_test.go:10\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "    Expected\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "        <bool>: true\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "    to equal\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "        <bool>: false\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: fmt.Sprintf("    %s/test_assets/failing/failing_test.go:11\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "------------------------------\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "--- FAIL: level 1 A test 1 fails (TIME)\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "fail", Test: "level 1 A test 1 fails", Output: "\n"},
			}))

			Expect(groups[2]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 A test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "=== RUN   level 1 A test 2 fails\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "level 1 A \n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "  test 2 fails\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: fmt.Sprintf("  %s/test_assets/failing/failing_test.go:14\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "• Failure [TIME]\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "level 1\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: fmt.Sprintf("%s/test_assets/failing/failing_test.go:8\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "  A\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: fmt.Sprintf("  %s/test_assets/failing/failing_test.go:9\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "    test 2 fails [It]\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: fmt.Sprintf("    %s/test_assets/failing/failing_test.go:14\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "    Expected\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "        <bool>: true\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "    to equal\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "        <bool>: false\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: fmt.Sprintf("    %s/test_assets/failing/failing_test.go:15\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "------------------------------\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "--- FAIL: level 1 A test 2 fails (TIME)\n"},
				{Action: "output", Test: "level 1 A test 2 fails", Output: "\n"},
				{Action: "fail", Test: "level 1 A test 2 fails", Output: "\n"},
			}))

			Expect(groups[3]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "=== RUN   level 1 B test 1 fails\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "level 1 B \n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "  test 1 fails\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: fmt.Sprintf("  %s/test_assets/failing/failing_test.go:20\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "• Failure [TIME]\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "level 1\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: fmt.Sprintf("%s/test_assets/failing/failing_test.go:8\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "  B\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: fmt.Sprintf("  %s/test_assets/failing/failing_test.go:19\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "    test 1 fails [It]\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: fmt.Sprintf("    %s/test_assets/failing/failing_test.go:20\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "    Expected\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "        <bool>: true\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "    to equal\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "        <bool>: false\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: fmt.Sprintf("    %s/test_assets/failing/failing_test.go:21\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "------------------------------\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "--- FAIL: level 1 B test 1 fails (TIME)\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "fail", Test: "level 1 B test 1 fails", Output: "\n"},
			}))

			Expect(groups[4]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "=== RUN   level 1 B test 2 fails\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "level 1 B \n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "  test 2 fails\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: fmt.Sprintf("  %s/test_assets/failing/failing_test.go:24\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "• Failure [TIME]\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "level 1\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: fmt.Sprintf("%s/test_assets/failing/failing_test.go:8\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "  B\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: fmt.Sprintf("  %s/test_assets/failing/failing_test.go:19\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "    test 2 fails [It]\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: fmt.Sprintf("    %s/test_assets/failing/failing_test.go:24\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "    Expected\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "        <bool>: true\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "    to equal\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "        <bool>: false\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: fmt.Sprintf("    %s/test_assets/failing/failing_test.go:25\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "------------------------------\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "--- FAIL: level 1 B test 2 fails (TIME)\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "fail", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "FAIL\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "Summarizing 4 Failures:\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "[Fail] level 1 A [It] test 1 fails \n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: fmt.Sprintf("%s/test_assets/failing/failing_test.go:11\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "[Fail] level 1 A [It] test 2 fails \n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: fmt.Sprintf("%s/test_assets/failing/failing_test.go:15\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "[Fail] level 1 B [It] test 1 fails \n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: fmt.Sprintf("%s/test_assets/failing/failing_test.go:21\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "[Fail] level 1 B [It] test 2 fails \n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: fmt.Sprintf("%s/test_assets/failing/failing_test.go:25\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "Ran 4 of 4 Specs in TIME\n"},
				{Action: "output", Test: "level 1 B test 2 fails", Output: "FAIL! -- 0 Passed | 4 Failed | 0 Pending | 0 Skipped\n"},
			}))

			Expect(groups[5]).To(Equal([]testJsonEntry{
				{Action: "output", Test: "TestFailing", Output: "--- FAIL: TestFailing (TIME)\n"},
				{Action: "fail", Test: "TestFailing", Output: "--- FAIL: TestFailing (TIME)\n"},
				{Action: "output", Test: "TestFailing", Output: "FAIL\n"},
				{Action: "output", Test: "TestFailing", Output: "\n"},
				{Action: "output", Test: "TestFailing", Output: "Ginkgo ran 1 suite in TIME\n"},
				{Action: "output", Test: "TestFailing", Output: "Test Suite Failed\n"},
				{Action: "fail", Test: "TestFailing", Output: "Test Suite Failed\n"},
			}))
		})
	})

	When("the some tests pass and some fail", func() {
		It("outputs them in a format that GoLand parses as nested", func() {
			lines := testOutputLines("./test_assets/mixed")
			groups := groupByTest(lines)

			Expect(groups).To(HaveLen(6))

			Expect(groups[0]).To(Equal([]testJsonEntry{
				{Action: "output", Test: "", Output: "Running Suite: Mixed Suite\n"},
				{Action: "output", Test: "", Output: "==========================\n"},
				{Action: "output", Test: "", Output: "Random Seed: 1234\n"},
				{Action: "output", Test: "", Output: "Will run 4 of 4 specs\n"},
				{Action: "output", Test: "", Output: "\n"},
				{Action: "output", Test: "", Output: "\n"},
			}))

			Expect(groups[1]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "=== RUN   level 1 A test 1 fails\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "level 1 A \n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "  test 1 fails\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: fmt.Sprintf("  %s/test_assets/mixed/mixed_test.go:10\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "• Failure [TIME]\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "level 1\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: fmt.Sprintf("%s/test_assets/mixed/mixed_test.go:8\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "  A\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: fmt.Sprintf("  %s/test_assets/mixed/mixed_test.go:9\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "    test 1 fails [It]\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: fmt.Sprintf("    %s/test_assets/mixed/mixed_test.go:10\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "    Expected\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "        <bool>: true\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "    to equal\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "        <bool>: false\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: fmt.Sprintf("    %s/test_assets/mixed/mixed_test.go:11\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "------------------------------\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "--- FAIL: level 1 A test 1 fails (TIME)\n"},
				{Action: "output", Test: "level 1 A test 1 fails", Output: "\n"},
				{Action: "fail", Test: "level 1 A test 1 fails", Output: "\n"},
			}))

			Expect(groups[2]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 A test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "=== RUN   level 1 A test 2 passes\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "level 1 A \n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "  test 2 passes\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: fmt.Sprintf("  %s/test_assets/mixed/mixed_test.go:14\n", projectRoot)},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "•\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "--- PASS: level 1 A test 2 passes (TIME)\n"},
				{Action: "output", Test: "level 1 A test 2 passes", Output: "\n"},
				{Action: "pass", Test: "level 1 A test 2 passes", Output: "\n"},
			}))

			Expect(groups[3]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "=== RUN   level 1 B test 1 fails\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "------------------------------\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "level 1 B \n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "  test 1 fails\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: fmt.Sprintf("  %s/test_assets/mixed/mixed_test.go:20\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "• Failure [TIME]\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "level 1\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: fmt.Sprintf("%s/test_assets/mixed/mixed_test.go:8\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "  B\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: fmt.Sprintf("  %s/test_assets/mixed/mixed_test.go:19\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "    test 1 fails [It]\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: fmt.Sprintf("    %s/test_assets/mixed/mixed_test.go:20\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "    Expected\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "        <bool>: true\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "    to equal\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "        <bool>: false\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: fmt.Sprintf("    %s/test_assets/mixed/mixed_test.go:21\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "------------------------------\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "--- FAIL: level 1 B test 1 fails (TIME)\n"},
				{Action: "output", Test: "level 1 B test 1 fails", Output: "\n"},
				{Action: "fail", Test: "level 1 B test 1 fails", Output: "\n"},
			}))

			Expect(groups[4]).To(Equal([]testJsonEntry{
				{Action: "run", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "=== RUN   level 1 B test 2 passes\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "level 1 B \n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "  test 2 passes\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: fmt.Sprintf("  %s/test_assets/mixed/mixed_test.go:24\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "•\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "--- PASS: level 1 B test 2 passes (TIME)\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "pass", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "FAIL\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "Summarizing 2 Failures:\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "[Fail] level 1 A [It] test 1 fails \n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: fmt.Sprintf("%s/test_assets/mixed/mixed_test.go:11\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "[Fail] level 1 B [It] test 1 fails \n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: fmt.Sprintf("%s/test_assets/mixed/mixed_test.go:21\n", projectRoot)},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "Ran 4 of 4 Specs in TIME\n"},
				{Action: "output", Test: "level 1 B test 2 passes", Output: "FAIL! -- 2 Passed | 2 Failed | 0 Pending | 0 Skipped\n"},
			}))

			Expect(groups[5]).To(Equal([]testJsonEntry{
				{Action: "output", Test: "TestMixed", Output: "--- FAIL: TestMixed (TIME)\n"},
				{Action: "fail", Test: "TestMixed", Output: "--- FAIL: TestMixed (TIME)\n"},
				{Action: "output", Test: "TestMixed", Output: "FAIL\n"},
				{Action: "output", Test: "TestMixed", Output: "\n"},
				{Action: "output", Test: "TestMixed", Output: "Ginkgo ran 1 suite in TIME\n"},
				{Action: "output", Test: "TestMixed", Output: "Test Suite Failed\n"},
				{Action: "fail", Test: "TestMixed", Output: "Test Suite Failed\n"},
			}))
		})
	})
})

func groupByTest(lines []testJsonEntry) [][]testJsonEntry {
	if len(lines) == 0 {
		return nil
	}

	var (
		groups       [][]testJsonEntry
		currentGroup []testJsonEntry
		currentTest  = lines[0].Test
	)

	for _, line := range lines {
		if currentTest != line.Test {
			groups = append(groups, currentGroup)
			currentGroup = make([]testJsonEntry, 0)
			currentTest = line.Test
		}

		currentGroup = append(currentGroup, line)
	}
	groups = append(groups, currentGroup)

	return groups
}

var (
	timeRegexp = regexp.MustCompile("\\b\\d+\\.\\d+(s|ms| seconds)\\b")
)

func standardizeTime(text string) string {
	return timeRegexp.ReplaceAllString(text, "TIME")
}

func testOutputLines(testPath string) []testJsonEntry {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("BILOBA_INTEGRATION_TEST=true ginkgo -v -noColor -seed 1234 %s | go tool test2json", testPath))
	stdOut := gbytes.NewBuffer()
	session, err := gexec.Start(cmd, stdOut, GinkgoWriter)

	Expect(err).NotTo(HaveOccurred())
	Eventually(session, 5*time.Second).Should(gexec.Exit(0))

	var (
		lines       []testJsonEntry
		currentLine testJsonEntry
	)

	scanner := bufio.NewScanner(stdOut)
	for scanner.Scan() {
		stdTime := standardizeTime(scanner.Text())
		Expect(
			json.Unmarshal([]byte(stdTime), &currentLine),
		).To(Succeed())
		lines = append(lines, currentLine)
	}

	return lines
}
