package memory

import (
	"fmt"
	"path/filepath"

	"github.com/EliasLd/snap-memories-processor/internal/model"
)

func BuildCollection(
	extractions []model.Extraction,
) ([]model.Media, error) {

	var collection []model.Media

	for _, extraction := range extractions {

		jsonPath := filepath.Join(
			extraction.Path,
			"json",
			"memories_history.json",
		)

		metadata, err := LoadMetadata(
			jsonPath,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"load metadata from %s: %w",
				extraction.ArchiveName,
				err,
			)
		}

		memoriesDir := filepath.Join(
			extraction.Path,
			"memories",
		)

		medias, err := ScanMemories(
			memoriesDir,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"scan memories from %s: %w",
				extraction.ArchiveName,
				err,
			)
		}

		medias, matches := MatchMetadata(
			medias,
			metadata,
		)

		if matches != len(medias) {

			return nil, fmt.Errorf(
				"%s: matched %d/%d media",
				extraction.ArchiveName,
				matches,
				len(medias),
			)
		}

		collection = append(
			collection,
			medias...,
		)
	}

	return collection, nil
}
