package verifier

import (
	"bytes"
	"errors"

	"github.com/farzadamr/greq-cli/internal/model"
)

func VerifySuite(responses []model.SuiteResult) (*[]model.VerifyResult, error) {
	var verifyResults []model.VerifyResult
	for _, suite := range responses {
		var testResults []model.TestResult
		for _, test := range suite.TestsResponses {
			var result model.TestResult
			if test.HTTPResponse.StatusCode != test.Assertion.Status {
				result.Assert = append(result.Assert, model.AssertResult{
					Title:   "status",
					Passed:  false,
					Message: "Status code does not match",
				})
			}

			for _, f := range test.Assertion.Contains {
				if !bytes.Contains(test.HTTPResponse.Body, []byte(f)) {
					result.Assert = append(result.Assert, model.AssertResult{
						Title:   "contains",
						Passed:  false,
						Message: "Response body does not contain expected string: " + f,
					})
				}
			}
			result.Method = test.HTTPResponse.Method
			result.Path = test.HTTPResponse.Path
			result.Duration = test.Duration
			testResults = append(testResults, result)
		}
		verifyResults = append(verifyResults, model.VerifyResult{
			Tag:         suite.Tag,
			TestResults: testResults,
		})
	}
	if len(verifyResults) == 0 {
		return nil, errors.New("no verify results found")
	}
	return &verifyResults, nil
}
