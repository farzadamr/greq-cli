package ui

import (
	"fmt"

	"github.com/farzadamr/greq-cli/internal/model"
)

func RenderHeader(fileName string, testCount int, env string) {
	fmt.Println(Theme.Title.Render("🔍 Loading suite: "), fileName)
	fmt.Println(Theme.Title.Render("🌐 Environment: "), env)
	fmt.Println(Line)
	fmt.Println(Theme.PurpleText.Render(fmt.Sprintf("➜ Running %d tests...", testCount)))
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

		fmt.Print(Theme.Title.Render(counter))
		fmt.Println(Theme.PurpleText.Render(res.Name))
		fmt.Println(Theme.Title.Render(status))
		fmt.Println(Theme.Title.Render(time))
		fmt.Println(Theme.Title.Render("Assertions:"))
		fmt.Println("\t", Theme.PurpleText.Render(assertionStatus))
		if res.Passed {
			fmt.Println(Theme.Success.Render("Result:\t✅ PASS"))
		} else {
			fmt.Println(Theme.Error.Render("Result:\t❌ FAIL"))
		}

		fmt.Println(Line)
	}
}

func RenderTestsResultSummery(results []model.TestResult) {
	var timeSum int64
	var passed int

	fmt.Println(Theme.PurpleText.Render(fmt.Sprintf("➜ Running %d tests...", len(results))))

	for _, res := range results {
		passText := "FAIL"
		if res.Passed {
			passed++
			passText = "PASS"
		}

		line := fmt.Sprintf("➜ %s\t%s\t%d ms", res.Name, passText, res.Time.Microseconds())
		fmt.Println(Theme.Title.Render(line))

		timeSum += res.Time.Microseconds()
	}

	fmt.Print(Theme.Success.Render(fmt.Sprintf("%d passed, ", passed)))
	fmt.Print(Theme.Error.Render(fmt.Sprintf("%d failed, ", len(results)-passed)))
	fmt.Println(Theme.PurpleText.Render(fmt.Sprintf("duration %d ms", timeSum)))
}
