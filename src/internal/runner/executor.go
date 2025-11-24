package runner

import (
	"fmt"
	"net/http"
	"time"

	"github.com/farzadamr/greq-cli/internal/model"
)

type TestResult struct {
	Name   string
	Passed bool
	Time   time.Duration
	Resp   int
	Expect int
	Err    error
}

func RunSuite(suite *model.TestSuite) []TestResult {
	results := []TestResult{}

	for _, t := range suite.Tests {
		start := time.Now()
		passed := false
		var err error

		resp, e := http.Get(t.URL)
		if e != nil {
			err = e
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == t.Expect.Status {
				passed = true
			}
		}

		results = append(results, TestResult{
			Name:   t.Name,
			Passed: passed,
			Time:   time.Since(start),
			Resp:   resp.StatusCode,
			Expect: t.Expect.Status,
			Err:    err,
		})
	}

	return results

}

func PrintResults(results []TestResult) {
	for _, r := range results {
		var status string
		if r.Passed {
			status = "✅"
			fmt.Printf("[%s] %s (%v)\n", status, r.Name, r.Time)
		} else {
			status = "❌"
			fmt.Printf("[%s] %s (%v) => (%d, %d)\n", status, r.Name, r.Time, r.Resp, r.Expect)
		}

	}
}
