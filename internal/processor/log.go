package processor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func WriteErrorLog(
	results []model.ProcessResult,
	outputDir string,
) error {

	var failures []model.ProcessResult

	for _, result := range results {

		if result.Success {
			continue
		}

		failures = append(
			failures,
			result,
		)
	}

	if len(failures) == 0 {
		return nil
	}

	logPath := filepath.Join(
		outputDir,
		"errors.log",
	)

	file, err := os.Create(logPath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, failure := range failures {

		_, err := fmt.Fprintf(
			file,
			"INPUT : %s\nOUTPUT: %s\nERROR : %s\n\n",
			failure.InputFile,
			failure.OutputFile,
			failure.Error,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
