package processor

import "github.com/EliasLd/snap-memories-processor/internal/model"

func CountSuccess(
	results []model.ProcessResult,
) int {

	success := 0

	for _, result := range results {

		if result.Success {
			success++
		}
	}

	return success
}

func CountFailures(
	results []model.ProcessResult,
) int {

	failures := 0

	for _, result := range results {

		if !result.Success {
			failures++
		}
	}

	return failures
}
