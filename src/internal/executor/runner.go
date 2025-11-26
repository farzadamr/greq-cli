package executor

import (
	"time"

	"github.com/farzadamr/greq-cli/internal/model"
)

func RunSuite(suite *model.TestSuite) []model.TestResult {
	results := []model.TestResult{}

	for _, t := range suite.Tests {
		start := time.Now()
		passed := false
		var err error
		response, e := SendHTTPRequest(t.Request)
		if e != nil {
			err = e
		}
		requestTime := time.Since(start)

		if response.StatusCode == t.Expect.Status {
			passed = true
		}

		results = append(results, model.TestResult{
			Name:     t.Request.Method + " " + t.Request.URL,
			Passed:   passed,
			Time:     requestTime,
			Response: *response,
			Expect:   t.Expect,
			Err:      &err,
		})
	}

	return results

}
