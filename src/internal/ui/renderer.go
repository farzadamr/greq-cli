package ui

import (
	"fmt"

	"github.com/farzadamr/greq-cli/internal/model"
)

func RenderHeader(fileName string, testCount int, env string) {
	fmt.Println(Theme.Title.Render("🔍 Loading suite: "), fileName)
	fmt.Println(Theme.Title.Render("🌐 Environment: "), env)
	fmt.Println(Line)
	fmt.Printf(Theme.PurpleText.Render("➜ Running %d tests...\n"), testCount)
}

func RenderTestsResult(results []model.TestResult) {
	for i, res := range results {
		counter := fmt.Sprintf("%d. ", i+1)
		status := fmt.Sprintf("Status:\t%d", res.Response.StatusCode)
		time := fmt.Sprintf("Time:\t%d ms", res.Time.Milliseconds())
		assertionStatus := fmt.Sprintf("✗ Status code not equal! expected: %d, but: %d", res.Expect.Status, res.Response.StatusCode)
		if res.Passed {
			assertionStatus = fmt.Sprintf("✔ Status code is %d", res.Response.StatusCode)
		}
		result := fmt.Sprintln("Result:\t❌ FAIL")
		if res.Passed {
			result = fmt.Sprintln("Result:\t✅ PASS")
		}

		fmt.Print(Theme.Title.Render(counter))
		fmt.Println(Theme.PurpleText.Render(res.Name))
		fmt.Println(Theme.Title.Render(status))
		fmt.Println(Theme.Title.Render(time))
		fmt.Println(Theme.Title.Render("Assertions:"))
		fmt.Println("\t", Theme.PurpleText.Render(assertionStatus))
		if res.Passed {
			fmt.Println(Theme.Success.Render(result))
		} else {
			fmt.Println(Theme.Error.Render(result))
		}

		fmt.Println(Line)
	}
}
